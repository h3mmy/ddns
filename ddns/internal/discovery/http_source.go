package discovery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/netip"

	"github.com/h3mmy/ddns/ddns/internal/log"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HTTPSource struct {
	url    string
	logger *zap.Logger
	parser models.ContentParser
}

func NewHTTPSource(url string, parser models.ContentParser) *HTTPSource {
	return &HTTPSource{
		url:    url,
		logger: log.NewLoggerForSource(url),
		parser: parser,
	}
}

func (hs *HTTPSource) GetIP() (netip.Addr, error) {
	logger := hs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetIP"},
	)
	res, err := http.Get(hs.url)
	if err != nil {
		logger.Error("error getting response from http source", zap.Error(err))
		return netip.Addr{}, err
	}
	if res.StatusCode > 400 {
		logger.Error(fmt.Sprintf("recv response status %s", res.Status))
		return netip.Addr{}, errors.New("response status not OK")
	}
	addr, err := hs.parser(res)
	if err != nil {
		logger.Error("error parsing response", zap.Error(err))
		return netip.Addr{}, err
	}
	return addr, err
}

func (hs *HTTPSource) GetIPWithContext(ctx context.Context, tcpVersion models.TCPVersion) (netip.Addr, error) {
	logger := hs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetIPWithContext"},
		zapcore.Field{Key: "tcpVersion", Type: zapcore.StringType, String: string(tcpVersion)},
	)
	request, err := http.NewRequestWithContext(ctx, "GET", hs.url, nil)
	if err != nil {
		logger.Error("error generating request", zap.Error(err))
		return netip.Addr{}, err
	}
	httpClient := getHttpClientWithVersion(tcpVersion)
	res, err := httpClient.Do(request)
	if err != nil {
		logger.Error("error getting response from http source", zap.Error(err))
		return netip.Addr{}, err
	}
	if res.StatusCode > 400 {
		logger.Error(fmt.Sprintf("recv response status %s", res.Status))
		return netip.Addr{}, errors.New("response status not OK")
	}
	addr, err := hs.parser(res)
	if err != nil {
		logger.Error("error parsing response", zap.Error(err))
		return netip.Addr{}, err
	}
	logger.Debug(fmt.Sprintf("No Error. Got %v", addr))
	return addr, err
}
