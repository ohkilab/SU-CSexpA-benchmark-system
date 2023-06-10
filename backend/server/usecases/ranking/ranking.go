package ranking

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
)

type RankingInteractor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *RankingInteractor {
	return &RankingInteractor{entClient, logger}
}

func (i *RankingInteractor) GetRanking(ctx context.Context, containGuest bool, contestID int) ([]*pb.GetRankingResponse_Record, error) {
	query := i.entClient.Group.Query().WithSubmits().Where(group.HasSubmitsWith(submit.HasContestsWith(contest.ID(contestID))))
	if !containGuest {
		query.Where(group.RoleEQ(group.RoleContestant))
	}
	groups, err := query.All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch groups", err)
		return nil, err
	}
	pbGroups := lo.Map(groups, func(group *ent.Group, i int) *pb.GetRankingResponse_Record {
		maxSubmit := lo.MaxBy(group.Edges.Submits, func(x, max *ent.Submit) bool {
			return x.Score > max.Score
		})
		return &pb.GetRankingResponse_Record{
			Rank: int32(i + 1),
			Group: &pb.Group{
				Id:    group.Name,
				Score: int32(maxSubmit.Score),
				Role:  pb.Role(pb.Role_value[group.Role.String()]),
			},
		}
	})
	return pbGroups, nil
}
