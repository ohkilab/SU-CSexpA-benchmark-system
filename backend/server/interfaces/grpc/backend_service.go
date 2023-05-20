package grpc

import (
	"context"
	"net/url"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/ranking"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type backendServiceServer struct {
	authInteractor    *auth.AuthInteractor
	rankingInteractor *ranking.RankingInteractor
	submitInteractor  *submit.Interactor
}

func NewBackendService(secret []byte, entClient *ent.Client, worker worker.Worker) pb.BackendServiceServer {
	authInteractor := auth.NewInteractor(secret, entClient)
	rankingInteractor := ranking.NewInteractor(entClient)
	submitInteractor := submit.NewInteractor(entClient, worker)
	return &backendServiceServer{authInteractor, rankingInteractor, submitInteractor}
}

func (s *backendServiceServer) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingResponse, error) {
	records, err := s.rankingInteractor.GetRanking(ctx, req.ContainGuest, int(req.Year))
	if err != nil {
		return nil, err
	}
	return &pb.GetRankingResponse{Records: records}, nil
}

func (s *backendServiceServer) PostSubmit(ctx context.Context, req *pb.PostSubmitRequest) (*pb.PostSubmitResponse, error) {
	if _, err := url.ParseRequestURI(req.Url); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return s.submitInteractor.PostSubmit(ctx, req)
}

func (s *backendServiceServer) GetSubmit(req *pb.GetSubmitRequest, stream pb.BackendService_GetSubmitServer) error {
	return s.submitInteractor.GetSubmit(req, stream)
}

func (s *backendServiceServer) PostLogin(ctx context.Context, req *pb.PostLoginRequest) (*pb.PostLoginResponse, error) {
	return s.authInteractor.PostLogin(ctx, req.Id, req.Password)
}

func (s *backendServiceServer) ListSubmits(ctx context.Context, req *pb.ListSubmitsRequest) (*pb.ListSubmitsResponse, error) {
	return s.submitInteractor.ListSubmits(ctx, req.GroupId, req.Status)
}
