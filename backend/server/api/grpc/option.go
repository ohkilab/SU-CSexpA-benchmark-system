package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
)

type option struct {
	entClient *ent.Client
	jwtSecret string
}

type optionFunc func(*option)

func WithEntClient(entClient *ent.Client) optionFunc {
	return func(o *option) {
		o.entClient = entClient
	}
}

func WithJwtSecret(jwtSecret string) optionFunc {
	return func(o *option) {
		o.jwtSecret = jwtSecret
	}
}
