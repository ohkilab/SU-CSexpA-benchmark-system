package submit

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/benchmark"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	threadNum    int = 5
	attemptCount int = 100
)

type Interactor struct {
	entClient *ent.Client
	worker    worker.Worker
}

func NewInteractor(entClient *ent.Client, worker worker.Worker) *Interactor {
	return &Interactor{entClient, worker}
}

func (i *Interactor) PostSubmit(ctx context.Context, req *backendpb.PostSubmitRequest) (*backendpb.PostSubmitResponse, error) {
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
		SetContestsID(int(req.ContestId)).
		SetGroupsID(claims.GroupID).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// add a task to worker
	executeRequest := buildExecuteRequest(submit.IPAddr, strconv.Itoa(claims.GroupID))
	i.worker.Push(executeRequest)

	return &backendpb.PostSubmitResponse{
		Id:         int32(submit.ID),
		IpAddr:     submit.IPAddr,
		ContestId:  req.ContestId,
		SubmitedAt: timestamppb.New(submit.SubmitedAt),
	}, nil
}

func buildExecuteRequest(ipAddr, groupID string) *benchmarkpb.ExecuteRequest {
	tags := generateRandomTags(50)
	return &benchmarkpb.ExecuteRequest{
		GroupId: groupID,
		Tasks: lo.Map(tags, func(tag string, _ int) *benchmarkpb.Task {
			return &benchmarkpb.Task{
				Request: &benchmarkpb.HttpRequest{
					Url:         fmt.Sprintf("http://%s?tag=%s", ipAddr, tag),
					Method:      benchmarkpb.HttpMethod_GET,
					ContentType: "application/x-www-form-urlencoded",
					Body:        "",
				},
				ThreadNum:    int32(threadNum),
				AttemptCount: int32(attemptCount),
			}
		}),
	}
}

// TODO: flickr の tag api を叩いて取るとか？
func generateRandomTags(n int) []string {
	tags := make([]string, n)
	for i := range tags {
		tags[i] = lo.RandomString(10, lo.AlphanumericCharset)
	}
	return tags
}

func (i *Interactor) GetSubmit(req *backendpb.GetSubmitRequest, stream backendpb.BackendService_GetSubmitServer) error {
	for {
		s, err := i.entClient.Submit.Query().Where(submit.IDEQ(int(req.SubmitId))).WithGroups().WithTaskResults().Only(stream.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return status.Error(codes.NotFound, "no such submit")
			}
			return status.Error(codes.Internal, err.Error())
		}
		if err := stream.Send(toGetSubmitResponse(s)); err != nil {
			return err
		}
		// ベンチマーク処理が完了していたら結果を返す
		if !s.CompletedAt.IsZero() {
			break
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}

func toGetSubmitResponse(submit *ent.Submit) *backendpb.GetSubmitResponse {
	var completedAt *timestamppb.Timestamp
	if !submit.CompletedAt.IsZero() {
		completedAt = timestamppb.New(submit.CompletedAt)
	}
	return &backendpb.GetSubmitResponse{
		Submit: &backendpb.Submit{
			Id:      int32(submit.ID),
			GroupId: int32(submit.Edges.Groups.ID),
			Year:    int32(submit.Year),
			Score:   int32(submit.Score),
			// Language: submit.Language,
			SubmitedAt:  timestamppb.New(submit.SubmitedAt),
			CompletedAt: completedAt,
			TaskResults: lo.Map(submit.Edges.TaskResults, func(taskResult *ent.TaskResult, _ int) *backendpb.TaskResult {
				var requestBody *string
				if taskResult.RequestBody != "" {
					requestBody = &taskResult.RequestBody
				}
				return &backendpb.TaskResult{
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
					CreatedAt:           timestamppb.New(taskResult.CreatedAt),
				}
			}),
		},
	}
}
