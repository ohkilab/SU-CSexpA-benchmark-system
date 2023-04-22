package grpc

import (
	"context"

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

func (s *healthcheckServiceServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	if len(req.Ping) == 0 {
		return nil, status.Error(codes.InvalidArgument, "field ping: must not be empty")
	}
	if len(req.Ping) > 100000 {
		return nil, status.Error(codes.InvalidArgument, "field ping: length must be less than 100000")
	}
	return s.healthcheckInteractor.Ping(ctx, req.Ping)
}
