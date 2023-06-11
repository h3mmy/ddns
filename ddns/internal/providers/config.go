package providers

import "github.com/h3mmy/ddns/ddns/internal/config"

func GetConfig() *config.DDNSConfig {
	return config.LoadConfig()
}
