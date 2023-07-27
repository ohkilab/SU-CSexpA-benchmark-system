package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func Test_ListSubmits(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithLimit(3))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	now := timejst.Now()
	group1 := utils.CreateGroup(ctx, t, entClient, "test1", "hoge", group.RoleContestant)
	group2 := utils.CreateGroup(ctx, t, entClient, "test2", "hoge", group.RoleContestant)
	c := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	submit1, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -6)).
		SetTaskNum(50).
		SetScore(100).
		Save(ctx)
	require.NoError(t, err)
	submit2, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -5)).
		SetTaskNum(50).
		SetScore(200).
		Save(ctx)
	require.NoError(t, err)
	submit3, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group2).
		SetStatus(pb.Status_CONNECTION_FAILED.String()).
		SetSubmitedAt(now.AddDate(0, 0, -4)).
		SetTaskNum(50).
		Save(ctx)
	require.NoError(t, err)
	submit4, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -3)).
		SetTaskNum(50).
		SetScore(300).
		Save(ctx)
	require.NoError(t, err)
	submit5, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -2)).
		SetTaskNum(50).
		SetScore(400).
		Save(ctx)
	require.NoError(t, err)
	submit6, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group2).
		SetStatus(pb.Status_CONNECTION_FAILED.String()).
		SetSubmitedAt(now.AddDate(0, 0, -1)).
		SetTaskNum(50).
		Save(ctx)
	require.NoError(t, err)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group1.ID, group1.CreatedAt.Year())
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err := client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, submit6.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit5.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit4.ID, int(resp.Submits[2].Id))
	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        2,
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, submit3.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit2.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit1.ID, int(resp.Submits[2].Id))

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
		GroupName:   lo.ToPtr("test2"),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 2)
	assert.Equal(t, submit6.ID, int(resp.Submits[0].Id))
	assert.Equal(t, group2.Name, resp.Submits[0].GroupName)
	assert.Equal(t, submit3.ID, int(resp.Submits[1].Id))
	assert.Equal(t, group2.Name, resp.Submits[1].GroupName)

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
		SortBy:      pb.ListSubmitsRequest_SCORE.Enum(),
		Status:      pb.Status_SUCCESS.Enum(),
		IsDesc:      lo.ToPtr(false),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, submit1.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit2.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit4.ID, int(resp.Submits[2].Id))
}
