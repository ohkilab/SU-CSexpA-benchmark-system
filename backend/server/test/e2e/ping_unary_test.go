package e2e

import (
	"context"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PingUnary(t *testing.T) {
	ctx := context.Background()
	conn, closeFunc := utils.LaunchGrpcServer(t)
	defer closeFunc()
	client := pb.NewHealthcheckServiceClient(conn)
	resp, err := client.PingUnary(ctx, &pb.PingUnaryRequest{Ping: "ping"})
	require.NoError(t, err)
	assert.Equal(t, "ping pong:)", resp.Pong)
}
