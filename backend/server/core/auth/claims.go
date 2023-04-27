package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type Claims struct {
	GroupID string `json:"group_id"`
	Year    int    `json:"year"`
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

func GenerateJWTToken(secret, groupID string, year int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		GroupID: groupID,
		Year:    year,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	})
	return jwtToken.SignedString(secret)
}
