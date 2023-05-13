package main

import (
	"context"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/config"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"golang.org/x/exp/slog"
	pkgrpc "google.golang.org/grpc"
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
	// migration
	if err := entClient.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	// benchmark worker
	conn, err := pkgrpc.Dial(fmt.Sprintf("%s:%s", config.BenchmarkHost, config.BenchmarkPort), pkgrpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	benchmarkClient := benchmark.NewBenchmarkServiceClient(conn)
	benchmarkWorker := worker.New(entClient, benchmarkClient)
	go benchmarkWorker.Run()

	// build grpc server
	grpcServer := grpc.NewServer(grpc.WithEntClient(entClient), grpc.WithJwtSecret(config.JwtSecret), grpc.WithWorker(benchmarkWorker))

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
