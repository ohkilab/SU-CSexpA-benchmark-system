package service

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
)

type service struct {
	client *benchmark.Client
	pb.UnimplementedBenchmarkServiceServer
}

func New(client *benchmark.Client) pb.BenchmarkServiceServer {
	return &service{client, pb.UnimplementedBenchmarkServiceServer{}}
}
