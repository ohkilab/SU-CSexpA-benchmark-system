package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_CreateContest(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tagRepository := tag.MockRepository(nil, nil,
		func(contestSlug string, tags []string) error {
			return nil
		},
		func(contestSlug string, tagsList [][]string) error {
			return nil
		},
	)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	group := utils.CreateGroup(ctx, t, entClient, "test-group", "test-group", group.RoleGuest)
	ctx = utils.WithJWT(ctx, t, group.ID, group.CreatedAt.Year())

	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(1, 0, 0))
	resp, err := client.CreateContest(ctx, &pb.CreateContestRequest{
		Title:       "test contest",
		Slug:        "test-contest",
		StartAt:     startAt,
		EndAt:       endAt,
		SubmitLimit: 329,
		TagSelection: &pb.CreateContestRequest_Auto{
			Auto: &pb.TagSelectionLogicAuto{
				Type: pb.TagSelectionLogicType_AUTO,
				Tags: &pb.Tags{
					Tags: []string{"tag1", "tag2"},
				},
			},
		},
		TimeLimitPerTask: 30,
	})
	require.NoError(t, err)
	assert.Equal(t, "test contest", resp.Contest.Title)
	assert.Equal(t, "test-contest", resp.Contest.Slug)
	assert.Equal(t, 329, int(resp.Contest.SubmitLimit))
	assert.Equal(t, pb.TagSelectionLogicType_AUTO, resp.Contest.TagSelectionLogic)
}
