package contest

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *Interactor {
	return &Interactor{entClient, logger}
}

func (i *Interactor) ListContests(ctx context.Context, req *pb.ListContestsRequest) (*pb.ListContestsResponse, error) {
	contests, err := i.entClient.Contest.Query().All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch contests", err)
		return nil, status.Error(codes.Internal, "failed to fetch contests")
	}
	return &pb.ListContestsResponse{
		Contests: lo.Map(contests, func(contest *ent.Contest, i int) *pb.Contest {
			return ToPbContest(contest)
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
		Contest: ToPbContest(contest),
	}, nil
}

func ToPbContest(c *ent.Contest) *pb.Contest {
	return &pb.Contest{
		Id:          int32(c.ID),
		Slug:        c.Slug,
		Title:       c.Title,
		StartAt:     timestamppb.New(c.StartAt),
		EndAt:       timestamppb.New(c.EndAt),
		SubmitLimit: int32(c.SubmitLimit),
		TagSelectionLogic: func() pb.TagSelectionLogicType {
			switch c.TagSelectionLogic {
			case contest.TagSelectionLogicAuto:
				return pb.TagSelectionLogicType_AUTO
			case contest.TagSelectionLogicManual:
				return pb.TagSelectionLogicType_MANUAL
			default:
				return -1 // unreachable
			}
		}(),
		Validator: pb.Validator(pb.Validator_value[c.Validator]),
	}
}
