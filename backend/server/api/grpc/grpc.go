package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(optionFuncs ...optionFunc) *grpc.Server {
	opt := &option{}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Auth(opt.jwtSecret)),
	)

	backendService := interfaces.NewBackendService(opt.jwtSecret, opt.entClient)
	backend.RegisterBackendServiceServer(grpcServer, backendService)
	healthcheckService := interfaces.NewHealthcheckService()
	backend.RegisterHealthcheckServiceServer(grpcServer, healthcheckService)
	reflection.Register(grpcServer)

	return grpcServer
}
