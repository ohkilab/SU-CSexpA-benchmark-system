package service

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
)

type service struct {
	client       *benchmark.Client
	validatorMap map[string]validation.Validator
}

func New(client *benchmark.Client, validatorMap map[string]validation.Validator) pb.BenchmarkServiceServer {
	return &service{client, validatorMap}
}
