package ranking

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

type RankingInteractor struct {
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger) *RankingInteractor {
	return &RankingInteractor{entClient, logger}
}

func (i *RankingInteractor) GetRanking(ctx context.Context, containGuest bool, contestSlug string) ([]*pb.GetRankingResponse_Record, error) {
	query := i.entClient.Group.Query().WithSubmits(func(sq *ent.SubmitQuery) {
		sq.Where(submit.HasContestsWith(contest.Slug(contestSlug))).Order(submit.ByScore(sql.OrderDesc()))
	})
	if !containGuest {
		query.Where(group.RoleEQ(group.RoleContestant))
	}
	groups, err := query.All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch groups", err)
		return nil, err
	}
	slices.SortFunc(groups, func(left, right *ent.Group) bool {
		return compareGroup(left, right) == -1
	})

	rank := int32(1)
	pbGroups := lo.Map(groups, func(group *ent.Group, i int) *pb.GetRankingResponse_Record {
		var score *int32
		if len(group.Edges.Submits) > 0 {
			score = lo.ToPtr(int32(group.Edges.Submits[0].Score))
		}
		if i-1 >= 0 {
			if compareGroup(groups[i-1], group) != 0 {
				rank++
			}
		}
		return &pb.GetRankingResponse_Record{
			Rank: rank,
			Group: &pb.Group{
				Name: group.Name,
				Role: pb.Role(pb.Role_value[group.Role.String()]),
			},
			Score: score,
		}
	})
	return pbGroups, nil
}

func compareGroup(left, right *ent.Group) int {
	leftScore := 0
	var leftDate time.Time
	if len(left.Edges.Submits) > 0 {
		leftScore = left.Edges.Submits[0].Score
		leftDate = left.Edges.Submits[0].SubmitedAt
	}
	rightScore := 0
	var rightDate time.Time
	if len(right.Edges.Submits) > 0 {
		rightScore = right.Edges.Submits[0].Score
		rightDate = right.Edges.Submits[0].SubmitedAt
	}
	if leftScore > rightScore {
		return -1
	} else if leftScore == rightScore {
		return leftDate.Compare(rightDate)
	} else {
		return 1
	}
}
