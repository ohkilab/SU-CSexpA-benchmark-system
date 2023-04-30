package e2e

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/enttest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	pkggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func launchGrpcServer(optionFuncs ...grpc.OptionFunc) (*pkggrpc.ClientConn, error) {
	server := grpc.NewServer(optionFuncs...)
	lsnr, err := net.Listen("tcp", ":3776")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := server.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	return pkggrpc.Dial("localhost:3776", pkggrpc.WithTransportCredentials(insecure.NewCredentials()))
}

func enttestOpen(t *testing.T) *ent.Client {
	t.Helper()
	dsn := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("TEST_MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	driver, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	db := driver.DB()
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Minute)
	return enttest.NewClient(t, enttest.WithOptions(ent.Driver(driver)))
}
