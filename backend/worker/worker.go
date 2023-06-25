package worker

import (
	"io"
	"sync"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const Timeout = 3 * time.Minute

type Worker interface {
	Push(*Task)
	Run()
}

type Task struct {
	Req      *benchmarkpb.ExecuteRequest
	SubmitID int
	GroupID  int
}

type worker struct {
	entClient       *ent.Client
	benchmarkClient benchmarkpb.BenchmarkServiceClient
	queue           *Queue[Task]
	logger          *slog.Logger
}

func New(entClient *ent.Client, benchmarkClient benchmarkpb.BenchmarkServiceClient, logger *slog.Logger) *worker {
	return &worker{entClient, benchmarkClient, &Queue[Task]{}, logger}
}

func (w *worker) Push(task *Task) {
	w.queue.Push(task)
}

func (w *worker) Run() {
	for {
		time.Sleep(100 * time.Millisecond)

		task := w.queue.Pop()
		if task == nil {
			continue
		}

		ctx := context.Background()

		w.logger.Info("start benchmark", slog.Any("req", task.Req))
		if err := w.runBenchmarkTask(task); err != nil {
			w.logger.Error("failed to run benchmark", err)
			_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
				SetScore(0).
				SetMessage("Internal Server Error(Please contact administrator)").
				SetStatus(backendpb.Status_INTERNAL_ERROR.String()).
				SetCompletedAt(timejst.Now()).
				SetUpdatedAt(timejst.Now()).
				Save(ctx)
			if err != nil {
				w.logger.Error("failed to update submit", err)
			}
			continue
		}
		w.logger.Info("benchmark succeeded")
	}
}

func (w *worker) runBenchmarkTask(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	stream, err := w.benchmarkClient.Execute(ctx, task.Req)
	if err != nil {
		w.logger.Error("failed to connect to benchmark-service", err)
		return err
	}
	_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
		SetStatus(backendpb.Status_IN_PROGRESS.String()).
		SetUpdatedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		w.logger.Error("failed to update submit", err)
		return err
	}

	eg := &errgroup.Group{}
	scores := make([]int, 0, len(task.Req.Tasks))
	pbStatus := backendpb.Status_SUCCESS
	mu := sync.Mutex{}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			if err := stream.CloseSend(); err != nil {
				w.logger.Error("failed to close stream", err)
				return err
			}

			submit, err := w.entClient.Submit.Query().WithTaskResults().Where(submit.ID(task.SubmitID)).Only(ctx)
			if err != nil {
				w.logger.Error("failed to get submit", err)
				return err
			}
			if len(submit.Edges.TaskResults) == 0 {
				w.logger.Error("failed to get task results", err)
				return err
			}
			if mp := lo.Associate(submit.Edges.TaskResults, func(tr *ent.TaskResult) (string, struct{}) {
				return tr.ErrorMessage, struct{}{}
			}); len(mp) == 1 {
				if _, err := w.entClient.Submit.UpdateOneID(task.SubmitID).SetMessage(submit.Edges.TaskResults[0].ErrorMessage).Save(ctx); err != nil {
					w.logger.Error("failed to update submit", err)
					return err
				}
			}
			break
		}
		if err != nil {
			if code := status.Code(err); code == codes.DeadlineExceeded || code == codes.Canceled {
				now := timejst.Now()
				_, err := w.entClient.Submit.
					UpdateOneID(task.SubmitID).
					SetCompletedAt(now).
					SetUpdatedAt(now).
					SetStatus(backendpb.Status_TIMEOUT.String()).
					Save(context.Background())
				return err
			}
			w.logger.Error("failed to receive benchmark response", err)
			return err
		}
		w.logger.Info("received benchmark response", slog.Any("resp", resp))

		eg.Go(func() error {
			resp := resp

			mu.Lock()
			if resp.Ok {
				scores = append(scores, int(resp.RequestsPerSecond))
			}
			if needsUpdateStatus(pbStatus, resp.Status) {
				w.logger.Info("update status", slog.Any("current status", pbStatus), slog.Any("next status", resp.Status))
				pbStatus = resp.Status
			}
			mu.Unlock()

			var errorMessage string
			if resp.ErrorMessage != nil {
				errorMessage = *resp.ErrorMessage
			}
			if _, err := w.entClient.TaskResult.Create().
				SetRequestPerSec(int(resp.RequestsPerSecond)).
				SetURL(resp.Task.Request.Url).
				SetMethod(resp.Task.Request.Method.String()).
				SetErrorMessage(errorMessage).
				SetRequestContentType(resp.Task.Request.ContentType).
				SetRequestBody(resp.Task.Request.Body).
				SetThreadNum(int(resp.Task.ThreadNum)).
				SetAttemptCount(int(resp.Task.AttemptCount)).
				SetStatus(resp.Status.String()).
				SetCreatedAt(timejst.Now()).
				SetSubmitsID(task.SubmitID).
				Save(ctx); err != nil {
				w.logger.Error("failed to save task result", err)
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	score := 0
	if pbStatus == backendpb.Status_SUCCESS {
		score = lo.Sum(scores)
	}

	now := timejst.Now()
	if _, err := w.entClient.Submit.
		UpdateOneID(task.SubmitID).
		SetCompletedAt(now).
		SetUpdatedAt(now).
		SetScore(score).
		SetStatus(pbStatus.String()).
		Save(ctx); err != nil {
		w.logger.Error("failed to update submit", err)
		return err
	}

	maxSubmit, err := w.entClient.Submit.Query().WithGroups().Where(submit.HasGroupsWith(group.ID(task.GroupID))).Order(ent.Desc(submit.FieldScore)).First(ctx)
	if err != nil {
		w.logger.Error("failed to get max submit", "error", err)
		return err
	}
	if maxSubmit.Score < score {
		if _, err := w.entClient.Group.
			UpdateOneID(maxSubmit.Edges.Groups.ID).
			SetUpdatedAt(now).
			Save(ctx); err != nil {
			w.logger.Error("failed to update group", err)
			return err
		}
	}
	return nil
}

func needsUpdateStatus(current, next backendpb.Status) bool {
	priorityMap := map[backendpb.Status]int{
		backendpb.Status_WAITING:           0,
		backendpb.Status_IN_PROGRESS:       1,
		backendpb.Status_SUCCESS:           2,
		backendpb.Status_CONNECTION_FAILED: 3,
		backendpb.Status_VALIDATION_ERROR:  4,
		backendpb.Status_INTERNAL_ERROR:    5,
		backendpb.Status_TIMEOUT:           6,
	}
	return priorityMap[current] < priorityMap[next]
}
