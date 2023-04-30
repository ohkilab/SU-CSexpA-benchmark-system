package auth

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInteractor struct {
	secret    []byte
	entClient *ent.Client
}

func NewInteractor(secret []byte, entClient *ent.Client) *AuthInteractor {
	return &AuthInteractor{secret, entClient}
}

func (i *AuthInteractor) PostLogin(ctx context.Context, id, password string) (*backend.PostLoginResponse, error) {
	group, err := i.entClient.Group.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := bcrypt.CompareHashAndPassword([]byte(group.EncryptedPassword), []byte(password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "password mismatch")
	}
	jwtToken, err := auth.GenerateJWTToken(i.secret, group.ID, group.Year)
	if err != nil {
		return nil, err
	}
	return &backend.PostLoginResponse{
		Group: &backend.Group{
			Id:    group.ID,
			Year:  int32(group.Year),
			Role:  backend.Role(backend.Role_value[group.Role.String()]),
			Score: int32(group.Score),
		},
		Token: jwtToken,
	}, nil
}
