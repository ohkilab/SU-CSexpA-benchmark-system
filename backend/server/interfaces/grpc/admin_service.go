package grpc

import (
	"context"
	"strings"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/admin"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
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
	if err := interceptor.RequireAdmin(ctx); err != nil {
		return nil, err
	}
	if req.TimeLimitPerTask <= 0 {
		return nil, status.Error(codes.InvalidArgument, "time_limit_per_task must be positive")
	}
	if strings.TrimSpace(req.Slug) == "" {
		return nil, status.Error(codes.InvalidArgument, "slug is required")
	}
	if strings.TrimSpace(req.Title) == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}
	if req.StartAt == nil {
		return nil, status.Error(codes.InvalidArgument, "start_at is required")
	}
	if req.EndAt == nil {
		return nil, status.Error(codes.InvalidArgument, "end_at is required")
	}
	if req.SubmitLimit <= 0 {
		return nil, status.Error(codes.InvalidArgument, "submit_limit must be positive")
	}
	if _, ok := pb.Validator_name[int32(req.Validator)]; !ok {
		return nil, status.Error(codes.InvalidArgument, "validator is invalid")
	}
	switch selection := req.TagSelection.(type) {
	case *pb.CreateContestRequest_Auto:
		if selection.Auto == nil || selection.Auto.Tags == nil || len(selection.Auto.Tags.Tags) == 0 {
			return nil, status.Error(codes.InvalidArgument, "auto tags are required")
		}
	case *pb.CreateContestRequest_Manual:
		if selection.Manual == nil || len(selection.Manual.TagsList) == 0 {
			return nil, status.Error(codes.InvalidArgument, "manual tags are required")
		}
		for _, tags := range selection.Manual.TagsList {
			if tags == nil || len(tags.Tags) == 0 {
				return nil, status.Error(codes.InvalidArgument, "manual tags must not contain empty attempts")
			}
		}
	default:
		return nil, status.Error(codes.InvalidArgument, "tag_selection is required")
	}
	return s.adminInteractor.CreateContest(ctx, req)
}

func (s *adminServiceServer) UpdateContest(ctx context.Context, req *pb.UpdateContestRequest) (*pb.UpdateContestResponse, error) {
	if err := interceptor.RequireAdmin(ctx); err != nil {
		return nil, err
	}
	if strings.TrimSpace(req.ContestSlug) == "" {
		return nil, status.Error(codes.InvalidArgument, "contest_slug is required")
	}
	if req.Title != nil && strings.TrimSpace(*req.Title) == "" {
		return nil, status.Error(codes.InvalidArgument, "title must not be empty")
	}
	if req.SubmitLimit != nil && *req.SubmitLimit <= 0 {
		return nil, status.Error(codes.InvalidArgument, "submit_limit must be positive")
	}
	if req.Validator != nil {
		if _, ok := pb.Validator_name[int32(*req.Validator)]; !ok {
			return nil, status.Error(codes.InvalidArgument, "validator is invalid")
		}
	}
	return s.adminInteractor.UpdateContest(ctx, req)
}

func (s *adminServiceServer) CreateGroups(ctx context.Context, req *pb.CreateGroupsRequest) (*pb.CreateGroupsResponse, error) {
	if err := interceptor.RequireAdmin(ctx); err != nil {
		return nil, err
	}
	if len(req.Groups) == 0 {
		return nil, status.Error(codes.InvalidArgument, "groups are required")
	}
	for _, g := range req.Groups {
		if g == nil {
			return nil, status.Error(codes.InvalidArgument, "group must not be nil")
		}
		if strings.TrimSpace(g.Name) == "" {
			return nil, status.Error(codes.InvalidArgument, "group name is required")
		}
		if strings.TrimSpace(g.Password) == "" {
			return nil, status.Error(codes.InvalidArgument, "group password is required")
		}
		if g.Year <= 0 {
			return nil, status.Error(codes.InvalidArgument, "group year must be positive")
		}
		if _, ok := pb.Role_name[int32(g.Role)]; !ok {
			return nil, status.Error(codes.InvalidArgument, "group role is invalid")
		}
	}
	return s.adminInteractor.CreateGroups(ctx, req)
}
