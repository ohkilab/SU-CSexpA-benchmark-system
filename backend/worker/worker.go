package worker

import (
	"io"
	"log"
	"sync"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type Worker interface {
	Push(*Task)
	Run()
}

type Task struct {
	Req      *pb.ExecuteRequest
	SubmitID int
}

type worker struct {
	entClient       *ent.Client
	benchmarkClient pb.BenchmarkServiceClient
	queue           *Queue[Task]
}

func New(entClient *ent.Client, benchmarkClient pb.BenchmarkServiceClient) *worker {
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
		stream, err := w.benchmarkClient.Execute(ctx, task.Req)
		if err != nil {
			log.Println(err)
		}

		eg := &errgroup.Group{}
		score := 0
		mu := sync.Mutex{}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				if err := stream.CloseSend(); err != nil {
					log.Println(err)
				}
				break
			}
			if err != nil {
				log.Println(err)
				return
			}

			eg.Go(func() error {
				resp := resp

				mu.Lock()
				if resp.Ok {
					score += int(resp.RequestsPerSecond)
				}
				mu.Unlock()

				log.Println(resp.RequestsPerSecond)
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
			Save(ctx); err != nil {
			log.Println("ERROR", err)
		}
	}
}
