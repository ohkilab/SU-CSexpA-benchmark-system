package main

import (
	"context"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/config"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"golang.org/x/exp/slog"
	pkggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// config from env
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	// mysql client
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer entClient.Close()
	// migration
	if err := entClient.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	// logger
	logger := slog.Default()

	// benchmark worker
	conn, err := pkggrpc.Dial(
		fmt.Sprintf("%s:%s", config.BenchmarkHost, config.BenchmarkPort),
		pkggrpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	benchmarkClient := benchmarkpb.NewBenchmarkServiceClient(conn)
	benchmarkWorker := worker.New(entClient, benchmarkClient, logger)
	go benchmarkWorker.Run()

	// build grpc server
	grpcServer, err := grpc.NewServer(
		context.Background(),
		grpc.WithEntClient(entClient),
		grpc.WithJwtSecret(config.JwtSecret),
		grpc.WithWorker(benchmarkWorker),
		grpc.WithLogger(logger),
		grpc.WithTagRepository(tag.NewRespository(".")),
		grpc.UseLogMiddleware(),
	)
	if err != nil {
		panic(err)
	}

	// launch grpc server
	listener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		panic(err)
	}

	slog.Info("server launched")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
