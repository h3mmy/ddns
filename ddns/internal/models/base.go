package models

import (
	"context"
	"net/http"
	"net/netip"
)

type CtxKey string
type TCPVersion string

var (
	V4 TCPVersion = "tcp4"
	V6 TCPVersion = "tcp6"
	ProcessId CtxKey = "process_id"
)

type DNSUpdater interface {
}

type DiscoveryTask interface {
	GetResultSet(ctx context.Context) (*IPSet, error)
}

type DiscoveryProvider interface {
	GetIPv4(ctx context.Context) (netip.Addr, error)
	GetIPv6(ctx context.Context) (netip.Addr, error)
}

type IPSet struct {
	GlobalIPv4 netip.Addr
	GlobalIPv6 netip.Addr
}

type ContentParser func(response *http.Response) (netip.Addr, error)
