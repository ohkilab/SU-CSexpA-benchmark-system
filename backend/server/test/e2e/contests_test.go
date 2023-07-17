package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_ListContests(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	contest1 := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest1", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	contest2 := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest2", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)

	resp, err := client.ListContests(ctx, &pb.ListContestsRequest{})
	require.NoError(t, err)
	require.Equal(t, 2, len(resp.Contests))
	assert.Equal(t, contest1.ID, int(resp.Contests[0].Id))
	assert.Equal(t, contest2.ID, int(resp.Contests[1].Id))
}

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
	client := pb.NewBackendServiceClient(conn)

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
	})
	require.NoError(t, err)
	assert.Equal(t, "test contest", resp.Contest.Title)
	assert.Equal(t, "test-contest", resp.Contest.Slug)
	assert.Equal(t, 329, int(resp.Contest.SubmitLimit))
	assert.Equal(t, pb.TagSelectionLogicType_AUTO, resp.Contest.TagSelectionLogic)
}
