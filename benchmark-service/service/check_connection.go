package service

import (
	"context"
	"log"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
)

func (s *service) CheckConnection(ctx context.Context, req *pb.CheckConnectionRequest) (*pb.CheckConnectionResponse, error) {
	if err := s.client.CheckConnection(req.Url); err != nil {
		log.Println(err)
		return &pb.CheckConnectionResponse{
			Ok:           false,
			ErrorMessage: toPtr(err.Error()),
		}, nil
	}
	return &pb.CheckConnectionResponse{
		Ok: true,
	}, nil
}

func toPtr[T any](t T) *T {
	return &t
}
