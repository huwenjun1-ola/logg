//go:build zap
// +build zap

package logg

import "go.uber.org/zap"

func Info(arg0 interface{}, args ...interface{}) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info(arg0, args)
}
func Version() string {
	return "zap"
}
