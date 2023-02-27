package discovery

import (
	"context"
	"net/http"

	"github.com/h3mmy/ddns/ddns/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HTTPSource struct {
	url    string
	logger *zap.Logger
}

func NewHTTPSource(url string) *HTTPSource {
	return &HTTPSource{
		url:    url,
		logger: log.NewLoggerForSource(url),
	}
}

func (hs *HTTPSource) GetRaw() string {
	logger := hs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetRaw"},
	)
	res, err := http.Get(hs.url)
	if err != nil {
		logger.Error("error getting response from http source", zap.Error(err))
		return ""
	}
	var response string
	err = parseResponse(res, &response)
	if err != nil {
		logger.Error("error parsing response", zap.Error(err))
		return ""
	}
	return response
}

func (hs *HTTPSource) GetRawWithContext(ctx context.Context) string {
	logger := hs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetRawWithContext"},
	)
	request, err := http.NewRequestWithContext(ctx, "GET", hs.url, nil)
	if err != nil {
		logger.Error("error generating request", zap.Error(err))
		return ""
	}
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.Error("error getting response from http source", zap.Error(err))
		return ""
	}
	var response string
	err = parseResponse(res, &response)
	if err != nil {
		logger.Error("error parsing response", zap.Error(err))
		return ""
	}
	return response
}
