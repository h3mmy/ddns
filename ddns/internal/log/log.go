package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	devLogMode = true
)

var zloggerCommonKey = &zap.Field{Key: "group", Type: zapcore.StringType, String: "internal"}

// Setup New Common Zap Logger
func BaseZapConfig(level zapcore.Level) zap.Config {
	var cfg zap.Config
	if devLogMode {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	cfg.Level = zap.NewAtomicLevelAt(level)
	return cfg
}

// Generate new zap logger
func NewZapLogger() *zap.Logger {
	zapConfig := zap.NewDevelopmentConfig()
	zlogger, _ := zapConfig.Build()
	return zlogger.With(*zloggerCommonKey)
}
