package e2e

import (
	"context"
	"testing"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
)

func Test_PingUnary(t *testing.T) {
	ctx := context.Background()
	conn, closeFunc := launchGrpcServer(t)
	defer closeFunc()
	client := pb.NewHealthcheckServiceClient(conn)
	resp, err := client.PingUnary(ctx, &pb.PingUnaryRequest{Ping: "ping"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ping pong:)", resp.Pong)
}
