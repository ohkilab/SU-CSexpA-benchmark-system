package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type Claims struct {
	GroupID string `json:"group_id"`
	jwt.RegisteredClaims
}

func GetClaimsFromToken(token, secret string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("invalid jwt token")
	}
	claims, _ := jwtToken.Claims.(*Claims)
	if claims == nil {
		return nil, errors.New("invalid jwt token")
	}
	return claims, nil
}
