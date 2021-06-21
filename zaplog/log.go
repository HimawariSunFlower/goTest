package zaplog

import "go.uber.org/zap"

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
