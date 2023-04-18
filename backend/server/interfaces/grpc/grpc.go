package grpc

import (
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type option struct {
	db *gorm.DB
}

type optionFunc func(*option)

func WithDB(db *gorm.DB) optionFunc {
	return func(o *option) {
		o.db = db
	}
}

func NewServer(optionFuncs ...optionFunc) *grpc.Server {
	grpcServer := grpc.NewServer()
	opt := &option{}
	for _, optionFunc := range optionFuncs {
		optionFunc(opt)
	}

	backendService := newBackendService()
	backend.RegisterBackendServiceServer(grpcServer, backendService)

	return grpcServer
}
