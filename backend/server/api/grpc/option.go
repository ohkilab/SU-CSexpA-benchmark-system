package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	"golang.org/x/exp/slog"
)

type option struct {
	entClient        *ent.Client
	jwtSecret        []byte
	worker           worker.Worker
	logger           *slog.Logger
	tagRepository    tag.Repository
	useLogMiddleware bool
	limit            int
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

func WithTagRepository(r tag.Repository) OptionFunc {
	return func(o *option) {
		o.tagRepository = r
	}
}

func UseLogMiddleware() OptionFunc {
	return func(o *option) {
		o.useLogMiddleware = true
	}
}

func WithLimit(limit int) OptionFunc {
	return func(o *option) {
		o.limit = limit
	}
}
