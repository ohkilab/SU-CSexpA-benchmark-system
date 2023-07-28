package e2e

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	mock_worker "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker/mock"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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
		SetRole(pb.Role_CONTESTANT.String()).
		SetYear(2023).
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

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.CreatedAt.Year(), group.Role)
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

func Test_GetLatestSubmit(t *testing.T) {
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
	group := utils.CreateGroup(ctx, t, entClient, "test", "hoge", 2023, pb.Role_CONTESTANT)
	contest := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, group)
	time.Sleep(time.Second)
	submit2 := utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, group)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group.ID, group.CreatedAt.Year(), group.Role)
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	resp, err := client.GetLatestSubmit(ctx, &pb.GetLatestSubmitRequest{})
	require.NoError(t, err)
	assert.Equal(t, int32(submit2.ID), resp.Submit.Id)
}

func Test_ListSubmits(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithLimit(3))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	now := timejst.Now()
	group1 := utils.CreateGroup(ctx, t, entClient, "test1", "hoge", 2023, pb.Role_CONTESTANT)
	group2 := utils.CreateGroup(ctx, t, entClient, "test2", "hoge", 2023, pb.Role_CONTESTANT)
	c := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)
	submit1, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -6)).
		SetTaskNum(50).
		SetScore(100).
		Save(ctx)
	require.NoError(t, err)
	submit2, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -5)).
		SetTaskNum(50).
		SetScore(200).
		Save(ctx)
	require.NoError(t, err)
	submit3, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group2).
		SetStatus(pb.Status_CONNECTION_FAILED.String()).
		SetSubmitedAt(now.AddDate(0, 0, -4)).
		SetTaskNum(50).
		Save(ctx)
	require.NoError(t, err)
	submit4, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -3)).
		SetTaskNum(50).
		SetScore(300).
		Save(ctx)
	require.NoError(t, err)
	submit5, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group1).
		SetStatus(pb.Status_SUCCESS.String()).
		SetSubmitedAt(now.AddDate(0, 0, -2)).
		SetTaskNum(50).
		SetScore(400).
		Save(ctx)
	require.NoError(t, err)
	submit6, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(c).
		SetGroups(group2).
		SetStatus(pb.Status_CONNECTION_FAILED.String()).
		SetSubmitedAt(now.AddDate(0, 0, -1)).
		SetTaskNum(50).
		Save(ctx)
	require.NoError(t, err)

	jwtToken, err := auth.GenerateJWTToken([]byte("secret"), group1.ID, group1.CreatedAt.Year(), group1.Role)
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	resp, err := client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, int32(2), resp.TotalPages)
	assert.Equal(t, submit6.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit5.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit4.ID, int(resp.Submits[2].Id))
	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        2,
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, submit3.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit2.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit1.ID, int(resp.Submits[2].Id))

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
		GroupName:   lo.ToPtr("test2"),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 2)
	assert.Equal(t, int32(1), resp.TotalPages)
	assert.Equal(t, submit6.ID, int(resp.Submits[0].Id))
	assert.Equal(t, group2.Name, resp.Submits[0].GroupName)
	assert.Equal(t, submit3.ID, int(resp.Submits[1].Id))
	assert.Equal(t, group2.Name, resp.Submits[1].GroupName)

	resp, err = client.ListSubmits(ctx, &pb.ListSubmitsRequest{
		ContestSlug: c.Slug,
		Page:        1,
		SortBy:      pb.ListSubmitsRequest_SCORE.Enum(),
		Status:      pb.Status_SUCCESS.Enum(),
		IsDesc:      lo.ToPtr(false),
	})
	require.NoError(t, err)
	require.Len(t, resp.Submits, 3)
	assert.Equal(t, int32(2), resp.TotalPages)
	assert.Equal(t, submit1.ID, int(resp.Submits[0].Id))
	assert.Equal(t, submit2.ID, int(resp.Submits[1].Id))
	assert.Equal(t, submit4.ID, int(resp.Submits[2].Id))
}

func Test_PostSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	worker := mock_worker.NewMockWorker(ctrl)
	worker.EXPECT().Push(gomock.Any()).AnyTimes()
	secret := []byte("secret")
	mockRepository := tag.MockRepository(
		func(contestSlug string, num int) ([]string, error) {
			return []string{"a", "b", "c"}, nil
		}, nil, nil, nil)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithWorker(worker), grpc.WithTagRepository(mockRepository))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	group, err := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(pb.Role_CONTESTANT.String()).
		SetYear(2023).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	require.NoError(t, err)
	contest := utils.CreateContest(ctx, t, entClient, "test contest", "test-contest", pb.Validator_V2023.String(), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC), 9999, contest.TagSelectionLogicAuto)

	jwtToken, err := auth.GenerateJWTToken(secret, group.ID, group.CreatedAt.Year(), group.Role)
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	req := &pb.PostSubmitRequest{
		Url:         "http://10.255.255.255",
		ContestSlug: contest.Slug,
	}
	resp, err := client.PostSubmit(ctx, req)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.Url, resp.Url)
	assert.Equal(t, req.ContestSlug, resp.ContestSlug)
	assert.NotEmpty(t, resp.SubmitedAt)

	// failed
	req = &pb.PostSubmitRequest{
		Url:         "http://10.255.255.255",
		ContestSlug: "",
	}
	_, err = client.PostSubmit(ctx, req)
	assert.Error(t, err)
}
