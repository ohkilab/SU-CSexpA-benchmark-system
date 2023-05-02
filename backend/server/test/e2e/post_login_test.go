package e2e

import (
	"context"
	"strings"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_PostLogin(t *testing.T) {
	ctx := context.Background()
	entClient := enttestOpen(ctx, t)
	defer entClient.Close()
	server, conn := launchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer server.GracefulStop()
	defer conn.Close()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	g, _ := entClient.Group.Create().
		SetID("test").
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(group.RoleContestant).
		SetScore(12345).
		SetYear(2023).
		Save(ctx)

	// success
	resp, err := client.PostLogin(ctx, &pb.PostLoginRequest{Id: "test", Password: "test"})
	assert.NoError(t, err)
	assert.Equal(t, g.ID, resp.Group.Id)
	assert.Equal(t, int32(g.Year), resp.Group.Year)
	assert.Equal(t, int32(g.Score), resp.Group.Score)
	assert.Equal(t, g.Role, group.Role(strings.ToLower(pb.Role_name[int32(resp.Group.Role)])))
	assert.NotEmpty(t, resp.Token)

	// failed
	resp, err = client.PostLogin(ctx, &pb.PostLoginRequest{Id: "test", Password: "tset"})
	assert.Nil(t, resp)
	assert.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code())

	// TODO: test using jwt token
}
