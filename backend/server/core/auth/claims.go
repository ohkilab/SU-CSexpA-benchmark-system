package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/pkg/errors"
)

type Claims struct {
	GroupID int    `json:"group_id"`
	Year    int    `json:"year"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

func GetClaimsFromToken(token string, secret []byte) (*Claims, error) {
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

func GenerateJWTToken(secret []byte, groupID, year int, role string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		GroupID: groupID,
		Year:    year,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timejst.Now().Add(7 * 24 * time.Hour)),
		},
	})
	return jwtToken.SignedString(secret)
}
