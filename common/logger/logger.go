package logger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func initLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}
	l, _ := zap.NewDevelopment()
	defer l.Sync()
	logger = l.Sugar()
	return logger
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		return initLogger()
	}
	return logger
}
