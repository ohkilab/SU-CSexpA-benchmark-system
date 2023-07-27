package grpc

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/admin"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type adminServiceServer struct {
	adminInteractor *admin.Interactor
	pb.UnimplementedAdminServiceServer
}

func NewAdminService(entClient *ent.Client, logger *slog.Logger, tagRepository tag.Repository) pb.AdminServiceServer {
	interactor := admin.NewInteractor(entClient, logger, tagRepository)
	return &adminServiceServer{interactor, pb.UnimplementedAdminServiceServer{}}
}

func (s *adminServiceServer) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestResponse, error) {
	if req.TimeLimitPerTask == 0 {
		return nil, status.Error(codes.InvalidArgument, "time_limit_per_task is required")
	}
	if req.Slug == "" {
		return nil, status.Error(codes.InvalidArgument, "slug is required")
	}
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}
	return s.adminInteractor.CreateContest(ctx, req)
}

func (s *adminServiceServer) UpdateContest(ctx context.Context, req *pb.UpdateContestRequest) (*pb.UpdateContestResponse, error) {
	if req.ContestSlug == "" {
		return nil, status.Error(codes.InvalidArgument, "contest_slug is required")
	}
	return s.adminInteractor.UpdateContest(ctx, req)
}

func (s *adminServiceServer) CreateGroups(ctx context.Context, req *pb.CreateGroupsRequest) (*pb.CreateGroupsResponse, error) {
	return s.adminInteractor.CreateGroups(ctx, req)
}
