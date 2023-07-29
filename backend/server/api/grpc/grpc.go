package grpc

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(ctx context.Context, optionFuncs ...OptionFunc) (*grpc.Server, error) {
	opt := &option{
		logger: slog.Default(),
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	serverOptions := []grpc.ServerOption{grpc.UnaryInterceptor(interceptor.Auth(opt.jwtSecret))}
	if opt.useLogMiddleware {
		serverOptions = append(serverOptions,
			grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(opt.logger))),
			grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(opt.logger))),
		)
	}
	grpcServer := grpc.NewServer(serverOptions...)

	backendService := interfaces.NewBackendService(opt.jwtSecret, opt.entClient, opt.worker, opt.logger, opt.tagRepository, opt.limit)
	backend.RegisterBackendServiceServer(grpcServer, backendService)
	healthcheckService := interfaces.NewHealthcheckService()
	backend.RegisterHealthcheckServiceServer(grpcServer, healthcheckService)
	adminService := interfaces.NewAdminService(opt.entClient, opt.logger, opt.tagRepository)
	backend.RegisterAdminServiceServer(grpcServer, adminService)
	reflection.Register(grpcServer)

	if opt.initAdminName != "" && opt.initAdminPassword != "" {
		if _, err := adminService.CreateGroups(ctx, &backend.CreateGroupsRequest{
			Groups: []*backend.CreateGroupsRequest_CreateGroupsGroup{
				{
					Name:     opt.initAdminName,
					Password: opt.initAdminPassword,
					Year:     int32(time.Now().Year()),
					Role:     backend.Role_ADMIN,
				},
			},
		}); err != nil {
			if !ent.IsConstraintError(err) {
				return nil, err
			}
		}
	}

	return grpcServer, nil
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
