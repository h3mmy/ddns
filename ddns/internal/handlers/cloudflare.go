package handlers

import (
	"github.com/h3mmy/ddns/ddns/internal/log"
	"go.uber.org/zap"
)

type CloudflareHandler struct {
	logger *zap.Logger
}

func NewCloudflareUpdater() *CloudflareHandler {
	return &CloudflareHandler{
		logger: log.NewLoggerForHandler("CloudflareUpdater"),
	}
}
