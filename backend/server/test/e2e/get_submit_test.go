package e2e

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
)

func Test_GetSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := enttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := launchGrpcServer(t, grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	submit, _ := entClient.Submit.Create().
		SetIPAddr("10.255.255.255").
		SetYear(2023).
		SetSubmitedAt(timejst.Now()).
		Save(ctx)
	go func() {

	}()

	stream, err := client.GetSubmit(ctx, &pb.GetSubmitRequest{SubmitId: int32(submit.ID)})
	if err != nil {
		t.Fatal(err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}

		for _, taskRes := range msg.Submit.TaskResults {
			dbtaskRes, err := entClient.TaskResult.Get(ctx, int(taskRes.Id))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, dbtaskRes.RequestPerSec, int(taskRes.RequestPerSec))
		}

		time.Sleep(time.Second)
	}

	submit, _ = entClient.Submit.Get(ctx, submit.ID)
	assert.NotEmpty(t, submit.Score)
	assert.NotNil(t, submit.CompletedAt)
}
