package benchmark

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

func Test_Run(t *testing.T) {
	port := launchTestServer(t)
	time.Sleep(time.Second)
	c := NewClient()
	ch := c.Run(context.Background(), fmt.Sprintf("0.0.0.0:%v", port), func(req *http.Request) {}, OptTimeout(5*time.Second))
	for res := range ch {
		if res.HttpResult != nil {
			log.Println(res.HttpResult.ResponseTime)
		}
	}
}

func launchTestServer(t *testing.T) int {
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
