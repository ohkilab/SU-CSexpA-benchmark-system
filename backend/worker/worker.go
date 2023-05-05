package worker

import (
	"io"
	"log"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type Worker struct {
	entClient       *ent.Client
	benchmarkClient pb.BenchmarkServiceClient
	queue           *Queue[pb.ExecuteRequest]
}

type Task struct {
	GroupID int
	Host    string
	Tags    []string
}

const (
	threadNum    int           = 5
	attemptCount int           = 100
	attemptTime  time.Duration = 10 * time.Second
)

func New(entClient *ent.Client, benchmarkClient pb.BenchmarkServiceClient) *Worker {
	return &Worker{entClient, benchmarkClient, &Queue[pb.ExecuteRequest]{}}
}

func (w *Worker) Push(task *pb.ExecuteRequest) {
	w.queue.Push(task)
}

func (w *Worker) Run() {
	for {
		time.Sleep(100 * time.Millisecond)

		task := w.queue.Pop()
		if task == nil {
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		stream, err := w.benchmarkClient.Execute(ctx, task)
		if err != nil {
			log.Println(err)
		}

		eg := &errgroup.Group{}
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
			}

			eg.Go(func() error {
				resp := resp
				// TODO: validation
				_, err := w.entClient.TaskResult.Create().
					SetRequestPerSec(int(resp.RequestsPerSecond)).
					SetURL(resp.Request.Url).
					SetMethod(resp.Request.Method.String()).
					SetRequestContentType(resp.Request.ContentType).
					SetRequestBody(resp.Request.Body).
					SetResponseCode(resp.Response.StatusCode).
					SetResponseContentType(resp.Response.ContentType).
					SetResponseBody(resp.Response.Body).
					SetThreadNum(threadNum).
					SetAttemptCount(attemptCount).
					SetAttemptTime(int(attemptTime.Seconds())).
					SetCreatedAt(timejst.Now()).
					Save(ctx)
				return err
			})
		}
		if err := eg.Wait(); err != nil {
			log.Println(err)
		}

		now := timejst.Now()
		if _, err := w.entClient.Submit.Update().SetCompletedAt(now).SetUpdatedAt(now).Save(ctx); err != nil {
			log.Println("ERROR", err)
		}
	}
}
