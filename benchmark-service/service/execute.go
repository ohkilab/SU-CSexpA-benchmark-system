package service

import (
	"log"
	"net"
	"net/http"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const taskKey = "task"

func (s *service) Execute(req *pb.ExecuteRequest, stream pb.BenchmarkService_ExecuteServer) error {
	ipAddr := net.ParseIP(req.IpAddr)
	if ipAddr == nil {
		return status.Error(codes.InvalidArgument, "invalid ip address")
	}
	if len(req.GroupId) == 0 {
		return status.Error(codes.InvalidArgument, "groupID must not be empty")
	}
	if len(req.GroupId) > 100 {
		return status.Error(codes.InvalidArgument, "groupID must be 100 or less")
	}

	interceptor := func(req *http.Request) {}
	resultChan := s.client.Run(stream.Context(), string(ipAddr), interceptor)

	for result := range resultChan {
		log.Println(result.HttpResult)
		if err := stream.Send(&pb.ExecuteResponse{
			Response: &pb.HttpResponse{},
		}); err != nil {
			return err
		}
	}

	return nil
}
