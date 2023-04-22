package healthcheck

import (
	"context"

	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
)

type HealthcheckInteractor struct{}

func NewInteractor() *HealthcheckInteractor {
	return &HealthcheckInteractor{}
}

func (i *HealthcheckInteractor) Ping(ctx context.Context, ping string) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Pong: ping + " Pong:)",
	}, nil
}
