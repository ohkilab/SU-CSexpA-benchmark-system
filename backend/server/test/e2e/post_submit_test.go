package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

func Test_PostSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := enttestOpen(ctx, t)
	defer cleanupFunc(t)
	secret := []byte("secret")
	conn, closeFunc := launchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	group, err := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(group.RoleContestant).
		SetScore(12345).
		SetYear(2023).
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	contest, err := entClient.Contest.Create().
		SetTitle("test contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetSubmitLimit(9999).
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}

	jwtToken, err := auth.GenerateJWTToken(secret, group.ID, group.Year)
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	req := &pb.PostSubmitRequest{
		IpAddr:    "10.123.456.789",
		ContestId: int32(contest.ID),
	}
	resp, err := client.PostSubmit(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.IpAddr, resp.IpAddr)
	assert.Equal(t, req.ContestId, resp.ContestId)
	assert.NotEmpty(t, resp.SubmitedAt)
}
