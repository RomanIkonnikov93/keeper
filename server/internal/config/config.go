package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config for launching an application.
type Config struct {
	DSN          string `env:"KEEPER_DSN"`
	GRPCAddress  string `env:"GRPC_PORT" envDefault:":3200"`
	JWTSecretKey string `env:"JWT_SECRET_KEY"`
	SecretKey    string `env:"KEEPER_SECRET_KEY"` // 32 byte
}

// GetConfig creates a new configuration.
func GetConfig() (*Config, error) {

	cfg := &Config{}

	// flags
	flag.StringVar(&cfg.DSN, "d", cfg.DSN, "KEEPER_DSN")
	flag.StringVar(&cfg.GRPCAddress, "g", cfg.GRPCAddress, "GRPC_PORT")
	flag.Parse()

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
