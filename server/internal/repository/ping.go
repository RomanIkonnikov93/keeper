package repository

import (
	"context"

	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/conn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Pinger for ping data base.
type Pinger interface {
	PingDB() error
}

// Ping for ping data base.
type Ping struct {
	pool *pgxpool.Pool
}

// NewPing for ping data base.
func NewPing(cfg config.Config) (*Ping, error) {

	pool := conn.NewConnection(cfg)

	return &Ping{
		pool: pool,
	}, nil
}

// PingDB for ping data base.
func (p *Ping) PingDB() error {

	pool := p.pool
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	err := pool.Ping(ctx)

	return err
}
