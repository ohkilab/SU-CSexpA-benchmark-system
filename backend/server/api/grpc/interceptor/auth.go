package interceptor

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var claimsKey struct{}

// authentication with JWT
func Auth(secret string) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
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
