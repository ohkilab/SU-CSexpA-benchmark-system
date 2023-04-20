package grpc

import (
	"context"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
)

type backendServiceServer struct {
}

func NewBackendService() pb.BackendServiceServer {
	return &backendServiceServer{}
}

func (s *backendServiceServer) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingResponse, error) {
	// !!!unimplemented!!!
	return nil, nil
}

func (s *backendServiceServer) PostSubmit(ctx context.Context, req *pb.PostSubmitRequest) (*pb.PostSubmitResponse, error) {
	// !!!unimplemented!!!
	return nil, nil
}

func (s *backendServiceServer) GetSubmit(req *pb.GetSubmitRequest, stream pb.BackendService_GetSubmitServer) error {
	// !!!unimplemented!!!
	return nil
}

func (s *backendServiceServer) PostLogin(ctx context.Context, req *pb.PostLoginRequest) (*pb.PostLoginResponse, error) {
	// !!!unimplemented!!!
	return nil, nil
}
