// Package config contains environment variables.
package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config for launching an application.
type Config struct {
	ServerAddress     string `env:"KEEPER_SERVER_ADDRESS" envDefault:":3200"`
	DownloadFilesPath string `env:"KEEPER_SERVER_ADDRESS" envDefault:"/keeper/download"`
	UploadFilesPath   string `env:"KEEPER_SERVER_ADDRESS" envDefault:"/keeper/upload/"`
}

// GetConfig creates a new configuration.
func GetConfig() (*Config, error) {

	cfg := &Config{}

	// flags
	flag.StringVar(&cfg.ServerAddress, "s", cfg.ServerAddress, "KEEPER_SERVER_ADDRESS")
	flag.StringVar(&cfg.ServerAddress, "d", cfg.ServerAddress, "KEEPER_SERVER_ADDRESS")
	flag.StringVar(&cfg.ServerAddress, "u", cfg.ServerAddress, "KEEPER_SERVER_ADDRESS")
	flag.Parse()

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
