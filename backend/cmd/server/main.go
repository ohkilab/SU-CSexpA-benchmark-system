package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/config"
)

func main() {
	// config from env
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// mysql client
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// build grpc server
	grpcServer := grpc.NewServer(grpc.WithEntClient(entClient))

	// launch grpc server
	listener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
