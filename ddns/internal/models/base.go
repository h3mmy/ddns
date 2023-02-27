package models

import (
	"context"
	"net/netip"
)

type CtxKey string

var (
	ProcessId CtxKey = "process_id"
)

type DNSUpdater interface {
}

type DiscoveryTask interface {
	DiscoverIPs(ctx context.Context) error
	GetResultSet() *IPSet
}

type DiscoveryProvider interface {
	GetIPv4() (netip.Addr, error)
	GetIPv6() (netip.Addr, error)
}

type IPSet struct {
	GlobalIPv4 netip.Addr
	GlobalIPv6 netip.Addr
}

type ContentParser func(content string) (netip.Addr, error)
