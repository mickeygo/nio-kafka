package rpc

import (
	"github.com/afex/hystrix-go/hystrix"
)

// HTTPClient client
type HTTPClient struct {
	// URL 请求地址
	URL string
	// Timeout 超时时间（秒）
	Timeout int
}

const _httpCircuitBreaker string = "rpc_http"
const _defaultTimeout int = 30

// NewHTTPClient new a client
func NewHTTPClient() HTTPClient {
	client := HTTPClient{Timeout: _defaultTimeout}

	return client
}

// Get GET HTTPMethod
func (client HTTPClient) Get(uri string) {
	hystrix.Go(_httpCircuitBreaker, func() error {
		return nil
	}, nil)
}

// Post POST HTTPMethod
func (client HTTPClient) Post() {

}

// Put PUT HTTPMethod
func (client HTTPClient) Put() {

}

// Delete DELETE HTTPMethod
func (client HTTPClient) Delete() {

}
