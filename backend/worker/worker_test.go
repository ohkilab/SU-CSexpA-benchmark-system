package worker

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/migrate"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type blockingBenchmarkServer struct {
	benchmarkpb.UnimplementedBenchmarkServiceServer
	responses []*benchmarkpb.ExecuteResponse
}

func (s *blockingBenchmarkServer) CheckConnection(context.Context, *benchmarkpb.CheckConnectionRequest) (*benchmarkpb.CheckConnectionResponse, error) {
	return &benchmarkpb.CheckConnectionResponse{Ok: true}, nil
}

func (s *blockingBenchmarkServer) Execute(_ *benchmarkpb.ExecuteRequest, stream benchmarkpb.BenchmarkService_ExecuteServer) error {
	for _, resp := range s.responses {
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func TestRunBenchmarkTaskMarksTimeoutWhenOverallTimeoutExpires(t *testing.T) {
	ctx := context.Background()
	entClient := openWorkerTestEntClient(ctx, t)
	defer cleanupWorkerTestEntClient(ctx, t, entClient)

	task := testWorkerTask(ctx, t, entClient)
	client, cleanup := launchBlockingBenchmarkServer(t, nil)
	defer cleanup()

	w := New(entClient, client, slog.Default(), WithTimeout(50*time.Millisecond))
	require.NoError(t, w.runBenchmarkTask(task))

	submit, err := entClient.Submit.Get(ctx, task.SubmitID)
	require.NoError(t, err)
	require.Equal(t, backendpb.Status_TIMEOUT.String(), submit.Status)
	require.Equal(t, 0, submit.Score)
}

func TestRunBenchmarkTaskKeepsPositiveScoreWhenOverallTimeoutExpiresAfterSuccess(t *testing.T) {
	ctx := context.Background()
	entClient := openWorkerTestEntClient(ctx, t)
	defer cleanupWorkerTestEntClient(ctx, t, entClient)

	task := testWorkerTask(ctx, t, entClient)
	client, cleanup := launchBlockingBenchmarkServer(t, []*benchmarkpb.ExecuteResponse{
		{
			Ok:                true,
			RequestsPerSecond: 7,
			TotalRequests:     1,
			Task:              task.Req.Tasks[0],
			Status:            backendpb.Status_SUCCESS,
		},
	})
	defer cleanup()

	w := New(entClient, client, slog.Default(), WithTimeout(50*time.Millisecond))
	require.NoError(t, w.runBenchmarkTask(task))

	submit, err := entClient.Submit.Get(ctx, task.SubmitID)
	require.NoError(t, err)
	require.Equal(t, backendpb.Status_SUCCESS.String(), submit.Status)
	require.Equal(t, 7, submit.Score)
}

func TestFinalBenchmarkStatus(t *testing.T) {
	tests := []struct {
		name       string
		current    backendpb.Status
		scores     []int
		hasTimeout bool
		wantStatus backendpb.Status
		wantScore  int
	}{
		{
			name:       "overall timeout without positive scores becomes timeout",
			current:    backendpb.Status_SUCCESS,
			hasTimeout: true,
			wantStatus: backendpb.Status_TIMEOUT,
			wantScore:  0,
		},
		{
			name:       "overall timeout after positive score stays success",
			current:    backendpb.Status_SUCCESS,
			scores:     []int{7},
			hasTimeout: true,
			wantStatus: backendpb.Status_SUCCESS,
			wantScore:  7,
		},
		{
			name:       "validation error still wins over timeout",
			current:    backendpb.Status_VALIDATION_ERROR,
			scores:     []int{7},
			hasTimeout: true,
			wantStatus: backendpb.Status_VALIDATION_ERROR,
			wantScore:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, score := finalBenchmarkStatus(tt.current, tt.scores, tt.hasTimeout)
			require.Equal(t, tt.wantStatus, status)
			require.Equal(t, tt.wantScore, score)
		})
	}
}

func launchBlockingBenchmarkServer(t *testing.T, responses []*benchmarkpb.ExecuteResponse) (benchmarkpb.BenchmarkServiceClient, func()) {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := grpc.NewServer()
	benchmarkpb.RegisterBenchmarkServiceServer(server, &blockingBenchmarkServer{responses: responses})
	go func() {
		_ = server.Serve(listener)
	}()

	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	return benchmarkpb.NewBenchmarkServiceClient(conn), func() {
		server.Stop()
		_ = conn.Close()
		_ = listener.Close()
	}
}

func testWorkerTask(ctx context.Context, t *testing.T, entClient *ent.Client) *Task {
	t.Helper()

	submit, err := entClient.Submit.Create().
		SetURL("http://example.com").
		SetStatus(backendpb.Status_WAITING.String()).
		SetTaskNum(1).
		SetSubmitedAt(time.Now()).
		Save(ctx)
	require.NoError(t, err)

	task := &benchmarkpb.Task{
		Request: &benchmarkpb.HttpRequest{
			Url:         "http://example.com?tag=test",
			Method:      benchmarkpb.HttpMethod_GET,
			ContentType: "application/x-www-form-urlencoded",
		},
		ThreadNum:    1,
		AttemptCount: 1,
	}
	return &Task{
		Req: &benchmarkpb.ExecuteRequest{
			Tasks:            []*benchmarkpb.Task{task},
			GroupId:          "1",
			ContestSlug:      "test-contest",
			Validator:        backendpb.Validator_V2023,
			TimeLimitPerTask: int64(time.Second),
		},
		SubmitID: submit.ID,
		GroupID:  1,
	}
}

func openWorkerTestEntClient(ctx context.Context, t *testing.T) *ent.Client {
	t.Helper()

	if os.Getenv("TEST_MYSQL_PORT") == "" || os.Getenv("MYSQL_DATABASE") == "" {
		t.Skip("TEST_MYSQL_PORT and MYSQL_DATABASE are required for worker tests")
	}
	dsn := fmt.Sprintf(
		"root:%s@tcp(localhost:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("TEST_MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)

	entClient := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
	require.NoError(t, entClient.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)))
	return entClient
}

func cleanupWorkerTestEntClient(ctx context.Context, t *testing.T, entClient *ent.Client) {
	t.Helper()

	_, err := entClient.TaskResult.Delete().Exec(ctx)
	require.NoError(t, err)
	_, err = entClient.Submit.Delete().Exec(ctx)
	require.NoError(t, err)
	require.NoError(t, entClient.Close())
}
