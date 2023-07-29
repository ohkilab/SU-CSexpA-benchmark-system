package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	g := utils.CreateGroup(ctx, t, entClient, "test", "test", 2023, pb.Role_CONTESTANT)

	// success
	resp, err := client.PostLogin(ctx, &pb.PostLoginRequest{Id: "test", Password: "test"})
	require.NoError(t, err)
	assert.Equal(t, g.Name, resp.Group.Name)
	assert.Equal(t, g.Role, pb.Role_name[int32(resp.Group.Role)])
	assert.NotEmpty(t, resp.Token)

	// failed
	resp, err = client.PostLogin(ctx, &pb.PostLoginRequest{Id: "test", Password: "tset"})
	assert.Nil(t, resp)
	assert.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code())

	// TODO: test using jwt token
}

func Test_VerifyToken(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithEntClient(entClient), grpc.WithJwtSecret("secret"))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), 1, 2023, pb.Role_ADMIN.String())
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err := client.VerifyToken(ctx, &pb.VerifyTokenRequest{})
	require.NoError(t, err)
	assert.True(t, resp.Ok)

	jwtToken, err = auth.GenerateJWTToken([]byte("sec"), 1, 2023, pb.Role_ADMIN.String())
	require.NoError(t, err)
	meta = metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err = client.VerifyToken(ctx, &pb.VerifyTokenRequest{})
	require.NoError(t, err)
	assert.False(t, resp.Ok)
}
