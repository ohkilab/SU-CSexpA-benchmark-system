package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/migrate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
	pkggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func launchGrpcServer(t *testing.T, optionFuncs ...grpc.OptionFunc) (*pkggrpc.ClientConn, func()) {
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

type cleanupFunc func(t *testing.T)

func enttestOpen(ctx context.Context, t *testing.T) (*ent.Client, cleanupFunc) {
	dsn := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("TEST_MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
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
		defer entClient.Close()

		rows, err := db.Query("SHOW TABLES;")
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				t.Fatal(err)
			}
			// for truncating forcely
			if _, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0;"); err != nil {
				t.Fatal(err)
			}
			_, _ = db.Exec("TRUNCATE TABLE " + tableName + ";")
			if _, err := db.Exec("SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
				t.Fatal(err)
			}
		}
	}
	return entClient, cleanup
}
