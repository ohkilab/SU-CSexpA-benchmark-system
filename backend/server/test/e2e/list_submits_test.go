package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
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
	group, _ := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string("hoge")).
		SetRole(group.RoleContestant).
		SetScore(12345).
		SetYear(2023).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetYear(2023).
		SetSubmitLimit(9999).
		SetCreatedAt(now).
		Save(ctx)
	submit1, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetStatus(pb.Status_SUCCESS.String()).
		SetYear(2023).
		SetSubmitedAt(now.AddDate(0, 0, -2)).
		Save(ctx)
	submit2, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetStatus(pb.Status_SUCCESS.String()).
		SetYear(2023).
		SetSubmitedAt(now.AddDate(0, 0, -1)).
		Save(ctx)
	submit3, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetStatus(pb.Status_SUCCESS.String()).
		SetYear(2023).
		SetSubmitedAt(now).
		Save(ctx)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.Year)
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
}
