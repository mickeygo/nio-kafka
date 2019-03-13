package rpc

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// HTTPClient client
type HTTPClient struct {
	// URL 请求地址
	URL string
	// Timeout 超时时间（秒）
	Timeout int
}

const _httpCircuitBreaker string = "rpc_http_circuit"
const _defaultTimeout int = 30 * 1000 // milliseconds

func init() {
	hystrix.ConfigureCommand(_httpCircuitBreaker, hystrix.CommandConfig{
		Timeout: _defaultTimeout,
	})
}

// NewHTTPClient new a client
func NewHTTPClient() HTTPClient {
	client := HTTPClient{Timeout: _defaultTimeout}

	return client
}

// Get GET HTTPMethod
func (c HTTPClient) Get(url string, header http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}

	return c.Do(req)
}

// Post POST HTTPMethod
func (c HTTPClient) Post(url, contentType string, body io.Reader, header http.Header) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// PostAsJSON 以 json 数据格式提交 post 请求
// `data` 参数应该为字符串 json 格式
func (c HTTPClient) PostAsJSON(url, data string, header http.Header) (*http.Response, error) {
	return c.Post(url, "application/json", strings.NewReader(data), header)
}

// Do execute an http request
func (c HTTPClient) Do(req *http.Request) (*http.Response, error) {
	if req == nil {
		panic("nil http.Request")
	}

	hystrix.Go(_httpCircuitBreaker, func() error {
		return nil
	}, func(err error) error {
		return err
	})

	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(_defaultTimeout)*time.Millisecond)
	defer cancel()

	req.WithContext(ctx)

	return http.DefaultClient.Do(req)
}
