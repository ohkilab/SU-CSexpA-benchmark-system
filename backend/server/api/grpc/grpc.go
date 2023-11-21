package grpc

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	interfaces "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
		b, err := bcrypt.GenerateFromPassword([]byte(opt.initAdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to generate password")
		}
		group, err := opt.entClient.Group.Create().
			SetName(opt.initAdminName).
			SetEncryptedPassword(string(b)).
			SetRole(backend.Role_ADMIN.String()).
			SetYear(2023).
			SetCreatedAt(timejst.Now()).
			Save(ctx)
		if err != nil {
			if !ent.IsConstraintError(err) {
				return nil, err
			}
		}
		if group != nil {
			log.Println("initial admin account created", group)
		} else {
			log.Println("skipped creatint initial admin account")
		}
	}

	return grpcServer, nil
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
