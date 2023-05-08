package utils

import (
	"net"
	"net/http"
	"testing"
	"time"
)

func LaunchTestServer(t *testing.T) int {
	t.Helper()

	lsnr, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})
	go func() {
		if err := http.Serve(lsnr, nil); err != nil {
			t.Log(err)
		}
	}()
	return lsnr.Addr().(*net.TCPAddr).Port
}
