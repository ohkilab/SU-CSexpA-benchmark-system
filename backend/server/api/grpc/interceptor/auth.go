package interceptor

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var claimsKey struct{}
var excludeMethodSet = map[string]struct{}{
	backend.BackendService_PostLogin_FullMethodName:                   {},
	backend.BackendService_GetSubmit_FullMethodName:                   {},
	backend.BackendService_ListContests_FullMethodName:                {},
	backend.BackendService_VerifyToken_FullMethodName:                 {},
	backend.HealthcheckService_PingUnary_FullMethodName:               {},
	backend.HealthcheckService_PingServerSideStreaming_FullMethodName: {},
}

// authentication with JWT
func Auth(secret []byte) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
		method, ok := grpc.Method(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "failed to get method name from context")
		}
		// skip authentication
		if _, ok := excludeMethodSet[method]; ok {
			return ctx, nil
		}

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		claims, err := auth.GetClaimsFromToken(token, secret)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "failed to parse jwt")
		}
		return context.WithValue(ctx, claimsKey, claims), nil
	})
}

func GetClaimsFromContext(ctx context.Context) *auth.Claims {
	return ctx.Value(claimsKey).(*auth.Claims)
}
