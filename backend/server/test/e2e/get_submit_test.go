package e2e

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Test_GetSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	connBenchmark, closeBenchmark := utils.LaunchBenchmarkGrpcServer(t)
	defer closeBenchmark()
	benchmarkClient := benchmarkpb.NewBenchmarkServiceClient(connBenchmark)
	worker := worker.New(entClient, benchmarkClient, slog.Default())
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
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	contest, _ := entClient.Contest.Create().
		SetTitle("test contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetYear(2023).
		SetSubmitLimit(9999).
		SetCreatedAt(timejst.Now()).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		Save(ctx)
	submit, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetStatus(pb.Status_SUCCESS.String()).
		SetYear(2023).
		SetTaskNum(50).
		SetSubmitedAt(timejst.Now()).
		SetCompletedAt(timejst.Now()).
		Save(ctx)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.Year)
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	stream, err := client.GetSubmit(ctx, &pb.GetSubmitRequest{SubmitId: int32(submit.ID)})
	require.NoError(t, err)
	var ok bool
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		ok = true

		for _, taskRes := range msg.Submit.TaskResults {
			dbtaskRes, err := entClient.TaskResult.Get(ctx, int(taskRes.Id))
			require.NoError(t, err)
			t.Log(taskRes.Url, dbtaskRes.ErrorMessage)
			assert.Equal(t, dbtaskRes.RequestPerSec, int(taskRes.RequestPerSec))
		}

		time.Sleep(time.Second)
	}
	assert.True(t, ok)
	dbSubmit, err := entClient.Submit.Get(ctx, int(submit.ID))
	require.NoError(t, err)
	assert.NotNil(t, dbSubmit)

	// not found
	stream, err = client.GetSubmit(ctx, &pb.GetSubmitRequest{SubmitId: 999})
	require.NoError(t, err)
	_, err = stream.Recv()
	assert.Equal(t, codes.NotFound, status.Code(err))
}
