package main

import (
	"log"
	"net"
	"os"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := benchmark.NewClient()

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
	benchmarkService := service.New(client)
	reflection.Register(grpcServer)
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, benchmarkService)

	if err := grpcServer.Serve(lsnr); err != nil {
		log.Fatal(err)
	}
}
