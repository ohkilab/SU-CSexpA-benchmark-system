package utils

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/auth"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/migrate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/benchmark"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/service"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/validation"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/benchmark-service"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	pkggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func LaunchBenchmarkGrpcServer(t *testing.T) (*pkggrpc.ClientConn, func()) {
	t.Helper()
	server := pkggrpc.NewServer()
	client := benchmark.NewClient()
	benchmarkService := service.New(client, make(map[string]validation.Validator))
	server.RegisterService(&pb.BenchmarkService_ServiceDesc, benchmarkService)
	lsnr, err := net.Listen("tcp", ":3777")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		if err := server.Serve(lsnr); err != nil {
			t.Log(err)
		}
	}()
	conn, err := pkggrpc.Dial("localhost:3777", pkggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	return conn, func() {
		server.GracefulStop()
		conn.Close()
	}
}

func LaunchGrpcServer(t *testing.T, optionFuncs ...grpc.OptionFunc) (*pkggrpc.ClientConn, func()) {
	t.Helper()
	server := grpc.NewServer(optionFuncs...)
	lsnr, err := net.Listen("tcp", ":3776")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		if err := server.Serve(lsnr); err != nil {
			t.Log(err)
		}
	}()
	conn, err := pkggrpc.Dial("localhost:3776", pkggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	return conn, func() {
		server.GracefulStop()
		conn.Close()
	}
}

func LaunchTestServer(t *testing.T) int {
	t.Helper()

	lsnr, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go func() {
		if err := http.Serve(lsnr, nil); err != nil {
			t.Log(err)
		}
	}()
	return lsnr.Addr().(*net.TCPAddr).Port
}

type cleanupFunc func(t *testing.T)

func EnttestOpen(ctx context.Context, t *testing.T) (*ent.Client, cleanupFunc) {
	dsn := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s?parseTime=true", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("TEST_MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	entClient := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
	// migration
	for i := time.Duration(2); ; i = i * 2 {
		if err := entClient.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			// sometimes "mysql: querying mysql version driver: bad connection" error occurs, so retry it
			time.Sleep(i * time.Second)
		} else {
			break
		}
	}
	cleanup := func(t *testing.T) {
		_, err = entClient.Contest.Delete().Exec(ctx)
		require.NoError(t, err)
		_, err = entClient.Group.Delete().Exec(ctx)
		require.NoError(t, err)
		_, err = entClient.Submit.Delete().Exec(ctx)
		require.NoError(t, err)
		_, err = entClient.TaskResult.Delete().Exec(ctx)
		require.NoError(t, err)
		entClient.Close()
	}
	return entClient, cleanup
}

func CreateGroup(ctx context.Context, t *testing.T, entClient *ent.Client, name, password string, role group.Role) *ent.Group {
	t.Helper()
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	group, err := entClient.Group.Create().
		SetName(name).
		SetEncryptedPassword(string(encryptedPassword)).
		SetRole(role).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	require.NoError(t, err)
	return group
}

func CreateSubmit(ctx context.Context, t *testing.T, entClient *ent.Client, score int, status string, contest *ent.Contest, group *ent.Group) *ent.Submit {
	t.Helper()
	submit, err := entClient.Submit.Create().
		SetURL("http://localhost:8080/program").
		SetContests(contest).
		SetGroups(group).
		SetScore(score).
		SetStatus(status).
		SetSubmitedAt(timejst.Now()).
		SetTaskNum(50).
		Save(ctx)
	require.NoError(t, err)
	return submit
}

func WithJWT(ctx context.Context, t *testing.T, id, year int) context.Context {
	token, err := auth.GenerateJWTToken([]byte("secret"), id, year)
	require.NoError(t, err)
	meta := metadata.New(map[string]string{"authorization": "Bearer " + token})
	return metadata.NewOutgoingContext(ctx, meta)
}
