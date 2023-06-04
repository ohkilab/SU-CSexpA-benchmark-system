package worker

import (
	"io"
	"sync"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

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
			_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
				SetScore(0).
				SetMessage("Internal Server Error(Please contact administrator)").
				SetStatus(backend.Status_INTERNAL_ERROR.String()).
				SetCompletedAt(timejst.Now()).
				SetUpdatedAt(timejst.Now()).
				Save(ctx)
			if err != nil {
				w.logger.Error("failed to update submit", err)
			}
			w.logger.Error("failed to run benchmark", err)
		}
		w.logger.Info("benchmark succeeded")
	}
}

func (w *worker) runBenchmarkTask(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
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
	score := 0
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
			w.logger.Error("failed to receive benchmark response", err)
			return err
		}
		w.logger.Info("received benchmark response", slog.Any("resp", resp))

		eg.Go(func() error {
			resp := resp

			mu.Lock()
			if resp.Ok {
				score += int(resp.RequestsPerSecond)
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

	now := timejst.Now()
	if _, err := w.entClient.Submit.
		UpdateOneID(task.SubmitID).
		SetCompletedAt(now).
		SetUpdatedAt(now).
		SetScore(score).
		SetStatus(backendpb.Status_SUCCESS.String()).
		Save(ctx); err != nil {
		w.logger.Error("failed to update submit", err)
		return err
	}

	group, err := w.entClient.Group.Get(ctx, task.GroupID)
	if err != nil {
		w.logger.Error("failed to get group", err)
		return err
	}
	if group.Score < score {
		if _, err := w.entClient.Group.
			UpdateOneID(group.ID).
			SetScore(score).
			SetUpdatedAt(now).
			Save(ctx); err != nil {
			w.logger.Error("failed to update group", err)
			return err
		}
	}
	return nil
}
