package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
)

type option struct {
	entClient *ent.Client
}

type optionFunc func(*option)

func WithEntClient(entClient *ent.Client) optionFunc {
	return func(o *option) {
		o.entClient = entClient
	}
}
