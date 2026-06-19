package admin

import (
	"context"
	"strings"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/entutil"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	u_contest "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/contest"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Interactor struct {
	entClient     *ent.Client
	logger        *slog.Logger
	tagRepository tag.Repository
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger, tagRepository tag.Repository) *Interactor {
	return &Interactor{entClient, logger, tagRepository}
}

func (i *Interactor) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestResponse, error) {
	slug := strings.TrimSpace(req.Slug)
	var tagSelectionLogic contest.TagSelectionLogic
	switch req.TagSelection.(type) {
	case *pb.CreateContestRequest_Auto:
		tagSelectionLogic = contest.TagSelectionLogicAuto
		if err := i.tagRepository.CreateRandomTag(slug, req.GetAuto().Tags.Tags); err != nil {
			i.logger.Error("failed to create random tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create random tags")
		}
	case *pb.CreateContestRequest_Manual:
		tagSelectionLogic = contest.TagSelectionLogicManual
		tagsList := make([][]string, 0, len(req.GetManual().TagsList))
		for _, tags := range req.GetManual().TagsList {
			tagsList = append(tagsList, tags.Tags)
		}
		if err := i.tagRepository.CreateTags(slug, tagsList); err != nil {
			i.logger.Error("failed to create manual tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create manual tags")
		}
	}

	contest, err := i.entClient.Contest.Create().
		SetTitle(strings.TrimSpace(req.Title)).
		SetSlug(slug).
		SetStartAt(req.StartAt.AsTime()).
		SetEndAt(req.EndAt.AsTime()).
		SetSubmitLimit(int(req.SubmitLimit)).
		SetTagSelectionLogic(tagSelectionLogic).
		SetValidator(validatorName(req.Validator)).
		SetCreatedAt(timejst.Now()).
		SetTimeLimitPerTask(int64(time.Duration(req.TimeLimitPerTask) * time.Second)).
		Save(ctx)
	if err != nil {
		i.logger.Error("failed to create contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to create contest")
	}

	return &pb.CreateContestResponse{
		Contest: u_contest.ToPbContest(contest),
	}, nil
}

func (i *Interactor) UpdateContest(ctx context.Context, req *pb.UpdateContestRequest) (*pb.UpdateContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(contest.Slug(req.ContestSlug)).Only(ctx)
	if err != nil {
		i.logger.Error("failed to fetch contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch contest")
	}

	update := contest.Update()
	if req.Title != nil {
		update.SetTitle(strings.TrimSpace(*req.Title))
	}
	if req.StartAt != nil {
		update.SetStartAt(req.StartAt.AsTime())
	}
	if req.EndAt != nil {
		update.SetEndAt(req.EndAt.AsTime())
	}
	if req.SubmitLimit != nil {
		update.SetSubmitLimit(int(*req.SubmitLimit))
	}
	if req.Validator != nil {
		update.SetValidator(validatorName(*req.Validator))
	}

	contest, err = update.Save(ctx)
	if err != nil {
		i.logger.Error("failed to update contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to update contest")
	}

	return &pb.UpdateContestResponse{
		Contest: u_contest.ToPbContest(contest),
	}, nil
}

func validatorName(validator pb.Validator) string {
	if int32(validator) == 2 {
		return "V2026"
	}
	if name, ok := pb.Validator_name[int32(validator)]; ok {
		return name
	}
	return validator.String()
}

func (i *Interactor) CreateGroups(ctx context.Context, req *pb.CreateGroupsRequest) (*pb.CreateGroupsResponse, error) {
	createdGroups := make([]*pb.Group, 0, len(req.Groups))
	err := entutil.RunInTransaction(ctx, i.entClient, func(tx *ent.Tx) error {
		for _, g := range req.Groups {
			b, err := bcrypt.GenerateFromPassword([]byte(g.Password), bcrypt.DefaultCost)
			if err != nil {
				return status.Error(codes.Internal, "failed to generate password")
			}
			group, err := tx.Group.Create().
				SetName(strings.TrimSpace(g.Name)).
				SetEncryptedPassword(string(b)).
				SetRole(g.Role.String()).
				SetYear(int(g.Year)).
				SetCreatedAt(timejst.Now()).
				Save(ctx)
			if err != nil {
				return status.Error(codes.Internal, "failed to create group: "+err.Error())
			}
			createdGroups = append(createdGroups, &pb.Group{
				Name: group.Name,
				Role: pb.Role(pb.Role_value[group.Role]),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupsResponse{
		Groups: createdGroups,
	}, nil
}
