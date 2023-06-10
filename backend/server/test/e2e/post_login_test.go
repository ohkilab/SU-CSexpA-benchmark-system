package e2e

import (
	"context"
	"strings"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_PostLogin(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	g, _ := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(group.RoleContestant).
		SetCreatedAt(timejst.Now()).
		Save(ctx)

	// success
	resp, err := client.PostLogin(ctx, &pb.PostLoginRequest{Id: "test", Password: "test"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, g.Name, resp.Group.Id)
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
