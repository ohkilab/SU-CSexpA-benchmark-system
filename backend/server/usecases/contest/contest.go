package contest

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	pkgcontest "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient     *ent.Client
	logger        *slog.Logger
	tagRepository tag.Repository
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger, tagRepository tag.Repository) *Interactor {
	return &Interactor{entClient, logger, tagRepository}
}

func (i *Interactor) ListContests(ctx context.Context, req *pb.ListContestsRequest) (*pb.ListContestsResponse, error) {
	contests, err := i.entClient.Contest.Query().All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch contests", err)
		return nil, status.Error(codes.Internal, "failed to fetch contests")
	}
	return &pb.ListContestsResponse{
		Contests: lo.Map(contests, func(contest *ent.Contest, i int) *pb.Contest {
			return toPbContest(contest)
		}),
	}, nil
}

func (i *Interactor) GetContest(ctx context.Context, req *pb.GetContestRequest) (*pb.GetContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		i.logger.Error("failed to fetch contest", err)
		return nil, status.Error(codes.Internal, "failed to fetch contest")
	}
	return &pb.GetContestResponse{
		Contest: toPbContest(contest),
	}, nil
}

func (i *Interactor) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestResponse, error) {
	var tagSelectionLogic pkgcontest.TagSelectionLogic
	switch req.TagSelection.(type) {
	case *pb.CreateContestRequest_Auto:
		tagSelectionLogic = pkgcontest.TagSelectionLogicAuto
		if err := i.tagRepository.CreateRandomTag(req.Slug, req.GetAuto().Tags.Tags); err != nil {
			i.logger.Error("failed to create random tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create random tags")
		}
	case *pb.CreateContestRequest_Manual:
		tagSelectionLogic = pkgcontest.TagSelectionLogicManual
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
		Save(ctx)
	if err != nil {
		i.logger.Error("failed to create contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to create contest")
	}

	return &pb.CreateContestResponse{
		Contest: toPbContest(contest),
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
		Contest: toPbContest(contest),
	}, nil
}

func toPbContest(contest *ent.Contest) *pb.Contest {
	return &pb.Contest{
		Id:          int32(contest.ID),
		Slug:        contest.Slug,
		Title:       contest.Title,
		StartAt:     timestamppb.New(contest.StartAt),
		EndAt:       timestamppb.New(contest.EndAt),
		SubmitLimit: int32(contest.SubmitLimit),
		TagSelectionLogic: func() pb.TagSelectionLogicType {
			switch contest.TagSelectionLogic {
			case pkgcontest.TagSelectionLogicAuto:
				return pb.TagSelectionLogicType_AUTO
			case pkgcontest.TagSelectionLogicManual:
				return pb.TagSelectionLogicType_MANUAL
			default:
				return -1 // unreachable
			}
		}(),
		Validator: pb.Validator(pb.Validator_value[contest.Validator]),
	}
}
