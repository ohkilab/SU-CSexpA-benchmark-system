package e2e

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	pkggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func launchGrpcServer(t *testing.T, optionFuncs ...grpc.OptionFunc) (*pkggrpc.Server, *pkggrpc.ClientConn) {
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
	return server, conn
}

func enttestOpen(ctx context.Context, t *testing.T) *ent.Client {
	dsn := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("TEST_MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	// migration
	for i := time.Duration(2); ; i = i * 2 {
		if err := entClient.Schema.Create(ctx); err != nil {
			// sometimes "mysql: querying mysql version driver: bad connection" error occurs, so retry it
			time.Sleep(i * time.Second)
		} else {
			break
		}
	}
	return entClient
}
