package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	"golang.org/x/exp/slog"
)

type option struct {
	entClient *ent.Client
	jwtSecret []byte
	worker    worker.Worker
	logger    *slog.Logger
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

func WithLogger(logger *slog.Logger) OptionFunc {
	return func(o *option) {
		o.logger = logger
	}
}
