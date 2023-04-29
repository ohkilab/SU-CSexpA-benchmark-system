package e2e

import (
	"context"
	"testing"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_PingUnary(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:3776", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewHealthcheckServiceClient(conn)
	resp, err := client.PingUnary(ctx, &pb.PingUnaryRequest{Ping: "ping"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ping pong:)", resp.Pong)
}