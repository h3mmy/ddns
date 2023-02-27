package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/fatih/color"
	"github.com/h3mmy/ddns/ddns"
	"golang.org/x/sync/errgroup"
)

var (
	//go:embed VERSION
	Version string
)

func main() {
	color.Magenta(fmt.Sprintf("DDNS Gopher: v%s", Version))

	ctx := context.Background()
	errGroup, ctx := errgroup.WithContext(ctx)
	ddnsWorker := ddns.NewDDNSWorker()

	errGroup.Go(func() error {
		return ddnsWorker.Start(ctx)
	})

	errGroup.Wait()
}
