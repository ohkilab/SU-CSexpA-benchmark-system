package grpc

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
)

type backendServiceServer struct {
}

func newBackendService() backend.BackendServiceServer {
	return &backendServiceServer{}
}

func (s *backendServiceServer) GetRanking(ctx context.Context, req *backend.GetRankingRequest) (*backend.GetRankingResponse, error) {
	// !!!unimplemented!!!
	return nil, nil
}
