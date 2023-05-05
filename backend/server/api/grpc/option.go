package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
)

type option struct {
	entClient *ent.Client
	jwtSecret []byte
	worker    worker.Worker
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

func WithWorker(worker worker.Worker) OptionFunc {
	return func(o *option) {
		o.worker = worker
	}
}
