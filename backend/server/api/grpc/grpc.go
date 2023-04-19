package grpc

import (
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc"
)

func NewServer(optionFuncs ...optionFunc) *grpc.Server {
	grpcServer := grpc.NewServer()
	opt := &option{}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	backendService := interfaces.NewBackendService()
	backend.RegisterBackendServiceServer(grpcServer, backendService)

	return grpcServer
}
