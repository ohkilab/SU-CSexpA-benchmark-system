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
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
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
	now := timejst.Now()
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("aaaa"), bcrypt.DefaultCost)
	a01, _ := entClient.Group.Create().SetName("a01").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
	a02, _ := entClient.Group.Create().SetName("a02").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
	a03, _ := entClient.Group.Create().SetName("a03").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
	szpp, _ := entClient.Group.Create().SetName("szpp").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleGuest).SetCreatedAt(timejst.Now()).Save(ctx)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest1").
		SetStartAt(timejst.Now()).
		SetEndAt(timejst.Now().AddDate(1, 0, 0)).
		SetSubmitLimit(9999).
		SetCreatedAt(timejst.Now()).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		Save(ctx)
	a01Submit, _ := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetScore(100).
		SetGroups(a01).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now).
		SetTaskNum(50).
		Save(ctx)
	a02Submit, _ := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetScore(900).
		SetGroups(a02).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now).
		SetTaskNum(50).
		Save(ctx)
	a03Submit, _ := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetScore(300).
		SetGroups(a03).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now).
		SetTaskNum(50).
		Save(ctx)
	szppSubmit, _ := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetScore(99999).
		SetGroups(a03).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now).
		SetTaskNum(50).
		Save(ctx)

	jwtToken, err := auth.GenerateJWTToken(secret, a01.ID, a01.CreatedAt.Year())
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// not contain guest
	resp, err := client.GetRanking(ctx, &pb.GetRankingRequest{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(a02.Name, a02Submit.Score, string(a02.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a03.Name, a03Submit.Score, string(a03.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a01.Name, a01Submit.Score, string(a01.Role)), resp.Records[2].Group)

	// contain guest
	resp, err = client.GetRanking(ctx, &pb.GetRankingRequest{ContestId: int32(contest.ID), ContainGuest: true})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(szpp.Name, szppSubmit.Score, string(szpp.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a02.Name, a02Submit.Score, string(a02.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a03.Name, a03Submit.Score, string(a03.Role)), resp.Records[2].Group)
	assert.Equal(t, int32(4), resp.Records[3].Rank)
	assert.Equal(t, newPbGroup(a01.Name, a01Submit.Score, string(a01.Role)), resp.Records[3].Group)
}

func newPbGroup(id string, score int, role string) *pb.Group {
	return &pb.Group{
		Id:    id,
		Score: int32(score),
		Role:  pb.Role(pb.Role_value[role]),
	}
}
