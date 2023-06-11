package main

import (
	_ "embed"
	"fmt"

	"github.com/fatih/color"
	"github.com/h3mmy/ddns/ddns"
)

var (
	//go:embed VERSION
	Version string
)

func main() {
	color.Magenta(fmt.Sprintf("DDNS Gopher: v%s", Version))

	fg := color.New(color.FgCyan).Add(color.Underline)
	fg.Println("Config Bits")
	cfg := ddns.GetConfig()
	fmt.Printf("logLevel: %s\n", cfg.LoggingConfig.Level)
	fmt.Printf("outfile: %s\n", cfg.LoggingConfig.OutFile)
	// ctx := context.Background()
	// errGroup, ctx := errgroup.WithContext(ctx)
	// ddnsWorker := ddns.NewDDNSWorker()
	port := 5051
	fmt.Printf("Starting Server on port %d...\n", port)
	err := ddns.StartServer(port)
	if err != nil {
		fmt.Printf("Server ended with error: %v\n", err)
	}
	// errGroup.Go(func() error {
	// 	// return ddnsWorker.Start(ctx)
	// 	s.Serve()
	// })

	// errGroup.Wait()
}
