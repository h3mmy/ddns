package log

import (
	"os"

	"github.com/h3mmy/ddns/ddns/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogRetentionDays    = 30
	LogRetentionMaxSize = 10 // megabytes
	loggingConfigKey    = "logging"
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
func NewCommonLogger() *zap.Logger {
	// Get logfile name
	logfileName := config.LoadConfig().GetLogConfig().OutFile

	// Build Config
	var cfg zap.Config
	var core zapcore.Core
	w := getWriterSyncer(logfileName)
	if devLogMode {
		cfg = BaseZapConfig(zapcore.DebugLevel)
	} else {
		cfg = BaseZapConfig(config.LoadConfig().LoggingConfig.GetLogLevel())
	}
	core = getCore(cfg, w)
	return zap.New(core).With(*zloggerCommonKey)
}

// Setup Custom Log Core with custom filename
func BaseLogCoreWithLogFile(level zapcore.Level, filename string) zapcore.Core {
	return getCore(BaseZapConfig(level), getWriterSyncer(filename))
}

// Gets zapcore.Core with proper encoder
func getCore(cfg zap.Config, w zapcore.WriteSyncer) zapcore.Core {
	if devLogMode {
		return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), w, cfg.Level)
	} else {
		return zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), w, cfg.Level)
	}
}

func getWriterSyncer(filename string) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    LogRetentionMaxSize,
			MaxBackups: 3,
			MaxAge:     LogRetentionDays,
		}),
		zapcore.AddSync(os.Stdout),
	)
}
