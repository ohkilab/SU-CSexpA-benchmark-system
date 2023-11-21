package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInteractor struct {
	secret    []byte
	entClient *ent.Client
	logger    *slog.Logger
}

func NewInteractor(secret []byte, entClient *ent.Client, logger *slog.Logger) *AuthInteractor {
	return &AuthInteractor{secret, entClient, logger}
}

func (i *AuthInteractor) PostLogin(ctx context.Context, id, password string) (*pb.PostLoginResponse, error) {
	groups, err := i.entClient.Group.Query().WithSubmits().Where(group.NameEQ(id)).All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch groups", "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch groups")
	}
	var group *ent.Group
	for _, g := range groups {
		err := bcrypt.CompareHashAndPassword([]byte(g.EncryptedPassword), []byte(password))
		if err == nil {
			group = g
			break
		}
	}
	if group == nil {
		return nil, status.Error(codes.Unauthenticated, "id or password is incorrect")
	}
	jwtToken, err := auth.GenerateJWTToken(i.secret, group.ID, group.CreatedAt.Year(), group.Role)
	if err != nil {
		i.logger.Error("failed to generate jwt token", "error", err)
		return nil, err
	}
	return &pb.PostLoginResponse{
		Group: &pb.Group{
			Name: group.Name,
			Role: pb.Role(pb.Role_value[group.Role]),
		},
		Token: jwtToken,
	}, nil
}

func (i *AuthInteractor) VerifyToken(ctx context.Context) *pb.VerifyTokenResponse {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		i.logger.Error("failed to get token from metadata", "error", err)
		return &pb.VerifyTokenResponse{
			Ok:      false,
			Message: "missing token",
		}
	}
	claims, err := auth.GetClaimsFromToken(token, i.secret)
	if err != nil {
		i.logger.Error("failed to get claims from token", "error", err)
		return &pb.VerifyTokenResponse{
			Ok:      false,
			Message: "invalid token",
		}
	}
	if timejst.Now().After(claims.ExpiresAt.Time) {
		return &pb.VerifyTokenResponse{
			Ok:      false,
			Message: "token expired",
		}
	}
	return &pb.VerifyTokenResponse{
		Ok:      true,
		Message: "ok",
	}
}
