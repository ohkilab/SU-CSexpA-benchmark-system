package e2e

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_CreateContest(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tmpPath, err := os.MkdirTemp("/tmp", "benchmark-system-e2e-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpPath)
	tagRepository := tag.NewRespository(tmpPath)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	group := utils.CreateGroup(ctx, t, entClient, "test-group", "test-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, group.ID, group.CreatedAt.Year(), group.Role)

	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(1, 0, 0))
	resp, err := client.CreateContest(ctx, validCreateContestRequest("test-contest", startAt, endAt))
	require.NoError(t, err)
	assert.Equal(t, "test contest", resp.Contest.Title)
	assert.Equal(t, "test-contest", resp.Contest.Slug)
	assert.Equal(t, 329, int(resp.Contest.SubmitLimit))
	assert.Equal(t, pb.TagSelectionLogicType_AUTO, resp.Contest.TagSelectionLogic)
	existTags(t, "test-contest", tmpPath, pb.TagSelectionLogicType_AUTO, [][]string{{"tag1", "tag2"}})

	_, err = client.CreateContest(ctx, &pb.CreateContestRequest{
		Title:       "test contest",
		Slug:        "test-contest-manual",
		StartAt:     startAt,
		EndAt:       endAt,
		SubmitLimit: 329,
		TagSelection: &pb.CreateContestRequest_Manual{
			Manual: &pb.TagSelectionLogicManual{
				Type: pb.TagSelectionLogicType_MANUAL,
				TagsList: []*pb.Tags{
					{
						Tags: []string{"tag1", "tag2"},
					},
					{
						Tags: []string{"tag3", "tag4"},
					},
					{
						Tags: []string{"tag5", "tag6"},
					},
				},
			},
		},
		TimeLimitPerTask: 30,
	})
	require.NoError(t, err)
	existTags(t, "test-contest-manual", tmpPath, pb.TagSelectionLogicType_MANUAL, [][]string{{"tag1", "tag2"}, {"tag3", "tag4"}, {"tag5", "tag6"}})
}

func Test_CreateV2026ContestWithExistingManualTags(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tmpPath, err := os.MkdirTemp("/tmp", "benchmark-system-e2e-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpPath)
	require.NoError(t, writeTagFile(tmpPath, "v2026", "1.txt", "tag=cat\n"))
	require.NoError(t, writeTagFile(tmpPath, "v2026", "2.txt", "tag=dog\n"))

	tagRepository := tag.NewRespository(tmpPath)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository), grpc.WithV2026ContestSlug("v2026"))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))

	resp, err := client.CreateContest(ctx, &pb.CreateContestRequest{
		Title:               "V2026 contest",
		StartAt:             startAt,
		EndAt:               endAt,
		SubmitLimit:         2,
		Validator:           pb.Validator_V2026,
		TimeLimitPerTask:    30,
		UseExistingTagFiles: true,
		TagSelection: &pb.CreateContestRequest_Manual{
			Manual: &pb.TagSelectionLogicManual{Type: pb.TagSelectionLogicType_MANUAL},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "v2026", resp.Contest.Slug)
	assert.Equal(t, pb.Validator_V2026, resp.Contest.Validator)
	assert.Equal(t, pb.TagSelectionLogicType_MANUAL, resp.Contest.TagSelectionLogic)
	existTags(t, "v2026", tmpPath, pb.TagSelectionLogicType_MANUAL, [][]string{{"tag=cat"}, {"tag=dog"}})
}

func Test_CreateV2026ContestRejectsMissingExistingTags(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tmpPath, err := os.MkdirTemp("/tmp", "benchmark-system-e2e-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpPath)
	require.NoError(t, writeTagFile(tmpPath, "v2026", "1.txt", "tag=cat\n"))

	tagRepository := tag.NewRespository(tmpPath)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository), grpc.WithV2026ContestSlug("v2026"))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))

	_, err = client.CreateContest(ctx, &pb.CreateContestRequest{
		Title:               "V2026 contest",
		StartAt:             startAt,
		EndAt:               endAt,
		SubmitLimit:         2,
		Validator:           pb.Validator_V2026,
		TimeLimitPerTask:    30,
		UseExistingTagFiles: true,
		TagSelection: &pb.CreateContestRequest_Manual{
			Manual: &pb.TagSelectionLogicManual{Type: pb.TagSelectionLogicType_MANUAL},
		},
	})
	require.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.InvalidArgument, s.Code())
	assert.Contains(t, s.Message(), "tags/v2026/2.txt")
}

