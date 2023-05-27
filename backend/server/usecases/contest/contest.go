package contest

import (
	"context"
	"log"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/contest"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
}

func NewInteractor(entClient *ent.Client) *Interactor {
	return &Interactor{entClient}
}

func (i *Interactor) ListContests(ctx context.Context, req *pb.ListContestsRequest) (*pb.ListContestsResponse, error) {
	contests, err := i.entClient.Contest.Query().Where(contest.Year(int(req.Year))).All(ctx)
	if err != nil {
		log.Println(err)
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
		Year:        int32(contest.Year),
	}
}
