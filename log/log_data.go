package log

import "time"

// DataLog 日志数据
type DataLog struct {
	Name    string
	Brokers string
	Topic   string
	// 消息
	Message string

	// rpc 请求响应
	RPC RPCRequest

	// 日期记录时间
	LogTime time.Time
}

// RPCRequest rpc request
type RPCRequest struct {
	// 请求的状态码
	StatusCode int

	// 远程服务器响应内容
	ResponseContent string

	// 错误内容
	Err error
}
