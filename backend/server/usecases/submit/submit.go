package submit

import (
	"context"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
}

func NewInteractor(entClient *ent.Client) *Interactor {
	return &Interactor{entClient}
}

func (i *Interactor) PostSubmit(ctx context.Context, req *pb.PostSubmitRequest) (*pb.PostSubmitResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)
	exists, err := i.entClient.Contest.Query().Where(contest.ID(int(req.ContestId))).Exist(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "no such contest")
	}
	submit, err := i.entClient.Submit.Create().
		SetIPAddr(req.IpAddr).
		SetYear(claims.Year).
		SetSubmitedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.PostSubmitResponse{
		Id:         int32(submit.ID),
		IpAddr:     submit.IPAddr,
		ContestId:  req.ContestId,
		SubmitedAt: timestamppb.New(submit.SubmitedAt),
	}, nil
}

func (i *Interactor) GetSubmit(req *pb.GetSubmitRequest, stream pb.BackendService_GetSubmitServer) error {
	s, err := i.entClient.Submit.Get(stream.Context(), int(req.SubmitId))
	if err != nil {
		if ent.IsNotFound(err) {
			return status.Error(codes.NotFound, "no such submit")
		}
		return status.Error(codes.Internal, err.Error())
	}

	for {
		time.Sleep(2 * time.Second)

		s, err = i.entClient.Submit.Get(stream.Context(), int(req.SubmitId))
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		// ベンチマーク処理が完了していたら結果を返す
		if !s.CompletedAt.IsZero() {
			break
		}
		if err := stream.Send(toGetSubmitResponse(s)); err != nil {
			return err
		}
	}

	return nil
}

func toGetSubmitResponse(submit *ent.Submit) *pb.GetSubmitResponse {
	var completedAt *timestamppb.Timestamp
	if !submit.CompletedAt.IsZero() {
		completedAt = timestamppb.New(submit.CompletedAt)
	}
	return &pb.GetSubmitResponse{
		Submit: &pb.Submit{
			Id:      int32(submit.ID),
			GroupId: int32(submit.Edges.Groups[0].ID),
			Year:    int32(submit.Year),
			Score:   int32(submit.Score),
			// Language: submit.Language,
			SubmitedAt:  timestamppb.New(submit.SubmitedAt),
			CompletedAt: completedAt,
			TaskResults: lo.Map(submit.Edges.TaskResults, func(taskResult *ent.TaskResult, _ int) *pb.TaskResult {
				var requestBody *string
				if taskResult.RequestBody != "" {
					requestBody = &taskResult.RequestBody
				}
				return &pb.TaskResult{
					Id:                  int32(taskResult.ID),
					RequestPerSec:       int32(taskResult.RequestPerSec),
					Url:                 taskResult.URL,
					Method:              taskResult.Method,
					RequestContentType:  taskResult.RequestContentType,
					RequestBody:         requestBody,
					ResponseCode:        taskResult.ResponseCode,
					ResponseContentType: taskResult.ResponseContentType,
					ResponseBody:        taskResult.ResponseBody,
					ThreadNum:           int32(taskResult.ThreadNum),
					AttemptCount:        int32(taskResult.AttemptCount),
					AttemptTime:         int32(taskResult.AttemptTime),
					CreatedAt:           timestamppb.New(taskResult.CreatedAt),
				}
			}),
		},
	}
}
