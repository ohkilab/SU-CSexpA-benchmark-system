package e2e

import (
	"log"
	"net"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc"
)

func init() {
	server := grpc.NewServer()
	lsnr, err := net.Listen("tcp", ":3776")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := server.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
}
