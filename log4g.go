//go:build log4go
// +build log4go

package logg

import "github.com/alecthomas/log4go"

func Info(arg0 interface{}, args ...interface{}) {
	log := log4go.NewLogger()
	defer log.Close()
	log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())
	log4go.Info(arg0, args...)
}
func Version() string {
	return "log4go"
}
