package auth

import (
	"context"
	"log"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
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
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to generateFromPassword")
	}
	group, err := i.entClient.Group.Query().Where(group.NameEQ(id), group.EncryptedPasswordEQ(string(encryptedPassword))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.Unauthenticated, "id or password is incorrect")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	jwtToken, err := auth.GenerateJWTToken(i.secret, group.ID, group.Year)
	if err != nil {
		return nil, err
	}
	return &backend.PostLoginResponse{
		Group: &backend.Group{
			Id:    group.Name,
			Year:  int32(group.Year),
			Role:  backend.Role(backend.Role_value[group.Role.String()]),
			Score: int32(group.Score),
		},
		Token: jwtToken,
	}, nil
}
