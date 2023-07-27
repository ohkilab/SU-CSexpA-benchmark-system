package admin

import (
	"context"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	u_contest "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/contest"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interactor struct {
	entClient     *ent.Client
	logger        *slog.Logger
	tagRepository tag.Repository
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger, tagRepository tag.Repository) *Interactor {
	return &Interactor{entClient, logger, tagRepository}
}

func (i *Interactor) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestResponse, error) {
	var tagSelectionLogic contest.TagSelectionLogic
	switch req.TagSelection.(type) {
	case *pb.CreateContestRequest_Auto:
		tagSelectionLogic = contest.TagSelectionLogicAuto
		if err := i.tagRepository.CreateRandomTag(req.Slug, req.GetAuto().Tags.Tags); err != nil {
			i.logger.Error("failed to create random tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create random tags")
		}
	case *pb.CreateContestRequest_Manual:
		tagSelectionLogic = contest.TagSelectionLogicManual
		tagsList := make([][]string, 0, len(req.GetManual().TagsList))
		for _, tags := range req.GetManual().TagsList {
			tagsList = append(tagsList, tags.Tags)
		}
		if err := i.tagRepository.CreateTags(req.Slug, tagsList); err != nil {
			i.logger.Error("failed to create manual tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create manual tags")
		}
	}

	contest, err := i.entClient.Contest.Create().
		SetTitle(req.Title).
		SetSlug(req.Slug).
		SetStartAt(req.StartAt.AsTime()).
		SetEndAt(req.EndAt.AsTime()).
		SetSubmitLimit(int(req.SubmitLimit)).
		SetTagSelectionLogic(tagSelectionLogic).
		SetValidator(req.Validator.String()).
		SetCreatedAt(timejst.Now()).
		SetTimeLimitPerTask(int64(time.Duration(req.TimeLimitPerTask) * time.Second)).
		Save(ctx)
	if err != nil {
		i.logger.Error("failed to create contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to create contest")
	}

	return &pb.CreateContestResponse{
		Contest: u_contest.ToPbContest(contest),
	}, nil
}

func (i *Interactor) UpdateContest(ctx context.Context, req *pb.UpdateContestRequest) (*pb.UpdateContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		i.logger.Error("failed to fetch contest", err)
		return nil, status.Error(codes.Internal, "failed to fetch contest")
	}

	if req.Title != nil {
		contest.Title = *req.Title
	}
	if req.StartAt != nil {
		contest.StartAt = req.StartAt.AsTime()
	}
	if req.EndAt != nil {
		contest.EndAt = req.EndAt.AsTime()
	}
	if req.SubmitLimit != nil {
		contest.SubmitLimit = int(*req.SubmitLimit)
	}
	if req.Validator != nil {
		contest.Validator = req.Validator.String()
	}

	contest, err = contest.Update().Save(ctx)
	if err != nil {
		i.logger.Error("failed to update contest", err)
		return nil, status.Error(codes.Internal, "failed to update contest")
	}

	return &pb.UpdateContestResponse{
		Contest: u_contest.ToPbContest(contest),
	}, nil
}

func (i *Interactor) CreateGroups(ctx context.Context, req *pb.CreateGroupsRequest) (*pb.CreateGroupsResponse, error) {
	createdGroups := make([]*pb.Group, 0, len(req.Groups))
	for _, g := range req.Groups {
		b, err := bcrypt.GenerateFromPassword([]byte(g.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to generate password")
		}
		group, err := i.entClient.Group.Create().
			SetName(g.Name).
			SetEncryptedPassword(string(b)).
			SetRole(g.Role.String()).
			SetCreatedAt(timejst.Now()).
			Save(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to create group: "+err.Error())
		}
		createdGroups = append(createdGroups, &pb.Group{
			Name: group.Name,
			Role: pb.Role(pb.Role_value[group.Role]),
		})
	}
	return &pb.CreateGroupsResponse{
		Groups: createdGroups,
	}, nil
}
