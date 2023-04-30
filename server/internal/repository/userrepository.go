// Package userrepository contains user ID storage options.
package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// IDRepository interface for user repository methods.
type IDRepository interface {
	AddUser(user string) error
	CheckUser(user string) (bool, error)
}

// usersIDpg struct for postgresql connection.
type usersIDpg struct {
	pool *pgxpool.Pool
}

// NewPGIDRepository create new postgresql user repository.
func NewPGIDRepository(pool *pgxpool.Pool) (*usersIDpg, error) {

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	migrationDir := filepath.Join(dir, "/internal/migrations/000001_users.up.sql")

	file, err := os.ReadFile(migrationDir)
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(ctx, string(file)); err != nil {
		return nil, err
	}

	return &usersIDpg{
		pool: pool,
	}, nil
}

// AddUser
func (p *usersIDpg) AddUser(user string) error {

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()

	if _, err := p.pool.Exec(ctx, `insert into users (user_id) values ($1)`, user); err != nil {
		return err
	}

	return nil
}

// CheckUser
func (p *usersIDpg) CheckUser(user string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()

	if _, err := p.pool.Exec(ctx, `insert into users (user_id) values ($1)`, user); err != nil {
		pgerr, ok := err.(*pgconn.PgError)
		if ok {
			if pgerr.Code == "23505" {
				return true, nil
			}
		}
	}
	return false, nil
}
