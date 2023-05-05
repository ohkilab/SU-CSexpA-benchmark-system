package service

import (
	"log"
	"net/url"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
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

	for _, task := range req.Tasks {
		uri, err := url.ParseRequestURI(task.Request.Url)
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid url")
		}
		uri.RawQuery = uri.Query().Encode()

		results, err := s.client.Run(stream.Context(), uri.String(), benchmark.OptThreadNum(int(task.ThreadNum)), benchmark.OptAttemptCount(int(task.AttemptCount)))
		if err != nil {
			return err
		}

		timeElapsed := time.Duration(0)
		for _, result := range results {
			if err := validation.Validate2022(uri, result.Body); err != nil {
				errMsg := err.Error()
				validationErr := &errMsg
				if err := stream.Send(&pb.ExecuteResponse{
					Ok:                false,
					ErrorMessage:      validationErr,
					TimeElapsed:       0,
					TotalRequests:     0,
					RequestsPerSecond: 0,
					Task:              task,
				}); err != nil {
					log.Println(err)
				}
				goto L1
			}
			timeElapsed += result.ResponseTime
		}

		if err := stream.Send(&pb.ExecuteResponse{
			Ok:                true,
			TimeElapsed:       timeElapsed.Microseconds(),
			TotalRequests:     task.AttemptCount,
			RequestsPerSecond: int32(float64(task.AttemptCount) / timeElapsed.Seconds()),
			Task:              task,
		}); err != nil {
			log.Println(err)
			return err
		}
	L1:
	}

	return nil
}
