package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var pkgKey = "package"

func NewLoggerForSource(sourceName string) *zap.Logger {
	return NewCommonLogger().Named(sourceName).With(
		zapcore.Field{Key: pkgKey, Type: zapcore.StringType, String: "source"},
	)
}

func NewLoggerForService(serviceName string) *zap.Logger {
	return NewCommonLogger().Named(serviceName).With(
		zapcore.Field{Key: pkgKey, Type: zapcore.StringType, String: "service"},
	)
}

func NewLoggerForHandler(handlerName string) *zap.Logger {
	return NewCommonLogger().Named(handlerName).With(
		zapcore.Field{Key: pkgKey, Type: zapcore.StringType, String: "handler"},
	)
}
