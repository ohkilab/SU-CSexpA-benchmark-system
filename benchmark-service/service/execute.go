package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"sync"
	"syscall"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Execute(req *pb.ExecuteRequest, stream pb.BenchmarkService_ExecuteServer) error {
	if len(req.GroupId) == 0 {
		return status.Error(codes.InvalidArgument, "groupID must not be empty")
	}
	if len(req.GroupId) > 100 {
		return status.Error(codes.InvalidArgument, "groupID must be 100 or less")
	}
	if len(req.Tasks) == 0 {
		return status.Error(codes.InvalidArgument, "tasks must not be empty")
	}

	log.Println("Start executing: ", req)
	validator, ok := s.validatorMap[req.Validator.String()]
	if !ok {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("the validator is not supported(slug: %s)", req.Validator.String()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.client.CheckConnection(ctx, req.Tasks[0].Request.Url); err != nil {
		log.Println(err)
		return status.Error(codes.FailedPrecondition, "failed to connect with the server")
	}

	mu := &sync.Mutex{}
	eg := &errgroup.Group{}
	for _, task := range req.Tasks {
		uri, err := url.ParseRequestURI(task.Request.Url)
		if err != nil {
			log.Println(err)
			return status.Error(codes.InvalidArgument, "invalid url")
		}
		uri.RawQuery = uri.Query().Encode()

		results, err := s.client.Run(stream.Context(), uri.String(), benchmark.OptThreadNum(int(task.ThreadNum)), benchmark.OptAttemptCount(int(task.AttemptCount)))
		if err != nil {
			log.Println(err)
			if errors.Is(err, syscall.ECONNREFUSED) {
				msg := "サーバーとの接続ができませんでした"
				if err := stream.Send(&pb.ExecuteResponse{
					Ok:           false,
					ErrorMessage: &msg,
					Task:         task,
					Status:       backendpb.Status_CONNECTION_FAILED,
				}); err != nil {
					log.Println(err)
				}
				continue
			}
			return status.Error(codes.Internal, "Internal Server Error")
		}

		task := task
		eg.Go(func() error {
			mu.Lock()
			defer mu.Unlock()

			if err := validateAndSend(stream, validator, uri, task, results); err != nil {
				log.Println(err)
			}
			return nil
		})
	}

	return eg.Wait()
}

func validateAndSend(
	stream pb.BenchmarkService_ExecuteServer,
	validator validation.Validator,
	uri *url.URL,
	task *pb.Task,
	results []*benchmark.HttpResult,
) error {
	timeElapsed := time.Duration(0)
	for _, result := range results {
		if err := validator.Validate(uri, result.Body); err != nil {
			errMsg := err.Error()
			validationErr := &errMsg

			if err := stream.Send(&pb.ExecuteResponse{
				Ok:                false,
				ErrorMessage:      validationErr,
				TimeElapsed:       0,
				TotalRequests:     0,
				RequestsPerSecond: 0,
				Task:              task,
				Status:            backendpb.Status_VALIDATION_ERROR,
			}); err != nil {
				return err
			}
			return nil
		}
		timeElapsed += result.ResponseTime
	}

	return stream.Send(&pb.ExecuteResponse{
		Ok:                true,
		TimeElapsed:       timeElapsed.Microseconds(),
		TotalRequests:     task.AttemptCount,
		RequestsPerSecond: int32(float64(task.AttemptCount) / timeElapsed.Seconds()),
		Task:              task,
		Status:            backendpb.Status_SUCCESS,
	})
}
