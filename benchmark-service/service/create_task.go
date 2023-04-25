package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/task"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const taskKey = "task"

func (s *service) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	ipAddr := net.ParseIP(req.IpAddr)
	if ipAddr == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid ip address")
	}
	if len(req.GroupId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "groupID must not be empty")
	}
	if len(req.GroupId) > 100 {
		return nil, status.Error(codes.InvalidArgument, "groupID must be 100 or less")
	}

	task, key := task.NewTask(ipAddr.String(), req.GroupId)
	if err := s.taskClient.SetTask(ctx, key, task); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &pb.CreateTaskResponse{
		Id:          fmt.Sprintf("%x", sha256.Sum256([]byte(key))),
		GroupId:     task.GroupID,
		IpAddr:      task.IPAddr,
		SubmittedAt: timestamppb.New(task.CreatedAt),
	}, nil
}
