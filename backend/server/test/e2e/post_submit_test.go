package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/tag"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/test/utils"
	mock_worker "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker/mock"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

func Test_PostSubmit(t *testing.T) {
	ctx := context.Background()
	entClient, cleanupFunc := utils.EnttestOpen(ctx, t)
	defer cleanupFunc(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	worker := mock_worker.NewMockWorker(ctrl)
	worker.EXPECT().Push(gomock.Any()).AnyTimes()
	secret := []byte("secret")
	mockRepository := tag.MockRepository(
		func(contestSlug string, num int) ([]string, error) {
			return []string{"a", "b", "c"}, nil
		}, nil)
	conn, closeFunc := utils.LaunchGrpcServer(t, grpc.WithJwtSecret("secret"), grpc.WithEntClient(entClient), grpc.WithWorker(worker), grpc.WithTagRepository(mockRepository))
	defer closeFunc()
	client := pb.NewBackendServiceClient(conn)

	// prepare
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	group, err := entClient.Group.Create().
		SetName("test").
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(group.RoleContestant).
		SetScore(12345).
		SetYear(2023).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	contest, err := entClient.Contest.Create().
		SetTitle("test contest").
		SetSlug("test-contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetYear(2023).
		SetSubmitLimit(9999).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		t.Fatal(err)
	}

	jwtToken, err := auth.GenerateJWTToken(secret, group.ID, group.Year)
	if err != nil {
		t.Fatal(err)
	}
	meta := metadata.New(map[string]string{"authorization": "Bearer " + jwtToken})
	ctx = metadata.NewOutgoingContext(ctx, meta)

	// success
	req := &pb.PostSubmitRequest{
		Url:       "http://10.255.255.255",
		ContestId: int32(contest.ID),
	}
	t.Log(contest.ID)
	resp, err := client.PostSubmit(ctx, req)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.Url, resp.Url)
	assert.Equal(t, req.ContestId, resp.ContestId)
	assert.NotEmpty(t, resp.SubmitedAt)

	// failed
	req = &pb.PostSubmitRequest{
		Url:       "http://10.255.255.255",
		ContestId: 0,
	}
	_, err = client.PostSubmit(ctx, req)
	assert.Error(t, err)
}
