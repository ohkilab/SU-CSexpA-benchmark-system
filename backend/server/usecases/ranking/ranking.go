package ranking

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
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

func (i *RankingInteractor) GetRanking(ctx context.Context, containGuest bool, year int) ([]*pb.GetRankingResponse_Record, error) {
	query := i.entClient.Group.Query().Where(group.YearEQ(year)).Order(group.ByScore(sql.OrderDesc()))
	if !containGuest {
		query.Where(group.RoleEQ(group.RoleContestant))
	}
	groups, err := query.All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch groups", err)
		return nil, err
	}
	pbGroups := lo.Map(groups, func(group *ent.Group, i int) *pb.GetRankingResponse_Record {
		return &pb.GetRankingResponse_Record{
			Rank: int32(i + 1),
			Group: &pb.Group{
				Id:    group.Name,
				Score: int32(group.Score),
				Year:  int32(group.Year),
				Role:  pb.Role(pb.Role_value[group.Role.String()]),
			},
		}
	})
	return pbGroups, nil
}
