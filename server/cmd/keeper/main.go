package main

import (
	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/repository"
	"github.com/RomanIkonnikov93/keeper/server/internal/server"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"
)

func main() {

	logger := logging.GetLogger()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("GetConfig: %s", err)
	}

	rep, err := repository.NewReps(*cfg)
	if err != nil {
		logger.Fatalf("NewReps: %s", err)
	}

	err = server.StartServer(*rep, *cfg, logger)
	if err != nil {
		logger.Fatalf("StartServer: %s", err)
	}
}
