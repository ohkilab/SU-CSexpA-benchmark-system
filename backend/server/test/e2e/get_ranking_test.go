package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

func Test_GetRanking(t *testing.T) {
	ctx := context.Background()
	entClient := enttestOpen(ctx, t)
	defer entClient.Close()
	secret := []byte("secret")
	server, conn := launchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer server.GracefulStop()
	defer conn.Close()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("aaaa"), bcrypt.DefaultCost)
	a01, _ := entClient.Group.Create().SetID("a01").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(333).Save(ctx)
	a02, _ := entClient.Group.Create().SetID("a02").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(555).Save(ctx)
	a03, _ := entClient.Group.Create().SetID("a03").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(444).Save(ctx)
	szpp, _ := entClient.Group.Create().SetID("szpp").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleGuest).SetScore(9999).Save(ctx)

	jwtToken, err := auth.GenerateJWTToken(secret, a01.ID, a01.Year)
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// not contain guest
	resp, err := client.GetRanking(ctx, &pb.GetRankingRequest{Year: 2023})
	if err != nil {
		t.Fatal(err)
	}
	for _, record := range resp.Records {
		t.Log(record.Group.Id, record.Group.Score)
	}
	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(a02.ID, a02.Score, a02.Year, string(a02.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a03.ID, a03.Score, a03.Year, string(a03.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a01.ID, a01.Score, a01.Year, string(a01.Role)), resp.Records[2].Group)

	// contain guest
	resp, err = client.GetRanking(ctx, &pb.GetRankingRequest{Year: 2023, ContainGuest: true})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int32(1), resp.Records[0].Rank)
	assert.Equal(t, newPbGroup(szpp.ID, szpp.Score, szpp.Year, string(szpp.Role)), resp.Records[0].Group)
	assert.Equal(t, int32(2), resp.Records[1].Rank)
	assert.Equal(t, newPbGroup(a02.ID, a02.Score, a02.Year, string(a02.Role)), resp.Records[1].Group)
	assert.Equal(t, int32(3), resp.Records[2].Rank)
	assert.Equal(t, newPbGroup(a03.ID, a03.Score, a03.Year, string(a03.Role)), resp.Records[2].Group)
	assert.Equal(t, int32(4), resp.Records[3].Rank)
	assert.Equal(t, newPbGroup(a01.ID, a01.Score, a01.Year, string(a01.Role)), resp.Records[3].Group)
}

func newPbGroup(id string, score, year int, role string) *pb.Group {
	return &pb.Group{
		Id:    id,
		Score: int32(score),
		Year:  int32(year),
		Role:  pb.Role(pb.Role_value[role]),
	}
}
