package main

import (
	"log"
	"net"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"google.golang.org/grpc"
)

func main() {
	client := benchmark.NewClient()

	lsnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()

	grpcServer := grpc.NewServer()
	benchmarkService := service.New(client)
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, benchmarkService)

	if err := grpcServer.Serve(lsnr); err != nil {
		log.Fatal(err)
	}
}
