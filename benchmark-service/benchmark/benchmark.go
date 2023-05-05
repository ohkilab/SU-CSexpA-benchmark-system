package benchmark

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 100,
			},
		},
	}
}

func (c *Client) Run(ctx context.Context, url string, interceptor func(req *http.Request), options ...optionFunc) <-chan *Result {
	option := &option{
		threadNum: 5,
		timeout:   5 * time.Second,
	}
	for _, f := range options {
		f(option)
	}

	resultChan := make(chan *Result)
	go func() {
		defer close(resultChan)

		errChan := make(chan error)
		defer close(errChan)

		ctx, cancel := context.WithTimeout(ctx, option.timeout)
		defer cancel()

		for range make([]struct{}, option.threadNum) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						return
					}
				}()
				for {
					resp, took, err := c.request(url, interceptor, time.Minute)
					if err != nil {
						errChan <- err
						return
					}
					httpResult, err := NewResultWithHttpResult(resp, took)
					if err != nil {
						errChan <- err
						return
					}
					resultChan <- httpResult
				}
			}()
		}

		for {
			select {
			case <-ctx.Done():
				log.Println("ctx done")
				return
			case err := <-errChan:
				log.Println("errChan:", err)
				resultChan <- NewResultWithError(err)
				return
			}
		}
	}()

	return resultChan
}

func (c *Client) request(url string, interceptor func(req *http.Request), timeout time.Duration) (*http.Response, time.Duration, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	interceptor(req)
	now := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	return resp, time.Now().Sub(now), nil
}

type Result struct {
	HttpResult *HttpResult
	Err        error
}

type HttpResult struct {
	StatusCode   int
	ContentType  string
	Body         []byte
	ResponseTime time.Duration
}

func NewResultWithHttpResult(resp *http.Response, responseTime time.Duration) (*Result, error) {
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Result{
		HttpResult: &HttpResult{
			StatusCode:   resp.StatusCode,
			ContentType:  resp.Header.Get("Content-Type"),
			Body:         b,
			ResponseTime: responseTime,
		},
	}, nil
}

func NewResultWithError(err error) *Result {
	return &Result{Err: err}
}
