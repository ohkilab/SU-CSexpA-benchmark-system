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

type Worker interface {
	Push(*pb.ExecuteRequest)
	Run()
}

type worker struct {
	entClient       *ent.Client
	benchmarkClient pb.BenchmarkServiceClient
	queue           *Queue[pb.ExecuteRequest]
}

func New(entClient *ent.Client, benchmarkClient pb.BenchmarkServiceClient) *worker {
	return &worker{entClient, benchmarkClient, &Queue[pb.ExecuteRequest]{}}
}

func (w *worker) Push(task *pb.ExecuteRequest) {
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
		stream, err := w.benchmarkClient.Execute(ctx, task)
		if err != nil {
			log.Println(err)
		}

		eg := &errgroup.Group{}
		score := 0
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

			score += int(resp.RequestsPerSecond)
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
					SetThreadNum(int(resp.ThreadNum)).
					SetAttemptCount(int(resp.AttemptCount)).
					SetCreatedAt(timejst.Now()).
					Save(ctx)
				return err
			})
		}
		if err := eg.Wait(); err != nil {
			log.Println(err)
		}

		now := timejst.Now()
		if _, err := w.entClient.Submit.Update().
			SetCompletedAt(now).
			SetUpdatedAt(now).
			SetScore(score).
			Save(ctx); err != nil {
			log.Println("ERROR", err)
		}
	}
}
