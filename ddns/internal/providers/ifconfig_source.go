package providers

import (
	"context"
	"net/netip"

	"github.com/h3mmy/ddns/ddns/internal/discovery"
	"github.com/h3mmy/ddns/ddns/internal/handlers"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/services"
)

const (
	// string versions replace suffix /json with /ip
	ifconfigCo = "http://ifconfig.co/json"
)

func IfConfigTask() models.DiscoveryTask {
	return handlers.NewDiscoverIPsHandler(
		services.NewIPDiscoveryService(
			NewIfConfigSource(),
		),
	)
}

type IfConfigSource struct {
	client *discovery.HTTPSource
}

func NewIfConfigSource() *IfConfigSource {
	return &IfConfigSource{
		client: discovery.NewHTTPSource(ifconfigCo, discovery.IfConfigJsonParser),
	}
}

// Gets IPv4 address
func (ifs *IfConfigSource) GetIPv4(ctx context.Context) (netip.Addr, error) {
	return ifs.client.GetIPWithContext(ctx, models.V4)
}

// Gets IPv6 address
func (ifs *IfConfigSource) GetIPv6(ctx context.Context) (netip.Addr, error) {
	return ifs.client.GetIPWithContext(ctx, models.V6)
}
