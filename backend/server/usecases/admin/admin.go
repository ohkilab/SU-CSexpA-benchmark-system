package admin

import (
	"context"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/entutil"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	entgroup "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	u_contest "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/usecases/contest"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Interactor struct {
	entClient        *ent.Client
	logger           *slog.Logger
	tagRepository    tag.Repository
	v2026ContestSlug string
}

func NewInteractor(entClient *ent.Client, logger *slog.Logger, tagRepository tag.Repository, v2026ContestSlug string) *Interactor {
	if strings.TrimSpace(v2026ContestSlug) == "" {
		v2026ContestSlug = "v2026"
	}
	return &Interactor{entClient, logger, tagRepository, strings.TrimSpace(v2026ContestSlug)}
}

func (i *Interactor) CreateContest(ctx context.Context, req *pb.CreateContestRequest) (*pb.CreateContestResponse, error) {
	slug := strings.TrimSpace(req.Slug)
	if req.UseExistingTagFiles {
		if req.Validator != pb.Validator_V2026 {
			return nil, status.Error(codes.InvalidArgument, "use_existing_tag_files is only supported for V2026")
		}
		if slug == "" {
			slug = i.v2026ContestSlug
		}
	}
	if slug == "" {
		return nil, status.Error(codes.InvalidArgument, "slug is required")
	}

	var tagSelectionLogic contest.TagSelectionLogic
	switch req.TagSelection.(type) {
	case *pb.CreateContestRequest_Auto:
		tagSelectionLogic = contest.TagSelectionLogicAuto
		if req.UseExistingTagFiles {
			if err := i.validateExistingTagFiles(slug, tagSelectionLogic, int(req.SubmitLimit)); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			break
		}
		if err := i.tagRepository.CreateRandomTag(slug, req.GetAuto().Tags.Tags); err != nil {
			i.logger.Error("failed to create random tags", "error", err)
			return nil, status.Error(codes.Internal, "failed to create random tags")
		}
	case *pb.CreateContestRequest_Manual:
		tagSelectionLogic = contest.TagSelectionLogicManual
		if req.UseExistingTagFiles {
			if err := i.validateExistingTagFiles(slug, tagSelectionLogic, int(req.SubmitLimit)); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			break
		}
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
		if ent.IsConstraintError(err) {
			return nil, status.Error(codes.AlreadyExists, "contest slug already exists")
		}
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

func (i *Interactor) DeleteContest(ctx context.Context, req *pb.DeleteContestRequest) (*pb.DeleteContestResponse, error) {
	contest, err := i.entClient.Contest.Query().Where(contest.Slug(strings.TrimSpace(req.ContestSlug))).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "contest not found")
		}
		i.logger.Error("failed to fetch contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch contest")
	}

	submitCount, err := contest.QuerySubmits().Count(ctx)
	if err != nil {
		i.logger.Error("failed to count contest submits", "error", err)
		return nil, status.Error(codes.Internal, "failed to count contest submits")
	}
	if submitCount > 0 {
		return nil, status.Error(codes.FailedPrecondition, "contest has submits")
	}

	if err := i.entClient.Contest.DeleteOne(contest).Exec(ctx); err != nil {
		i.logger.Error("failed to delete contest", "error", err)
		return nil, status.Error(codes.Internal, "failed to delete contest")
	}
	if contest.Validator == "V2026" {
		return &pb.DeleteContestResponse{}, nil
	}
	if err := i.tagRepository.DeleteContestTags(contest.Slug); err != nil {
		i.logger.Error("failed to delete contest tags", "error", err)
		return nil, status.Error(codes.Internal, "failed to delete contest tags")
	}
	return &pb.DeleteContestResponse{}, nil
}

func (i *Interactor) validateExistingTagFiles(slug string, tagSelectionLogic contest.TagSelectionLogic, submitLimit int) error {
	switch tagSelectionLogic {
	case contest.TagSelectionLogicAuto:
		tags, err := i.tagRepository.GetRandomTags(slug, 1)
		if err != nil {
			return fmt.Errorf("existing auto tag file is required: tags/%s/random.txt", slug)
		}
		if len(tags) == 0 {
			return fmt.Errorf("existing auto tag file must not be empty: tags/%s/random.txt", slug)
		}
	case contest.TagSelectionLogicManual:
		for count := 1; count <= submitLimit; count++ {
			tags, err := i.tagRepository.GetTags(slug, count)
			if err != nil {
				return fmt.Errorf("existing manual tag file is required: tags/%s/%d.txt", slug, count)
			}
			if len(tags) == 0 {
				return fmt.Errorf("existing manual tag file must not be empty: tags/%s/%d.txt", slug, count)
			}
		}
	}
	return nil
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

func (i *Interactor) ListGroups(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	groups, err := i.entClient.Group.Query().
		Order(entgroup.ByYear(sql.OrderDesc()), entgroup.ByName()).
		All(ctx)
	if err != nil {
		i.logger.Error("failed to fetch groups", "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch groups")
	}

	return &pb.ListGroupsResponse{
		Groups: toPbAdminGroups(groups),
	}, nil
}

func toPbAdminGroups(groups []*ent.Group) []*pb.AdminGroup {
	pbGroups := make([]*pb.AdminGroup, 0, len(groups))
	for _, group := range groups {
		pbGroups = append(pbGroups, &pb.AdminGroup{
			Id:   int32(group.ID),
			Name: group.Name,
			Year: int32(group.Year),
			Role: pb.Role(pb.Role_value[group.Role]),
		})
	}
	return pbGroups
}

func (i *Interactor) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	group, err := i.entClient.Group.Get(ctx, int(req.GroupId))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "group not found")
		}
		i.logger.Error("failed to fetch group", "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch group")
	}

	submitCount, err := group.QuerySubmits().Count(ctx)
	if err != nil {
		i.logger.Error("failed to count group submits", "error", err)
		return nil, status.Error(codes.Internal, "failed to count group submits")
	}
	if submitCount > 0 {
		return nil, status.Error(codes.FailedPrecondition, "group has submits")
	}

	if err := i.entClient.Group.DeleteOne(group).Exec(ctx); err != nil {
		i.logger.Error("failed to delete group", "error", err)
		return nil, status.Error(codes.Internal, "failed to delete group")
	}
	return &pb.DeleteGroupResponse{}, nil
}
