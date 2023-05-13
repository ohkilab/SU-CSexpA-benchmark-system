package submit

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
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
		SetURL(req.Url).
		SetYear(claims.Year).
		SetSubmitedAt(timejst.Now()).
		SetContestsID(int(req.ContestId)).
		SetGroupsID(claims.GroupID).
		SetStatus(backendpb.Status_WAITING.String()).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// add a task to worker
	executeRequest := buildTask(submit.URL, strconv.Itoa(claims.GroupID), submit.ID)
	i.worker.Push(executeRequest)

	return &backendpb.PostSubmitResponse{
		Id:         int32(submit.ID),
		Url:        submit.URL,
		ContestId:  req.ContestId,
		SubmitedAt: timestamppb.New(submit.SubmitedAt),
	}, nil
}

func buildTask(url, groupID string, submitID int) *worker.Task {
	tags := generateRandomTags(50)
	return &worker.Task{
		Req: &benchmarkpb.ExecuteRequest{
			GroupId: groupID,
			Tasks: lo.Map(tags, func(tag string, _ int) *benchmarkpb.Task {
				return &benchmarkpb.Task{
					Request: &benchmarkpb.HttpRequest{
						Url:         fmt.Sprintf("%s?tag=%s", url, tag),
						Method:      benchmarkpb.HttpMethod_GET,
						ContentType: "application/x-www-form-urlencoded",
						Body:        "",
					},
					ThreadNum:    int32(threadNum),
					AttemptCount: int32(attemptCount),
				}
			}),
		},
		SubmitID: submitID,
	}
}

// TODO: flickr の tag api を叩いて取るとか？
func generateRandomTags(n int) []string {
	// tags := make([]string, n)
	// for i := range tags {
	// 	tags[i] = lo.RandomString(10, lo.AlphanumericCharset)
	// }
	// return tags

	return lo.Shuffle([]string{
		"陸上自衛隊",
		"陸光麵館",
		"陸前高田ボランティア",
		"陸前高田市",
		"陸家嘴",
		"陸羽茶室",
		"陸航",
		"険道",
		"陽光橋",
		"陽光橋夜景",
		"陽明公園",
		"陽明大學",
		"陽明山",
		"陽明山中國麗緻大飯店",
		"陽明山公園",
		"陽明山前山公園",
		"陽明山國家公園",
	})

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
		if err := stream.Send(&backendpb.GetSubmitResponse{
			Submit: toPbSubmit(s),
		}); err != nil {
			return err
		}
		// ベンチマーク処理が完了していたら結果を返す
		if !s.CompletedAt.IsZero() {
			log.Println("completed submit:", s)
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func toPbSubmit(submit *ent.Submit) *backendpb.Submit {
	return &backendpb.Submit{
		Id:      int32(submit.ID),
		GroupId: int32(submit.Edges.Groups.ID),
		Year:    int32(submit.Year),
		Score:   int32(submit.Score),
		// Language: submit.Language,
		ErrorMessage: &submit.Message,
		SubmitedAt:   timestamppb.New(submit.SubmitedAt),
		CompletedAt:  timestamppb.New(submit.CompletedAt),
		TaskResults: lo.Map(submit.Edges.TaskResults, func(taskResult *ent.TaskResult, _ int) *backendpb.TaskResult {
			var requestBody *string
			if taskResult.RequestBody != "" {
				requestBody = &taskResult.RequestBody
			}
			return &backendpb.TaskResult{
				Id:                 int32(taskResult.ID),
				RequestPerSec:      int32(taskResult.RequestPerSec),
				Url:                taskResult.URL,
				Method:             taskResult.Method,
				RequestContentType: taskResult.RequestContentType,
				RequestBody:        requestBody,
				ThreadNum:          int32(taskResult.ThreadNum),
				AttemptCount:       int32(taskResult.AttemptCount),
				Status:             backendpb.Status(backendpb.Status_value[taskResult.Status]),
				CreatedAt:          timestamppb.New(taskResult.CreatedAt),
			}
		}),
	}
}

func (i *Interactor) ListSubmits(ctx context.Context, groupID *string, status *backendpb.Status) (*backendpb.ListSubmitsResponse, error) {
	q := i.entClient.Submit.Query().WithGroups(func(gq *ent.GroupQuery) {
		if groupID != nil {
			gq.Where(group.NameContains(*groupID))
		}
	})
	if status != nil {
		q.Where(submit.StatusEQ(status.String()))
	}
	submits, err := q.Order(submit.BySubmitedAt(sql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	return &backendpb.ListSubmitsResponse{
		Submits: lo.Map(submits, func(submit *ent.Submit, _ int) *backendpb.Submit {
			pbSubmit := toPbSubmit(submit)
			pbSubmit.TaskResults = make([]*backendpb.TaskResult, 0)
			return pbSubmit
		}),
	}, nil
}
