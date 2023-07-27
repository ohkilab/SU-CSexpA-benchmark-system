package e2e

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
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

	group := utils.CreateGroup(ctx, t, entClient, "test-group", "test-group", group.RoleGuest)
	ctx = utils.WithJWT(ctx, t, group.ID, group.CreatedAt.Year())

	startAt := timestamppb.Now()
	endAt := timestamppb.New(startAt.AsTime().AddDate(1, 0, 0))
	resp, err := client.CreateContest(ctx, &pb.CreateContestRequest{
		Title:       "test contest",
		Slug:        "test-contest",
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
	})
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
		slices.SortFunc(filenames, func(a, b string) bool { return a < b })
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
