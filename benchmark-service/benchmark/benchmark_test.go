package benchmark

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func Test_Run(t *testing.T) {
	port := launchTestServer(t)
	time.Sleep(time.Second)
	c := NewClient()
	results, err := c.Run(context.Background(), fmt.Sprintf("http://0.0.0.0:%v", port), func(uri *url.URL, body io.ReadCloser) error {
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range results {
		log.Println(r)
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
