package services

import (
	"context"
	"net/netip"

	"github.com/h3mmy/ddns/ddns/internal/log"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"go.uber.org/zap"
)

type IPDiscoveryService struct {
	discoveryProvider models.DiscoveryProvider
	logger            *zap.Logger
}

func NewIPDiscoveryService(dProvider models.DiscoveryProvider) *IPDiscoveryService {
	lgr := log.NewLoggerForService("ip_discovery_service")
	return &IPDiscoveryService{
		discoveryProvider: dProvider,
		logger:            lgr,
	}
}

// Gets external IPv4 Address
func (ipds *IPDiscoveryService) GetSelfV4(ctx context.Context) (netip.Addr, error) {
	ipv4, err := ipds.discoveryProvider.GetIPv4(ctx)
	if err != nil {
		ipds.logger.Debug("error getting external ipv4 address", zap.Error(err))
		return netip.Addr{}, err
	}
	return ipv4, err
}

// Gets external IPv6 Address
func (ipds *IPDiscoveryService) GetSelfV6(ctx context.Context) (netip.Addr, error) {
	ipv6, err := ipds.discoveryProvider.GetIPv6(ctx)
	if err != nil {
		ipds.logger.Debug("error getting external ipv6 address", zap.Error(err))
		return netip.Addr{}, err
	}
	return ipv6, err
}
