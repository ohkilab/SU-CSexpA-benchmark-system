package benchmark

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        500,
				MaxIdleConnsPerHost: 100,
			},
		},
	}
}

func (c *Client) CheckConnection(ctx context.Context, url string) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()
	return err
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
				log.Println("check", attemptCount)
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
