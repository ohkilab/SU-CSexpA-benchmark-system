package benchmark

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/benchmark-service/test/utils"
)

func Test_Run(t *testing.T) {
	port := utils.LaunchTestServer(t)
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
