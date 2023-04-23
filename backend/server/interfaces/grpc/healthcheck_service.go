package grpc

import (
	"context"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/healthcheck"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type healthcheckServiceServer struct {
	healthcheckInteractor *healthcheck.HealthcheckInteractor
}

func NewHealthcheckService() pb.HealthcheckServiceServer {
	healthcheckInteractor := healthcheck.NewInteractor()
	return &healthcheckServiceServer{healthcheckInteractor}
}

func (s *healthcheckServiceServer) PingUnary(ctx context.Context, req *pb.PingUnaryRequest) (*pb.PingUnaryResponse, error) {
	if len(req.Ping) == 0 {
		return nil, status.Error(codes.InvalidArgument, "field ping: must not be empty")
	}
	if len(req.Ping) > 100000 {
		return nil, status.Error(codes.InvalidArgument, "field ping: length must be less than 100000")
	}
	resp := &pb.PingUnaryResponse{
		Pong: s.healthcheckInteractor.Ping(req.Ping),
	}
	return resp, nil
}

func (s *healthcheckServiceServer) PingServerSideStreaming(req *pb.PingServerSideStreamingRequest, stream pb.HealthcheckService_PingServerSideStreamingServer) error {
	if len(req.Ping) == 0 {
		return status.Error(codes.InvalidArgument, "field ping: must not be empty")
	}
	if len(req.Ping) > 100000 {
		return status.Error(codes.InvalidArgument, "field ping: length must be less than 100000")
	}
	for {
		if err := stream.Send(&pb.PingServerSideStreamingResponse{
			Pong: s.healthcheckInteractor.Ping(req.Ping),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}
