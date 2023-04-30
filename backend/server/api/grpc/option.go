package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
)

type option struct {
	entClient *ent.Client
	jwtSecret []byte
}

type OptionFunc func(*option)

func WithEntClient(entClient *ent.Client) OptionFunc {
	return func(o *option) {
		o.entClient = entClient
	}
}

func WithJwtSecret(jwtSecret string) OptionFunc {
	return func(o *option) {
		o.jwtSecret = []byte(jwtSecret)
	}
}
