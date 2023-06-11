package handlers

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/h3mmy/ddns/ddns/internal/log"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/services"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type DiscoverIPsHandler struct {
	logger           *zap.Logger
	discoveryService *services.IPDiscoveryService
	ipv4             netip.Addr
	ipv6             netip.Addr
}

func NewDiscoverIPsHandler(discoveryService *services.IPDiscoveryService) *DiscoverIPsHandler {
	lgr := log.NewLoggerForHandler("discover_ips_handler")
	return &DiscoverIPsHandler{
		logger:           lgr,
		discoveryService: discoveryService,
	}
}

func (dish *DiscoverIPsHandler) DiscoverIPs(ctx context.Context) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return dish.DiscoverIPv4(ctx)
	})
	errGroup.Go(func() error {
		return dish.DiscoverIPv6(ctx)
	})
	return errGroup.Wait()
}

func (dish *DiscoverIPsHandler) DiscoverIPv4(ctx context.Context) error {
	defer dish.logger.Sync()
	ipv4, err := dish.discoveryService.GetSelfV4(ctx)
	if err != nil {
		dish.logger.Error("error discovering IPv4", zap.Error(err))
		return err
	}
	dish.logger.Debug(fmt.Sprintf("IPv4: %s", ipv4.String()))
	dish.ipv4 = ipv4
	return nil
}

func (dish *DiscoverIPsHandler) DiscoverIPv6(ctx context.Context) error {
	defer dish.logger.Sync()
	ipv6, err := dish.discoveryService.GetSelfV6(ctx)
	if err != nil {
		dish.logger.Error("error discovering IPv6", zap.Error(err))
		return err
	}
	dish.logger.Debug(fmt.Sprintf("IPv6: %s", ipv6.String()))
	dish.ipv6 = ipv6
	return nil
}

func (dish *DiscoverIPsHandler) GetResultSet(ctx context.Context) (*models.IPSet, error) {
	err:= dish.DiscoverIPs(ctx)
	return &models.IPSet{
		GlobalIPv4: dish.ipv4,
		GlobalIPv6: dish.ipv6,
	}, err
}
