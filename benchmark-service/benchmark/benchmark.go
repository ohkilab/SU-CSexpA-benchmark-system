package benchmark

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	pkgurl "net/url"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Transport: &http.Transport{},
		},
	}
}

// check if the server is running. if the connection time is exceed 10 seconds, return error.
func (c *Client) CheckConnection(url string) error {
	uri, err := pkgurl.ParseRequestURI(url)
	if err != nil {
		log.Println(err)
		return status.Error(codes.InvalidArgument, "invalid url")
	}
	dialer := &net.Dialer{Timeout: 10 * time.Second}
	conn, err := dialer.Dial("tcp", uri.Host)
	if err != nil {
		log.Println(err)
		return status.Error(codes.FailedPrecondition, "failed to connect with the server")
	}
	defer conn.Close()
	return nil
}

func (c *Client) Run(ctx context.Context, url string, options ...optionFunc) ([]*HttpResult, error) {
	option := &option{
		threadNum:   5,
		attmptCount: 500,
	}
	for _, f := range options {
		f(option)
	}
	log.Printf("Start running benchmark(threadNum: %d, attemptCount: %d)\n", option.threadNum, option.attmptCount)

	attemptCount := 0
	results := make([]*HttpResult, 0)
	mu := sync.Mutex{}
	eg, ctx := errgroup.WithContext(ctx)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for range make([]struct{}, option.threadNum) {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					resp, took, err := c.request(url)
					if err != nil {
						return err
					}
					b, err := io.ReadAll(resp.Body)
					if err != nil {
						return err
					}
					resp.Body.Close()

					done := func() bool {
						mu.Lock()
						defer mu.Unlock()
						results = append(results, &HttpResult{
							StatusCode:   resp.StatusCode,
							ContentType:  resp.Header.Get("Content-Type"),
							Body:         b,
							ResponseTime: took,
						})
						attemptCount++
						return attemptCount == option.attmptCount
					}()
					if done {
						cancel()
						return nil
					}
				}
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) request(url string) (*http.Response, time.Duration, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	now := time.Now()
	for {
		resp, err := c.httpClient.Do(req)
		if errors.Is(err, syscall.ECONNREFUSED) {
			return nil, 0, err
		}
		if err == nil {
			return resp, time.Since(now), nil
		}
	}
}

type HttpResult struct {
	StatusCode   int
	ContentType  string
	Body         []byte
	ResponseTime time.Duration
}
