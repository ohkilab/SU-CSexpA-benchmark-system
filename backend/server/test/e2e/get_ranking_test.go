package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func Test_GetRanking(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	secret := []byte("secret")
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	a01 := utils.CreateGroup(ctx, t, entClient, "a01", "aaaa", group.RoleContestant)
	a02 := utils.CreateGroup(ctx, t, entClient, "a02", "aaaa", group.RoleContestant)
	a03 := utils.CreateGroup(ctx, t, entClient, "a03", "aaaa", group.RoleContestant)
	szpp := utils.CreateGroup(ctx, t, entClient, "szpp", "aaaa", group.RoleGuest)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest1").
		SetStartAt(timejst.Now()).
		SetEndAt(timejst.Now().AddDate(1, 0, 0)).
		SetSubmitLimit(9999).
		SetCreatedAt(timejst.Now()).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		Save(ctx)
	_ = utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, a01)
	a01Submit2 := utils.CreateSubmit(ctx, t, entClient, 1000, pb.Status_SUCCESS.String(), contest, a01)
	a02Submit := utils.CreateSubmit(ctx, t, entClient, 900, pb.Status_SUCCESS.String(), contest, a02)
	a03Submit := utils.CreateSubmit(ctx, t, entClient, 300, pb.Status_SUCCESS.String(), contest, a03)
	szppSubmit := utils.CreateSubmit(ctx, t, entClient, 99999, pb.Status_SUCCESS.String(), contest, szpp)
	jwtToken, err := auth.GenerateJWTToken(secret, a01.ID, a01.CreatedAt.Year())
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// not contain guest
	resp, err := client.GetRanking(ctx, &pb.GetRankingRequest{ContestId: int32(contest.ID)})
	require.NoError(t, err)
	require.Equal(t, 3, len(resp.Records))
	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(a01.Name, string(a01.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(a01Submit2.Score), *resp.Records[0].Score)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a02.Name, string(a02.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(a02Submit.Score), *resp.Records[1].Score)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a03.Name, string(a03.Role)), resp.Records[2].Group)
	assert.Equal(t, int32(a03Submit.Score), *resp.Records[2].Score)

	// contain guest
	resp, err = client.GetRanking(ctx, &pb.GetRankingRequest{ContestId: int32(contest.ID), ContainGuest: true})
	require.NoError(t, err)
	require.Equal(t, 4, len(resp.Records))
	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(szpp.Name, string(szpp.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(szppSubmit.Score), *resp.Records[0].Score)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a01.Name, string(a01.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(a01Submit2.Score), *resp.Records[1].Score)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a02.Name, string(a02.Role)), resp.Records[2].Group)
	assert.Equal(t, int32(a02Submit.Score), *resp.Records[2].Score)
	assert.Equal(t, int32(4), resp.Records[3].Rank)
	assert.Equal(t, newPbGroup(a03.Name, string(a03.Role)), resp.Records[3].Group)
	assert.Equal(t, int32(a03Submit.Score), *resp.Records[3].Score)
}

func newPbGroup(id string, role string) *pb.Group {
	return &pb.Group{
		Id:   id,
		Role: pb.Role(pb.Role_value[role]),
	}
}
