package goka

import (
	"net/http"
	"strings"
)

// CallHTTPServer Http Server
type CallHTTPServer struct {
	successFunc func(resp *http.Response)
	errorFunc   func(err error)
}

// NewCallHTTPServer new a http server
func NewCallHTTPServer(successFunc func(resp *http.Response), errFunc func(err error)) *CallHTTPServer {
	return &CallHTTPServer{
		successFunc,
		errFunc,
	}
}

// Get http get 请求
func (s *CallHTTPServer) Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		if s.errorFunc != nil {
			s.errorFunc(err)
		}
	}
	if s.successFunc != nil {
		s.successFunc(resp)
	}
}

// PostAsJSON 以 json 数据格式提交 post 请求
func (s *CallHTTPServer) PostAsJSON(url string, json string) {
	resp, err := http.Post(url, "application/json", strings.NewReader(json))
	if err != nil {
		if s.errorFunc != nil {
			s.errorFunc(err)
		}
	}
	if s.successFunc != nil {
		s.successFunc(resp)
	}
}
