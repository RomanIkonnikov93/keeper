// Package repository contains storage options for links and short links.
package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository interface for repository methods.
type Repository interface {
	Add(ctx context.Context, short, long, id string) error
	Get(ctx context.Context, short string) (string, error)
	GetAllByType(ctx context.Context) (int, int, error)
	UpdateByID(ctx context.Context, short string) (string, error)
	DeleteByID(ctx context.Context) (int, int, error)
}

// Pool struct for postgresql connection.
type Pool struct {
	pool *pgxpool.Pool
}

// NewPGRepository create new postgresql repository.
func NewPGRepository(pool *pgxpool.Pool) (*Pool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	migrationDir := filepath.Join(dir, "/internal/migrations/000001_store.up.sql")

	file, err := os.ReadFile(migrationDir)
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(ctx, string(file)); err != nil {
		return nil, err
	}

	p := &Pool{
		pool: pool,
	}

	return p, nil
}

// Add
func (p *Pool) Add(ctx context.Context, short, long, id string) error {
	return nil
}

// Get
func (p *Pool) Get(ctx context.Context, short string) (string, error) {
	return "", nil
}

// GetAllByType
func (p *Pool) GetAllByType(ctx context.Context) (int, int, error) {
	return 0, 0, nil
}

// UpdateByID
func (p *Pool) UpdateByID(ctx context.Context, short string) (string, error) {
	return "", nil
}

// DeleteByID
func (p *Pool) DeleteByID(ctx context.Context) (int, int, error) {
	return 0, 0, nil
}
