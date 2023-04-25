package main

import (
	"log"
	"net"
	"os"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/task"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	taskClient := task.NewClient(redisClient)

	lsnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()

	grpcServer := grpc.NewServer()
	benchmarkService := service.New(taskClient)
	grpcServer.RegisterService(&pb.BenchmarkService_ServiceDesc, benchmarkService)

	if err := grpcServer.Serve(lsnr); err != nil {
		log.Fatal(err)
	}
}
