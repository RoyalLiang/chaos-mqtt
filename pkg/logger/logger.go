package logger

import "go.uber.org/zap"

func Info(msg string) {
	zap.S().Info(msg)
}

func Warn(msg string) {
	zap.S().Warn(msg)
}

func Error(msg string) {
	zap.S().Error(msg)
}

func Fatal(msg string) {
	zap.S().Fatal(msg)
}

func Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}
