package service

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/task"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
)

type service struct {
	taskClient *task.Client
}

func New(taskClient *task.Client) pb.BenchmarkServiceServer {
	return &service{taskClient}
}
