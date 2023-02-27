package handlers

import (
	"github.com/h3mmy/ddns/ddns/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CloudflareHandler struct {
	logger *zap.Logger
}

func NewCloudflareUpdater() *CloudflareHandler {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Key: "handler", Type: zapcore.StringType, String: "CloudflareUpdater"},
	)
	return &CloudflareHandler{
		logger: lgr,
	}
}
