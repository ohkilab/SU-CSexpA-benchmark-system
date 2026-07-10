package test

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type passValidator struct{}

func (passValidator) Validate(*url.URL, []byte) error {
	return nil
}

func TestExecuteSendsTimeoutAndContinuesToNextTask(t *testing.T) {
	canceled := make(chan struct{})
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			closeOnce(canceled)
			return
		case <-time.After(500 * time.Millisecond):
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer slowServer.Close()

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer fastServer.Close()

	client, cleanup := launchBenchmarkServiceForTimeoutTest(t)
	defer cleanup()

	responses := executeAndCollect(t, client, &pb.ExecuteRequest{
		Tasks: []*pb.Task{
			testExecuteTask(slowServer.URL, 1),
			testExecuteTask(fastServer.URL, 1),
		},
		GroupId:          "group",
		Validator:        backendpb.Validator_V2023,
		TimeLimitPerTask: int64(50 * time.Millisecond),
	})
	require.Len(t, responses, 2)

	byURL := map[string]*pb.ExecuteResponse{}
	for _, resp := range responses {
		byURL[resp.Task.Request.Url] = resp
	}
	require.Equal(t, backendpb.Status_TIMEOUT, byURL[slowServer.URL].Status)
	require.False(t, byURL[slowServer.URL].Ok)
	require.Equal(t, int32(0), byURL[slowServer.URL].RequestsPerSecond)
	require.Equal(t, int32(0), byURL[slowServer.URL].TotalRequests)
	require.Equal(t, backendpb.Status_SUCCESS, byURL[fastServer.URL].Status)
	require.True(t, byURL[fastServer.URL].RequestsPerSecond > 0)

	select {
	case <-canceled:
	case <-time.After(time.Second):
		t.Fatal("slow request was not canceled by task timeout")
	}
}

func TestExecuteKeepsPositivePartialScoreWhenTaskTimesOut(t *testing.T) {
	var count int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&count, 1) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		<-r.Context().Done()
	}))
	defer server.Close()

	client, cleanup := launchBenchmarkServiceForTimeoutTest(t)
	defer cleanup()

	responses := executeAndCollect(t, client, &pb.ExecuteRequest{
		Tasks: []*pb.Task{
			testExecuteTask(server.URL, 2),
		},
		GroupId:          "group",
		Validator:        backendpb.Validator_V2023,
		TimeLimitPerTask: int64(50 * time.Millisecond),
	})
	require.Len(t, responses, 1)
	require.Equal(t, backendpb.Status_SUCCESS, responses[0].Status)
	require.True(t, responses[0].Ok)
	require.Equal(t, int32(1), responses[0].TotalRequests)
	require.True(t, responses[0].RequestsPerSecond > 0)
}

func executeAndCollect(t *testing.T, client pb.BenchmarkServiceClient, req *pb.ExecuteRequest) []*pb.ExecuteResponse {
	t.Helper()

	stream, err := client.Execute(context.Background(), req)
	require.NoError(t, err)

	responses := make([]*pb.ExecuteResponse, 0, len(req.Tasks))
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return responses
		}
		require.NoError(t, err)
		responses = append(responses, resp)
	}
}

func launchBenchmarkServiceForTimeoutTest(t *testing.T) (pb.BenchmarkServiceClient, func()) {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	grpcServer := grpc.NewServer()
	validatorMap := map[backendpb.Validator]validation.Validator{
		backendpb.Validator_V2023: passValidator{},
	}
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, service.New(benchmark.NewClient(), validatorMap))
	go func() {
		_ = grpcServer.Serve(listener)
	}()

	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	return pb.NewBenchmarkServiceClient(conn), func() {
		grpcServer.Stop()
		_ = conn.Close()
		_ = listener.Close()
	}
}

func testExecuteTask(rawURL string, attemptCount int32) *pb.Task {
	return &pb.Task{
		Request: &pb.HttpRequest{
			Url:         rawURL,
			Method:      pb.HttpMethod_GET,
			ContentType: "application/x-www-form-urlencoded",
		},
		ThreadNum:    1,
		AttemptCount: attemptCount,
	}
}

func closeOnce(ch chan struct{}) {
	select {
	case <-ch:
	default:
		close(ch)
	}
}
