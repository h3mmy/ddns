package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLoggerForSource(sourceName string) *zap.Logger {
	return NewZapLogger().With(
		zapcore.Field{Key: "source", Type: zapcore.StringType, String: sourceName},
	)
}

func NewLoggerForService(serviceName string) *zap.Logger {
	return NewZapLogger().With(
		zapcore.Field{Key: "service", Type: zapcore.StringType, String: serviceName},
	)
}

func NewLoggerForHandler(handlerName string) *zap.Logger {
	return NewZapLogger().With(
		zapcore.Field{Key: "handler", Type: zapcore.StringType, String: handlerName},
	)
}
