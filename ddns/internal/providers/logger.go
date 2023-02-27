package providers

import (
	"github.com/h3mmy/ddns/ddns/internal/log"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	return log.NewZapLogger()
}
