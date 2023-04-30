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
	server, conn := launchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer server.GracefulStop()
	defer conn.Close()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("aaaa"), bcrypt.DefaultCost)
	a01, _ := entClient.Group.Create().SetID("a01").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(333).Save(ctx)
	a02, _ := entClient.Group.Create().SetID("a02").SetYear(2023).SetRole(group.RoleContestant).SetScore(555).Save(ctx)
	a03, _ := entClient.Group.Create().SetID("a03").SetYear(2023).SetRole(group.RoleContestant).SetScore(444).Save(ctx)
	szpp, _ := entClient.Group.Create().SetID("szpp").SetYear(2023).SetRole(group.RoleGuest).SetScore(9999).Save(ctx)

	jwtToken, err := auth.GenerateJWTToken("secret", a01.ID, a01.Year)
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
	assert.Equal(t, resp.Records, []*pb.GetRankingResponse_Record{
		{
			Rank:  1,
			Group: newPbGroup(a02.ID, a02.Score, a02.Year, string(a02.Role)),
		},
		{
			Rank:  2,
			Group: newPbGroup(a03.ID, a03.Score, a03.Year, string(a03.Role)),
		},
		{
			Rank:  3,
			Group: newPbGroup(a01.ID, a01.Score, a01.Year, string(a01.Role)),
		},
	})

	// contain guest
	resp, err = client.GetRanking(ctx, &pb.GetRankingRequest{Year: 2023, ContainGuest: true})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, resp.Records, []*pb.GetRankingResponse_Record{
		{
			Rank:  1,
			Group: newPbGroup(szpp.ID, szpp.Score, szpp.Year, string(szpp.Role)),
		},
		{
			Rank:  1,
			Group: newPbGroup(a02.ID, a02.Score, a02.Year, string(a02.Role)),
		},
		{
			Rank:  2,
			Group: newPbGroup(a03.ID, a03.Score, a03.Year, string(a03.Role)),
		},
		{
			Rank:  3,
			Group: newPbGroup(a01.ID, a01.Score, a01.Year, string(a01.Role)),
		},
	})
}

func newPbGroup(id string, score, year int, role string) *pb.Group {
	return &pb.Group{
		Id:    id,
		Score: int32(score),
		Year:  int32(year),
		Role:  pb.Role(pb.Role_value[role]),
	}
}
