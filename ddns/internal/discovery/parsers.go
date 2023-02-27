package discovery

import (
	"net/netip"
	"strings"
)

func IfConfigParser(content string) (netip.Addr, error) {
	return netip.ParseAddr(strings.TrimSpace(content))
}
