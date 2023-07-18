package submit

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/predicate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	backendpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	benchmarkpb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	threadNum    int = 5
	attemptCount int = 100
)

type Interactor struct {
	entClient     *ent.Client
	worker        worker.Worker
	logger        *slog.Logger
	tagRepository tag.Repository
}

func NewInteractor(entClient *ent.Client, worker worker.Worker, logger *slog.Logger, tagRepository tag.Repository) *Interactor {
	return &Interactor{entClient, worker, logger, tagRepository}
}

func (i *Interactor) PostSubmit(ctx context.Context, req *backendpb.PostSubmitRequest) (*backendpb.PostSubmitResponse, error) {
	now := timejst.Now()

	claims := interceptor.GetClaimsFromContext(ctx)
	c, err := i.entClient.Contest.Query().Where(contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.InvalidArgument, "no such contest")
		}
		i.logger.Error("failed to fetch contest", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if now.Before(c.StartAt) {
		return nil, status.Error(codes.FailedPrecondition, "the contest is not started yet")
	}
	if now.After(c.EndAt) {
		return nil, status.Error(codes.FailedPrecondition, "the contest has been finished")
	}

	predicates := []predicate.Submit{
		submit.HasGroupsWith(group.ID(claims.GroupID)),
		submit.StatusEQ(backendpb.Status_SUCCESS.String()),
		submit.HasContestsWith(contest.Slug(req.ContestSlug)),
	}
	submitCount, err := i.entClient.Submit.Query().Where(predicates...).Count(ctx)
	if err != nil {
		i.logger.Error("failed to fetch submits", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if submitCount >= c.SubmitLimit {
		return nil, status.Error(codes.FailedPrecondition, "the count of submits is exceeded the limit")
	}

	var tags []string
	switch c.TagSelectionLogic {
	case contest.TagSelectionLogicAuto:
		tags, err = i.tagRepository.GetRandomTags(c.Slug, 50)
		if err != nil {
			i.logger.Error("failed to generate tags", err)
			return nil, status.Error(codes.Internal, "failed to generate tags")
		}
	case contest.TagSelectionLogicManual:
		tags, err = i.tagRepository.GetTags(c.Slug, submitCount+1)
		if err != nil {
			i.logger.Error("failed to generate tags", err)
			return nil, status.Error(codes.Internal, "failed to generate tags")
		}
	}

	submit, err := i.entClient.Submit.Create().
		SetURL(req.Url).
		SetSubmitedAt(timejst.Now()).
		SetContestsID(c.ID).
		SetGroupsID(claims.GroupID).
		SetStatus(backendpb.Status_WAITING.String()).
		SetTaskNum(len(tags)).
		Save(ctx)
	if err != nil {
		i.logger.Error("failed to create submit", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// add a task to worker
	executeRequest := buildTask(claims.GroupID, c.Slug, backendpb.Validator(backendpb.Validator_value[c.Validator]), submit, tags, time.Duration(c.TimeLimitPerTask))
	i.worker.Push(executeRequest)

	return &backendpb.PostSubmitResponse{
		Id:          int32(submit.ID),
		Url:         submit.URL,
		ContestSlug: c.Slug,
		SubmitedAt:  timestamppb.New(submit.SubmitedAt),
	}, nil
}

func buildTask(groupID int, contestSlug string, validator backendpb.Validator, submit *ent.Submit, tags []string, timeLimitPerTask time.Duration) *worker.Task {
	return &worker.Task{
		Req: &benchmarkpb.ExecuteRequest{
			GroupId: strconv.Itoa(groupID),
			Tasks: lo.Map(tags, func(tag string, _ int) *benchmarkpb.Task {
				return &benchmarkpb.Task{
					Request: &benchmarkpb.HttpRequest{
						Url:         fmt.Sprintf("%s?tag=%s", submit.URL, tag),
						Method:      benchmarkpb.HttpMethod_GET,
						ContentType: "application/x-www-form-urlencoded",
						Body:        "",
					},
					ThreadNum:    int32(threadNum),
					AttemptCount: int32(attemptCount),
				}
			}),
			ContestSlug:      contestSlug,
			Validator:        validator,
			TimeLimitPerTask: int64(timeLimitPerTask),
		},
		SubmitID: submit.ID,
		GroupID:  groupID,
	}
}

func (i *Interactor) GetSubmit(req *backendpb.GetSubmitRequest, stream backendpb.BackendService_GetSubmitServer) error {
	for {
		s, err := i.entClient.Submit.Query().Where(submit.IDEQ(int(req.SubmitId))).WithGroups().WithTaskResults().Only(stream.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				return status.Error(codes.NotFound, "no such submit")
			}
			i.logger.Error("failed to fetch submit", err)
			return status.Error(codes.Internal, err.Error())
		}
		if err := stream.Send(&backendpb.GetSubmitResponse{
			Submit: toPbSubmit(s),
		}); err != nil {
			i.logger.Error("failed to send submit", err)
			return err
		}
		// ベンチマーク処理が完了していたら結果を返す
		if !s.CompletedAt.IsZero() {
			i.logger.Info("completed submit", slog.Any("submit", s))
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func toPbSubmit(submit *ent.Submit) *backendpb.Submit {
	return &backendpb.Submit{
		Id:        int32(submit.ID),
		GroupName: submit.Edges.Groups.Name,
		Score:     int32(submit.Score),
		Status:    backendpb.Status(backendpb.Status_value[submit.Status]),
		// Language: submit.Language,
		TagCount:     int32(submit.TaskNum),
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
				ErrorMessage:       &taskResult.ErrorMessage,
				CreatedAt:          timestamppb.New(taskResult.CreatedAt),
			}
		}),
	}
}

func (i *Interactor) ListSubmits(ctx context.Context, req *backendpb.ListSubmitsRequest) (*backendpb.ListSubmitsResponse, error) {
	q := i.entClient.Submit.Query().WithGroups().Where(submit.HasContestsWith(contest.Slug(req.ContestSlug)))
	groupPredicates := []predicate.Group{}
	if req.GroupName != nil {
		groupPredicates = append(groupPredicates, group.NameContains(*req.GroupName))
	}
	groupPredicates = append(groupPredicates, group.RoleEQ(group.RoleContestant))
	if req.ContainsGuest != nil {
		if *req.ContainsGuest {
			groupPredicates = append(groupPredicates, group.RoleEQ(group.RoleGuest))
		}
	}
	q.Where(submit.HasGroupsWith(groupPredicates...))
	if req.Status != nil {
		q.Where(submit.StatusEQ(req.Status.String()))
	}
	submits, err := q.Order(submit.BySubmitedAt(sql.OrderDesc())).All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch submits", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &backendpb.ListSubmitsResponse{
		Submits: lo.Map(submits, func(submit *ent.Submit, _ int) *backendpb.Submit {
			pbSubmit := toPbSubmit(submit)
			pbSubmit.TaskResults = make([]*backendpb.TaskResult, 0)
			return pbSubmit
		}),
	}, nil
}

func (i *Interactor) GetLatestSubmit(ctx context.Context, groupID int) (*backendpb.GetLatestSubmitResponse, error) {
	submit, err := i.entClient.Submit.Query().
		Where(submit.HasGroupsWith(group.ID(groupID))).
		Order(submit.BySubmitedAt(sql.OrderDesc())).
		First(ctx)
	if err != nil {
		i.logger.Error("failed to get latest submit", "error", err)
		return nil, status.Error(codes.Internal, "failed to get latest submit")
	}
	return &backendpb.GetLatestSubmitResponse{
		Submit: toPbSubmit(submit),
	}, nil
}