func Test_AdminServiceRequiresAdminRole(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tmpPath, err := os.MkdirTemp("/tmp", "benchmark-system-e2e-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpPath)
	tagRepository := tag.NewRespository(tmpPath)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	guest := utils.CreateGroup(ctx, t, entClient, "guest-group", "guest-group", 2023, pb.Role_GUEST)
	guestCtx := utils.WithJWT(ctx, t, guest.ID, guest.CreatedAt.Year(), guest.Role)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))

	_, err = client.CreateContest(guestCtx, validCreateContestRequest("guest-contest", startAt, endAt))
	require.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.UpdateContest(guestCtx, &pb.UpdateContestRequest{ContestSlug: "guest-contest"})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.CreateGroups(guestCtx, &pb.CreateGroupsRequest{
		Groups: []*pb.CreateGroupsRequest_CreateGroupsGroup{
			{Name: "guest-created-group", Password: "password", Year: 2024, Role: pb.Role_CONTESTANT},
		},
	})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.DeleteContest(guestCtx, &pb.DeleteContestRequest{ContestSlug: "guest-contest"})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.ListGroups(guestCtx, &pb.ListGroupsRequest{})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.DeleteGroup(guestCtx, &pb.DeleteGroupRequest{GroupId: int32(guest.ID)})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	_, err = client.CreateContest(ctx, validCreateContestRequest("missing-token-contest", startAt, endAt))
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code())
}

func Test_UpdateContest(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))
	utils.CreateContest(ctx, t, entClient, "old title", "update-contest", pb.Validator_V2022.String(), startAt.AsTime(), endAt.AsTime(), 10, contest.TagSelectionLogicAuto)

	newTitle := "new title"
	newSubmitLimit := int32(20)
	validator := pb.Validator_V2023
	resp, err := client.UpdateContest(ctx, &pb.UpdateContestRequest{
		ContestSlug: "update-contest",
		Title:       &newTitle,
		SubmitLimit: &newSubmitLimit,
		Validator:   &validator,
	})
	require.NoError(t, err)
	assert.Equal(t, newTitle, resp.Contest.Title)
	assert.Equal(t, int32(20), resp.Contest.SubmitLimit)
	assert.Equal(t, pb.Validator_V2023, resp.Contest.Validator)
}

func Test_CreateGroups(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)

	resp, err := client.CreateGroups(ctx, &pb.CreateGroupsRequest{
		Groups: []*pb.CreateGroupsRequest_CreateGroupsGroup{
			{Name: "contestant-1", Password: "password", Year: 2024, Role: pb.Role_CONTESTANT},
			{Name: "guest-1", Password: "password", Year: 2024, Role: pb.Role_GUEST},
		},
	})
	require.NoError(t, err)
	require.Len(t, resp.Groups, 2)
	assert.Equal(t, pb.Role_CONTESTANT, resp.Groups[0].Role)
	assert.Equal(t, pb.Role_GUEST, resp.Groups[1].Role)
}

func Test_CreateGroupsRollsBackOnFailure(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	utils.CreateGroup(ctx, t, entClient, "duplicate", "password", 2024, pb.Role_CONTESTANT)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)

	_, err := client.CreateGroups(ctx, &pb.CreateGroupsRequest{
		Groups: []*pb.CreateGroupsRequest_CreateGroupsGroup{
			{Name: "created-before-failure", Password: "password", Year: 2024, Role: pb.Role_CONTESTANT},
			{Name: "duplicate", Password: "password", Year: 2024, Role: pb.Role_CONTESTANT},
		},
	})
	require.Error(t, err)

	count, err := entClient.Group.Query().Where(group.NameEQ("created-before-failure")).Count(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func Test_ListGroups(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	utils.CreateGroup(ctx, t, entClient, "contestant-1", "password", 2024, pb.Role_CONTESTANT)
	utils.CreateGroup(ctx, t, entClient, "guest-1", "password", 2024, pb.Role_GUEST)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)

	resp, err := client.ListGroups(ctx, &pb.ListGroupsRequest{})
	require.NoError(t, err)
	require.Len(t, resp.Groups, 3)
	assert.Equal(t, "contestant-1", resp.Groups[0].Name)
	assert.Equal(t, int32(2024), resp.Groups[0].Year)
	assert.Equal(t, pb.Role_CONTESTANT, resp.Groups[0].Role)
	assert.NotZero(t, resp.Groups[0].Id)
}

func Test_DeleteContest(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)

	tmpPath, err := os.MkdirTemp("/tmp", "benchmark-system-e2e-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpPath)
	tagRepository := tag.NewRespository(tmpPath)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithTagRepository(tagRepository))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))
	utils.CreateContest(ctx, t, entClient, "delete contest", "delete-contest", pb.Validator_V2023.String(), startAt.AsTime(), endAt.AsTime(), 10, contest.TagSelectionLogicAuto)
	require.NoError(t, tagRepository.CreateRandomTag("delete-contest", []string{"tag1"}))

	_, err = client.DeleteContest(ctx, &pb.DeleteContestRequest{ContestSlug: "delete-contest"})
	require.NoError(t, err)
	count, err := entClient.Contest.Query().Where(contest.SlugEQ("delete-contest")).Count(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)
	_, err = os.Stat(filepath.Join(tmpPath, "tags", "delete-contest"))
	assert.True(t, os.IsNotExist(err))

	utils.CreateContest(ctx, t, entClient, "delete V2026 contest", "delete-v2026-contest", pb.Validator_V2026.String(), startAt.AsTime(), endAt.AsTime(), 10, contest.TagSelectionLogicAuto)
	require.NoError(t, tagRepository.CreateRandomTag("delete-v2026-contest", []string{"tag1"}))

	_, err = client.DeleteContest(ctx, &pb.DeleteContestRequest{ContestSlug: "delete-v2026-contest"})
	require.NoError(t, err)
	_, err = os.Stat(filepath.Join(tmpPath, "tags", "delete-v2026-contest"))
	assert.NoError(t, err)

	contestWithSubmit := utils.CreateContest(ctx, t, entClient, "used contest", "used-contest", pb.Validator_V2026.String(), startAt.AsTime(), endAt.AsTime(), 10, contest.TagSelectionLogicAuto)
	utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contestWithSubmit, admin)
	_, err = client.DeleteContest(ctx, &pb.DeleteContestRequest{ContestSlug: "used-contest"})
	require.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.FailedPrecondition, s.Code())

	_, err = client.DeleteContest(ctx, &pb.DeleteContestRequest{ContestSlug: "missing-contest"})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.NotFound, s.Code())
}

