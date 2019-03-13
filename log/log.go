package log

// Logger log 接口
type Logger interface {
	// Log 记录日志
	Log(data DataLog) error
}
