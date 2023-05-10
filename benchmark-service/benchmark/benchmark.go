package benchmark

import (
	"context"
	"errors"
	"io"
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

func (c *Client) Run(ctx context.Context, url string, options ...optionFunc) ([]*HttpResult, error) {
	option := &option{
		threadNum:   5,
		attmptCount: 500,
	}
	for _, f := range options {
		f(option)
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	attemptCount := 0
	results := make([]*HttpResult, 0)
	mu := sync.Mutex{}
	eg, ctx := errgroup.WithContext(ctx)
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

					done := func() bool {
						mu.Lock()
						defer mu.Unlock()
						results = append(results, &HttpResult{
							StatusCode:   resp.StatusCode,
							ContentType:  resp.Header.Get("Content-Type"),
							Body:         resp.Body,
							ResponseTime: took,
						})
						attemptCount++
						if attemptCount == option.attmptCount {
							return true
						}
						return false
					}()
					if done {
						cancel()
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

		// error occured, so close and discard response body
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
}

type HttpResult struct {
	StatusCode   int
	ContentType  string
	Body         io.ReadCloser
	ResponseTime time.Duration
}
