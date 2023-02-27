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
// TODO: Use context for discovery requests
func (ipds *IPDiscoveryService) GetSelfV4(ctx context.Context) (*netip.Addr, error) {
	ipv4, err := ipds.discoveryProvider.GetIPv4()
	if err != nil {
		ipds.logger.Debug("error getting external ipv4 address", zap.Error(err))
		return nil, err
	}
	return &ipv4, nil
}

// Gets external IPv6 Address
// TODO: Use context for discovery requests
func (ipds *IPDiscoveryService) GetSelfV6(ctx context.Context) (*netip.Addr, error) {
	ipv6, err := ipds.discoveryProvider.GetIPv6()
	if err != nil {
		ipds.logger.Debug("error getting external ipv4 address", zap.Error(err))
		return nil, err
	}
	return &ipv6, nil
}
