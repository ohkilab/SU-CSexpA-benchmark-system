package grpc

import (
	"context"
	"net/url"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/ranking"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type backendServiceServer struct {
	authInteractor    *auth.AuthInteractor
	rankingInteractor *ranking.RankingInteractor
	submitInteractor  *submit.Interactor
	contestInteractor *contest.Interactor
	pb.UnimplementedBackendServiceServer
}

func NewBackendService(secret []byte, entClient *ent.Client, worker worker.Worker, logger *slog.Logger, tagRepository tag.Repository) pb.BackendServiceServer {
	authInteractor := auth.NewInteractor(secret, entClient, logger)
	rankingInteractor := ranking.NewInteractor(entClient, logger)
	submitInteractor := submit.NewInteractor(entClient, worker, logger, tagRepository)
	contestInteractor := contest.NewInteractor(entClient, logger)
	return &backendServiceServer{authInteractor, rankingInteractor, submitInteractor, contestInteractor, pb.UnimplementedBackendServiceServer{}}
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
	return s.submitInteractor.ListSubmits(ctx, req.GroupName, req.Status)
}

func (s *backendServiceServer) ListContests(ctx context.Context, req *pb.ListContestsRequest) (*pb.ListContestsResponse, error) {
	return s.contestInteractor.ListContests(ctx, req)
}

func (s *backendServiceServer) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	return s.authInteractor.VerifyToken(ctx), nil
}
