package main

import (
	"github.com/RomanIkonnikov93/keeper/client/internal/config"
	"github.com/RomanIkonnikov93/keeper/client/internal/gapi"
	"github.com/RomanIkonnikov93/keeper/client/internal/tui"
	"github.com/RomanIkonnikov93/keeper/client/pkg/logging"
)

func main() {

	logger := logging.GetLogger()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("GetConfig: %s", err)
	}

	client, err := gapi.InitServices(*cfg, logger)
	if err != nil {
		logger.Fatalf("GetConfig: %s", err)
	}

	t := tui.NewTUI(client)

	logger.Fatal(t.Run())
}
