package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc"
)

func NewServer(optionFuncs ...optionFunc) *grpc.Server {
	opt := &option{}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Auth(opt.jwtSecret)),
	)

	backendService := interfaces.NewBackendService()
	backend.RegisterBackendServiceServer(grpcServer, backendService)

	return grpcServer
}
