package e2e

import (
	"context"
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func Test_GetSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	connBenchmark, closeBenchmark := utils.LaunchBenchmarkGrpcServer(t)
	defer closeBenchmark()
	benchmarkClient := benchmarkpb.NewBenchmarkServiceClient(connBenchmark)
	testServerPort := utils.LaunchTestServer(t)
	worker := worker.New(entClient, benchmarkClient)
	go worker.Run()

	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithEntClient(entClient), grpc.WithJwtSecret("secret"), grpc.WithWorker(worker))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	group, _ := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string("hoge")).
		SetRole(group.RoleContestant).
		SetScore(12345).
		SetYear(2023).
		Save(ctx)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetYear(2023).
		SetSubmitLimit(9999).
		SetCreatedAt(time.Now()).
		Save(ctx)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.Year)
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)
	submit, err := client.PostSubmit(ctx, &pb.PostSubmitRequest{
		IpAddr:    "0.0.0.0:" + strconv.Itoa(testServerPort),
		ContestId: int32(contest.ID),
	})
	if err != nil {
		t.Fatal(err)
	}
	stream, err := client.GetSubmit(ctx, &pb.GetSubmitRequest{SubmitId: int32(submit.Id)})
	if err != nil {
		t.Fatal(err)
	}
	var ok bool
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		ok = true

		for _, taskRes := range msg.Submit.TaskResults {
			dbtaskRes, err := entClient.TaskResult.Get(ctx, int(taskRes.Id))
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, dbtaskRes.RequestPerSec, int(taskRes.RequestPerSec))
		}

		time.Sleep(time.Second)
	}
	assert.True(t, ok)

	dbSubmit, err := entClient.Submit.Get(ctx, int(submit.Id))
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, dbSubmit.Score)
	assert.NotNil(t, dbSubmit.CompletedAt)
}