func Test_DeleteGroup(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient))
	defer closeFunc()
	client := pb.NewAdminServiceClient(conn)

	admin := utils.CreateGroup(ctx, t, entClient, "admin-group", "admin-group", 2023, pb.Role_ADMIN)
	deletable := utils.CreateGroup(ctx, t, entClient, "deletable", "password", 2024, pb.Role_CONTESTANT)
	used := utils.CreateGroup(ctx, t, entClient, "used", "password", 2024, pb.Role_CONTESTANT)
	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(0, 1, 0))
	contest := utils.CreateContest(ctx, t, entClient, "contest", "delete-group-contest", pb.Validator_V2026.String(), startAt.AsTime(), endAt.AsTime(), 10, contest.TagSelectionLogicAuto)
	utils.CreateSubmit(ctx, t, entClient, 100, pb.Status_SUCCESS.String(), contest, used)
	ctx = utils.WithJWT(ctx, t, admin.ID, admin.CreatedAt.Year(), admin.Role)

	_, err := client.DeleteGroup(ctx, &pb.DeleteGroupRequest{GroupId: int32(deletable.ID)})
	require.NoError(t, err)
	count, err := entClient.Group.Query().Where(group.IDEQ(deletable.ID)).Count(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)

	_, err = client.DeleteGroup(ctx, &pb.DeleteGroupRequest{GroupId: int32(used.ID)})
	require.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, codes.FailedPrecondition, s.Code())

	_, err = client.DeleteGroup(ctx, &pb.DeleteGroupRequest{GroupId: int32(admin.ID)})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.FailedPrecondition, s.Code())

	_, err = client.DeleteGroup(ctx, &pb.DeleteGroupRequest{GroupId: 999999})
	require.Error(t, err)
	s, _ = status.FromError(err)
	assert.Equal(t, codes.NotFound, s.Code())
}

func writeTagFile(rootPath, slug, fileName, content string) error {
	dir := filepath.Join(rootPath, "tags", slug)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, fileName), []byte(content), 0o644)
}

func validCreateContestRequest(slug string, startAt, endAt *timestamppb.Timestamp) *pb.CreateContestRequest {
	return &pb.CreateContestRequest{
		Title:       "test contest",
		Slug:        slug,
		StartAt:     startAt,
		EndAt:       endAt,
		SubmitLimit: 329,
		TagSelection: &pb.CreateContestRequest_Auto{
			Auto: &pb.TagSelectionLogicAuto{
				Type: pb.TagSelectionLogicType_AUTO,
				Tags: &pb.Tags{
					Tags: []string{"tag1", "tag2"},
				},
			},
		},
		TimeLimitPerTask: 30,
	}
}

func existTags(t *testing.T, slug, tmpPath string, logicType pb.TagSelectionLogicType, tags [][]string) {
	switch logicType {
	case pb.TagSelectionLogicType_AUTO:
		f, err := os.Open(filepath.Join(tmpPath, "tags", slug, "random.txt"))
		require.NoError(t, err)
		defer f.Close()
		sc := bufio.NewScanner(f)
		i := 0
		for {
			require.True(t, sc.Scan())
			assert.Equal(t, tags[0][i], sc.Text())
			i++
			if i >= len(tags[0]) {
				break
			}
		}
	case pb.TagSelectionLogicType_MANUAL:
		rootPath := filepath.Join(tmpPath, "tags", slug)
		entries, err := os.ReadDir(rootPath)
		require.NoError(t, err)
		require.Equal(t, len(tags), len(entries))
		filenames := make([]string, 0, len(entries))
		for _, entry := range entries {
			filenames = append(filenames, entry.Name())
		}
		slices.SortFunc(filenames, func(a, b string) int {
			if a < b {
				return -1
			} else if a == b {
				return 0
			}
			return 1
		})
		for i, filename := range filenames {
			func() {
				f, err := os.Open(filepath.Join(rootPath, filename))
				require.NoError(t, err)
				defer f.Close()
				sc := bufio.NewScanner(f)
				j := 0
				for {
					require.True(t, sc.Scan())
					assert.Equal(t, tags[i][j], sc.Text())
					j++
					if j >= len(tags[i]) {
						break
					}
				}
			}()
		}
	}
}
