package grpc

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/auth"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
)

type backendServiceServer struct {
	authInteractor *auth.AuthInteractor
}

func NewBackendService(entClient *ent.Client) pb.BackendServiceServer {
	authInteractor := auth.NewInteractor(entClient)
	return &backendServiceServer{authInteractor}
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
	return s.authInteractor.PostLogin(ctx, req.Id, req.Password)
}
