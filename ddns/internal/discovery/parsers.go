package discovery

import (
	"fmt"
	"net/http"
	"net/netip"
	"strings"

	"github.com/h3mmy/ddns/ddns/internal/log"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"go.uber.org/zap"
)

// For simple string response parsin
func IfConfigStringParser(res *http.Response) (netip.Addr, error) {
	content, err := parseStringResponse(res)
	if err != nil {
		return netip.Addr{}, err
	}
	return netip.ParseAddr(strings.TrimSpace(*content))
}

// For json response parsing (ifconfig.co/json version)
func IfConfigJsonParser(response *http.Response) (netip.Addr, error) {
	logger := log.NewCommonLogger().Named("IfConfigJsonParser")
	defer logger.Sync()

	var parsedRes models.IfconfigJSONResponseCo
	err := parseJsonResponse(response, &parsedRes)
	if err != nil {
		logger.Error("parsing response failed", zap.Error(err))
		return netip.Addr{}, err
	}
	logger.Debug(fmt.Sprintf("parsing response success: %v", parsedRes))
	return netip.ParseAddr(parsedRes.IP)
}

// For json response parsing (ifconfig.me/all.json version)
func IfConfigMeJsonParser(response *http.Response) (netip.Addr, error) {
	var parsedRes models.IfconfigJSONResponseMe
	err := parseJsonResponse(response, &parsedRes)
	if err != nil {
		return netip.Addr{}, err
	}
	return netip.ParseAddr(parsedRes.IP)
}
