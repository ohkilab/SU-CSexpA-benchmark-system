package main

import (
	"log"
	"net"
	"os"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := benchmark.NewClient()

	logger := slog.Default()
	validatorMap := validation.NewValidator(logger)

	port := os.Getenv("BENCHMARK_GRPC_PORT")
	if port == "" {
		log.Fatal("BENCHMARK_GRPC_PORT is not set")
	}
	lsnr, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()

	grpcServer := grpc.NewServer()
	benchmarkService := service.New(client, validatorMap)
	reflection.Register(grpcServer)
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, benchmarkService)

	if err := grpcServer.Serve(lsnr); err != nil {
		log.Fatal(err)
	}
}
