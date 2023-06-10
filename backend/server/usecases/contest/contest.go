package contest

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
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
			return toPbContest(contest)
		}),
	}, nil
}

func toPbContest(contest *ent.Contest) *pb.Contest {
	return &pb.Contest{
		Id:          int32(contest.ID),
		Title:       contest.Title,
		StartAt:     timestamppb.New(contest.StartAt),
		EndAt:       timestamppb.New(contest.EndAt),
		SubmitLimit: int32(contest.SubmitLimit),
	}
}
