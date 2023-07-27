package e2e

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
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
	group, err := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string("hoge")).
		SetRole(group.RoleContestant).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	require.NoError(t, err)
	contest := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	require.NoError(t, err)
	submit, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetStatus(pb.Status_SUCCESS.String()).
		SetTaskNum(50).
		SetSubmitedAt(timejst.Now()).
		SetCompletedAt(timejst.Now()).
		Save(ctx)
	require.NoError(t, err)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.CreatedAt.Year())
	require.NoError(t, err)
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

func Test_GetContestantInfo(t *testing.T) {
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
	group := utils.CreateGroup(ctx, t, entClient, "test", "hoge", group.RoleContestant)
	contest := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, group)
	time.Sleep(time.Second)
	submit2 := utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, group)
	utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_CONNECTION_FAILED.String(), contest, group)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.CreatedAt.Year())
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	resp, err := client.GetContestantInfo(ctx, &pb.GetContestantInfoRequest{
		ContestSlug: contest.Slug,
	})
	require.NoError(t, err)
	assert.Equal(t, int32(submit2.ID), resp.LatestSubmit.Id)
	assert.Equal(t, int32(9997), resp.RemainingSubmitCount)
}
