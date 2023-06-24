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
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	now := timejst.Now()
	group1 := utils.CreateGroup(ctx, t, entClient, "test1", "hoge", group.RoleContestant)
	group2 := utils.CreateGroup(ctx, t, entClient, "test2", "hoge", group.RoleContestant)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetSubmitLimit(9999).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		SetCreatedAt(now).
		Save(ctx)
	submit1, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -2)).
		SetTaskNum(50).
		Save(ctx)
	submit2, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -1)).
		SetTaskNum(50).
		Save(ctx)
	submit3, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group2).
		SetStatus(pb.Status_CONNECTION_FAILED.String()).
		SetSubmitedAt(now).
		SetTaskNum(50).
		Save(ctx)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group1.ID, group1.CreatedAt.Year())
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err := client.ListSubmits(ctx, &pb.ListSubmitsRequest{})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, submit3.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit2.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit1.ID, int(resp.Submits[2].Id))

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		GroupName: lo.ToPtr("test2"),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 1)
	assert.Equal(t, submit3.ID, int(resp.Submits[0].Id))
	assert.Equal(t, group2.Name, resp.Submits[0].GroupName)

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		Status: lo.ToPtr(pb.Status_CONNECTION_FAILED),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 1)
	assert.Equal(t, submit3.ID, int(resp.Submits[0].Id))
	assert.Equal(t, group2.Name, resp.Submits[0].GroupName)
}
