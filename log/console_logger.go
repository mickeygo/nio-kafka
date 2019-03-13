package log

import "fmt"

// ConsoleLogger console
type ConsoleLogger struct {
}

// Log log message
func (l *ConsoleLogger) Log(data DataLog) error {
	fmt.Println(data)

	return nil
}
