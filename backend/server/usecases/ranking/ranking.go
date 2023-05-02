package ranking

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/samber/lo"
)

type RankingInteractor struct {
	entClient *ent.Client
}

func NewInteractor(entClient *ent.Client) *RankingInteractor {
	return &RankingInteractor{entClient}
}

func (i *RankingInteractor) GetRanking(ctx context.Context, containGuest bool, year int) ([]*pb.GetRankingResponse_Record, error) {
	query := i.entClient.Group.Query().Where(group.YearEQ(year)).Order(group.ByScore(sql.OrderDesc()))
	if !containGuest {
		query.Where(group.RoleEQ(group.RoleContestant))
	}
	groups, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	pbGroups := lo.Map(groups, func(group *ent.Group, i int) *pb.GetRankingResponse_Record {
		return &pb.GetRankingResponse_Record{
			Rank: int32(i + 1),
			Group: &pb.Group{
				Id:    group.ID,
				Score: int32(group.Score),
				Year:  int32(group.Year),
				Role:  pb.Role(pb.Role_value[group.Role.String()]),
			},
		}
	})
	return pbGroups, nil
}
