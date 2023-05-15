package worker

import (
	"io"
	"log"
	"sync"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/samber/lo"
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
}

type worker struct {
	entClient       *ent.Client
	benchmarkClient benchmarkpb.BenchmarkServiceClient
	queue           *Queue[Task]
}

func New(entClient *ent.Client, benchmarkClient benchmarkpb.BenchmarkServiceClient) *worker {
	return &worker{entClient, benchmarkClient, &Queue[Task]{}}
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		log.Println("Start benchmark", task.Req)
		stream, err := w.benchmarkClient.Execute(ctx, task.Req)
		if err != nil {
			log.Println(err)
			_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
				SetScore(0).
				SetMessage("failed to connect to benchmark-service").
				SetStatus(backend.Status_INTERNAL_ERROR.String()).
				SetCompletedAt(timejst.Now()).
				SetUpdatedAt(timejst.Now()).
				Save(ctx)
			if err != nil {
				log.Println(err)
			}
			continue
		}
		_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
			SetStatus(backendpb.Status_IN_PROGRESS.String()).
			SetUpdatedAt(timejst.Now()).
			Save(ctx)
		if err != nil {
			log.Println(err)
		}

		eg := &errgroup.Group{}
		score := 0
		mu := sync.Mutex{}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("received EOF")
				if err := stream.CloseSend(); err != nil {
					log.Println(err)
				}
				submit, err := w.entClient.Submit.Query().WithTaskResults().Where(submit.ID(task.SubmitID)).Only(ctx)
				if err != nil {
					log.Println(err)
					break
				}
				if len(submit.Edges.TaskResults) == 0 {
					break
				}
				if mp := lo.Associate(submit.Edges.TaskResults, func(tr *ent.TaskResult) (string, struct{}) {
					return tr.ErrorMessage, struct{}{}
				}); len(mp) == 1 {
					if _, err := w.entClient.Submit.UpdateOneID(task.SubmitID).SetMessage(submit.Edges.TaskResults[0].ErrorMessage).Save(ctx); err != nil {
						log.Println(err)
						break
					}
				}
				break
			}
			log.Println("received", resp)
			if err != nil {
				log.Println(err)
				_, err = w.entClient.Submit.UpdateOneID(task.SubmitID).
					SetScore(0).
					SetMessage(err.Error()).
					SetStatus(backend.Status_INTERNAL_ERROR.String()).
					SetCompletedAt(timejst.Now()).
					SetUpdatedAt(timejst.Now()).
					Save(ctx)
				if err != nil {
					log.Println(err)
				}
				break
			}

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
				_, err := w.entClient.TaskResult.Create().
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
					Save(ctx)
				return err
			})
		}
		if err := eg.Wait(); err != nil {
			log.Println(err)
		}

		now := timejst.Now()
		if _, err := w.entClient.Submit.
			UpdateOneID(task.SubmitID).
			SetCompletedAt(now).
			SetUpdatedAt(now).
			SetScore(score).
			SetStatus(backendpb.Status_SUCCESS.String()).
			Save(ctx); err != nil {
			log.Println("ERROR", err)
		}
	}
}
