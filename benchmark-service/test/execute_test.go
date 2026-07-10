package test

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/test/utils"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_Execute(t *testing.T) {
	port := utils.LaunchTestServer(t)

	lsnr, err := net.Listen("tcp", ":3776")
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()

	validatorMap := map[backendpb.Validator]validation.Validator{
		backendpb.Validator_V2023: passValidator{},
	}

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, service.New(benchmark.NewClient(), validatorMap))
	go func() {
		if err := grpcServer.Serve(lsnr); err != nil {
			panic(err)
		}
	}()
	defer grpcServer.GracefulStop()

	conn, err := grpc.Dial("0.0.0.0:3776", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewBenchmarkServiceClient(conn)
	// TODO: mock for validator
	stream, err := client.Execute(context.Background(), &pb.ExecuteRequest{
		Tasks: []*pb.Task{
			testTask(port),
			testTask(port),
			testTask(port),
			testTask(port),
			testTask(port),
		},
		GroupId:          "hoge",
		Validator:        backendpb.Validator_V2023,
		TimeLimitPerTask: int64(5 * time.Second),
	})
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
		t.Log(msg)
	}
}

func testTask(port int) *pb.Task {
	return &pb.Task{
		Request: &pb.HttpRequest{
			Url:         "http://0.0.0.0:" + strconv.Itoa(port),
			Method:      pb.HttpMethod_GET,
			ContentType: "x-www-form-urlencoded",
			Body:        "",
		},
		ThreadNum:    1,
		AttemptCount: 1,
	}
}
