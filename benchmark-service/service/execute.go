package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"syscall"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
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
	validator, ok := s.validatorMap[req.Validator]
	if !ok {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("the validator is not supported(slug: %s)", req.Validator.String()))
	}

	for _, task := range req.Tasks {
		uri, err := url.ParseRequestURI(task.Request.Url)
		if err != nil {
			log.Println(err)
			return status.Error(codes.InvalidArgument, "invalid url")
		}
		uri.RawQuery = uri.Query().Encode()

		ctx, cancel := context.WithTimeout(stream.Context(), time.Duration(req.TimeLimitPerTask))
		result, err := s.client.Run(ctx, uri.String(), benchmark.OptThreadNum(int(task.ThreadNum)), benchmark.OptAttemptCount(int(task.AttemptCount)))
		cancel()
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

		if err := validateAndSend(stream, validator, uri, task, result); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func validateAndSend(
	stream pb.BenchmarkService_ExecuteServer,
	validator validation.Validator,
	uri *url.URL,
	task *pb.Task,
	result *benchmark.RunResult,
) error {
	timeElapsed := time.Duration(0)
	for _, httpResult := range result.Results {
		if err := validator.Validate(uri, httpResult.Body); err != nil {
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
		timeElapsed += httpResult.ResponseTime
	}

	totalRequests := int32(len(result.Results))
	requestsPerSecond := calculateRequestsPerSecond(totalRequests, timeElapsed)
	if result.TimedOut && requestsPerSecond <= 0 {
		msg := "タイムアウトしました"
		return stream.Send(&pb.ExecuteResponse{
			Ok:                false,
			ErrorMessage:      &msg,
			TimeElapsed:       timeElapsed.Microseconds(),
			TotalRequests:     0,
			RequestsPerSecond: 0,
			Task:              task,
			Status:            backendpb.Status_TIMEOUT,
		})
	}

	return stream.Send(&pb.ExecuteResponse{
		Ok:                true,
		TimeElapsed:       timeElapsed.Microseconds(),
		TotalRequests:     totalRequests,
		RequestsPerSecond: requestsPerSecond,
		Task:              task,
		Status:            backendpb.Status_SUCCESS,
	})
}

func calculateRequestsPerSecond(totalRequests int32, timeElapsed time.Duration) int32 {
	if totalRequests <= 0 || timeElapsed <= 0 {
		return 0
	}
	return int32(float64(totalRequests) * 10 / timeElapsed.Seconds())
}
