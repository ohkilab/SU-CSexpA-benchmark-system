package grpc

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/ranking"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
)

type backendServiceServer struct {
	authInteractor    *auth.AuthInteractor
	rankingInteractor *ranking.RankingInteractor
}

func NewBackendService(secret string, entClient *ent.Client) pb.BackendServiceServer {
	authInteractor := auth.NewInteractor(secret, entClient)
	rankingInteractor := ranking.NewInteractor(entClient)
	return &backendServiceServer{authInteractor, rankingInteractor}
}

func (s *backendServiceServer) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingResponse, error) {
	records, err := s.rankingInteractor.GetRanking(ctx, req.ContainGuest, int(req.Year))
	if err != nil {
		return nil, err
	}
	return &pb.GetRankingResponse{Records: records}, nil
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
