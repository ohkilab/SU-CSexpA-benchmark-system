package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ListContests(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	contest1, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest1").
		SetStartAt(timejst.Now()).
		SetEndAt(timejst.Now().AddDate(1, 0, 0)).
		SetSubmitLimit(9999).
		SetCreatedAt(timejst.Now()).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		Save(ctx)
	contest2, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest2").
		SetStartAt(timejst.Now()).
		SetEndAt(timejst.Now().AddDate(1, 0, 0)).
		SetSubmitLimit(9999).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		SetCreatedAt(timejst.Now()).
		Save(ctx)

	resp, err := client.ListContests(ctx, &pb.ListContestsRequest{})
	require.NoError(t, err)
	require.Equal(t, 2, len(resp.Contests))
	assert.Equal(t, contest1.ID, int(resp.Contests[0].Id))
	assert.Equal(t, contest2.ID, int(resp.Contests[1].Id))
}
