package config

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

var (
	appName           = "ddns"
	defaultV4Strategy = IPMatchingStrategy{
		Type:         Prefix,
		PrefixLength: 32,
	}
	defaultV6Strategy = IPMatchingStrategy{
		Type:         Prefix,
		PrefixLength: 128,
	}
)

type IPMatchingStrategyKey string

const (
	Prefix IPMatchingStrategyKey = "prefix"
)

// DDNS Config
type DDNSConfig struct {
	config        *viper.Viper
	LoggingConfig LoggingConfig `mapstructure:"logging"`
	Providers     []DomainProviderConfig
	Strategy      *IPMatchingConfig `mapstructure:"strategy"`
}

type IPMatchingConfig struct {
	V4Strategy IPMatchingStrategy `mapstructure:"ipv4"`
	V6Strategy IPMatchingStrategy `mapstructure:"ipv6"`
}

type IPMatchingStrategy struct {
	Type         IPMatchingStrategyKey
	PrefixLength int
}

type LoggingConfig struct {
	OutFile string
	Level   string
}

type DomainProviderConfig struct {
	Type     string
	Enabled  bool
	ApiKey   string
	ApiToken string
	ApiUser  string
	Domains  []Domain
}

type Domain struct {
	URL string
}

// loads config
func LoadConfig() *DDNSConfig {
	// Parse Flags
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)
	// Set Defaults
	viper.SetDefault("devMode", false)
	viper.SetDefault("logLevel", "info")

	// Add config paths
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", appName))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config %w", err))
	}

	var logConfig LoggingConfig
	var domainProviders []DomainProviderConfig
	var v4Strategy IPMatchingStrategy
	var v6Strategy IPMatchingStrategy

	if viper.IsSet("logging") {
		err := viper.UnmarshalKey("logging", &logConfig)
		if err != nil {
			panic(fmt.Errorf("fatal error reading config %w", err))
		}
	}

	if viper.IsSet("providers") {
		err := viper.UnmarshalKey("providers", &domainProviders)
		if err != nil {
			panic(fmt.Errorf("fatal error reading config %w", err))
		}
	}

	if viper.IsSet("strategy.ipv4") {
		err := viper.UnmarshalKey("strategy.ipv4", &v4Strategy)
		if err != nil {
			panic(fmt.Errorf("fatal error reading config %w", err))
		}
	} else {
		v4Strategy = defaultV4Strategy
	}

	if viper.IsSet("strategy.ipv6") {
		err := viper.UnmarshalKey("strategy.ipv6", &v6Strategy)
		if err != nil {
			panic(fmt.Errorf("fatal error reading config %w", err))
		}
	} else {
		v6Strategy = defaultV6Strategy
	}

	return &DDNSConfig{
		LoggingConfig: logConfig,
		Providers:     domainProviders,
		Strategy: &IPMatchingConfig{
			V4Strategy: v4Strategy,
			V6Strategy: v6Strategy,
		},
	}
}

// GetConfig returns configuration
func (dc *DDNSConfig) GetRootConfig() *viper.Viper {
	return dc.config
}

func ParseLogLeve(logLevel string) zapcore.Level {
	l, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("Could not parse provided log level: %s. Defaulting to info\n", logLevel)
		return zapcore.InfoLevel
	}
	return l
}
