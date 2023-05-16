package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// UsersRepository interface for user repository methods.
type UsersRepository interface {
	AddUser(ctx context.Context, id, login, password string) error
	GetUser(ctx context.Context, login, password string) (string, string, error)
}

type UsersIDRepository struct {
	pool *pgxpool.Pool
}

// NewPGIDRepository create new postgresql user repository.
func NewPGIDRepository(pool *pgxpool.Pool) (*UsersIDRepository, error) {

	ctx, cancel := context.WithTimeout(context.Background(), models.TimeOut)
	defer cancel()

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	migrationDir := filepath.Join(dir, "/migrations/000001_users.up.sql")

	file, err := os.ReadFile(migrationDir)
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(ctx, string(file)); err != nil {
		return nil, err
	}

	return &UsersIDRepository{
		pool: pool,
	}, nil
}

// AddUser adds a new user and login details to the database.
func (p *UsersIDRepository) AddUser(ctx context.Context, ID, login, password string) error {

	c, cancel := context.WithTimeout(ctx, models.TimeOut)
	defer cancel()

	_, err := p.pool.Exec(c, `insert into keeper_auth (user_id, user_login, user_password) values ($1, $2, $3)`, ID, login, password)
	if err != nil {
		pgerr, ok := err.(*pgconn.PgError)
		if ok {
			if pgerr.Code == "23505" {
				return models.ErrConflict
			}
		}
		return err
	}

	return nil
}

// GetUser gets user data from database.
func (p *UsersIDRepository) GetUser(ctx context.Context, login, password string) (string, string, error) {

	user, pass, ID := "", "", ""
	err := p.pool.QueryRow(ctx, `select user_id, user_login, user_password from keeper_auth where user_login = $1`, login).
		Scan(&ID, &user, &pass)
	if err != nil {
		return "", "", models.ErrNotExist
	}

	return pass, ID, nil
}
