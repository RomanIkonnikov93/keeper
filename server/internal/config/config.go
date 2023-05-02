// Package config contains environment variables.
package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config for launching an application.
type Config struct {
	DSN          string `env:"KEEPER_DSN"`
	GRPCPort     string `env:"GRPC_PORT" envDefault:":3200"`
	JWTSecretKey string `env:"JWT_SECRET_KEY"`
	SecretKey    string `env:"KEEPER_SECRET_KEY"` // 32 byte
}

// GetConfig creates a new configuration.
func GetConfig() (*Config, error) {

	cfg := &Config{}

	// flags
	flag.StringVar(&cfg.DSN, "d", cfg.DSN, "KEEPER_DSN")
	flag.StringVar(&cfg.GRPCPort, "g", cfg.GRPCPort, "GRPC_PORT")
	flag.Parse()

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
