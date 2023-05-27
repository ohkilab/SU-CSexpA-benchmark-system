package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func Test_VerifyToken(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithEntClient(entClient), grpc.WithJwtSecret("secret"))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), 1, 2023)
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err := client.VerifyToken(ctx, &pb.VerifyTokenRequest{})
	require.NoError(t, err)
	assert.True(t, resp.Ok)

	jwtToken, err = auth.GenerateJWTToken([]byte("sec"), 1, 2023)
	require.NoError(t, err)
	meta = metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err = client.VerifyToken(ctx, &pb.VerifyTokenRequest{})
	require.NoError(t, err)
	assert.False(t, resp.Ok)
}
