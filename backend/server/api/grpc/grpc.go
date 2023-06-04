package grpc

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(optionFuncs ...OptionFunc) *grpc.Server {
	opt := &option{
		logger: slog.Default(),
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Auth(opt.jwtSecret)),
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(opt.logger)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(opt.logger)),
		),
	)

	backendService := interfaces.NewBackendService(opt.jwtSecret, opt.entClient, opt.worker, opt.logger, opt.tagRepository)
	backend.RegisterBackendServiceServer(grpcServer, backendService)
	healthcheckService := interfaces.NewHealthcheckService()
	backend.RegisterHealthcheckServiceServer(grpcServer, healthcheckService)
	reflection.Register(grpcServer)

	return grpcServer
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
