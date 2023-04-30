// Package conn for database connection.
package conn

import (
	"context"

	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewConnection database connection.
func NewConnection(cfg config.Config) *pgxpool.Pool {

	logger := logging.GetLogger()

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()
	pool, err := pgxpool.Connect(ctx, cfg.DSN)
	if err != nil {
		logger.Fatalf("Unable to connect to database: %v\n", err)
	}

	return pool
}
