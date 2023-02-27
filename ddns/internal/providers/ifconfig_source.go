package providers

import (
	"net/netip"

	"github.com/h3mmy/ddns/ddns/internal/discovery"
	"github.com/h3mmy/ddns/ddns/internal/handlers"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/services"
)

const (
	ifconfigV4 = "https://ifconfig.me/ip"
	ifconfigV6 = "https://ifconfig.co/ip"
)

func IfConfigTask() models.DiscoveryTask {
	return handlers.NewDiscoverIPsHandler(
		services.NewIPDiscoveryService(
			NewIfConfigSource(),
		),
	)
}

type IfConfigSource struct {
	clientV4 *discovery.HTTPSource
	clientV6 *discovery.HTTPSource
	parser   models.ContentParser
}

func NewIfConfigSource() *IfConfigSource {
	return &IfConfigSource{
		clientV4: discovery.NewHTTPSource(ifconfigV4),
		clientV6: discovery.NewHTTPSource(ifconfigV6),
		parser:   discovery.IfConfigParser,
	}
}

// Gets IPv4 address
func (ifs *IfConfigSource) GetIPv4() (netip.Addr, error) {
	return ifs.parser(ifs.clientV4.GetRaw())
}

// Gets IPv6 address
func (ifs *IfConfigSource) GetIPv6() (netip.Addr, error) {
	return ifs.parser(ifs.clientV6.GetRaw())
}
