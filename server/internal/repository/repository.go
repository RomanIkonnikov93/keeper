// Package repository contains storage options for links and short links.
package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository interface for repository methods.
type Repository interface {
	Add(ctx context.Context, record *pb.Record) error
	Get(ctx context.Context, short string) (string, error)
	GetAllByType(ctx context.Context) (int, int, error)
	UpdateByID(ctx context.Context, in *pb.Record) error
	DeleteByID(ctx context.Context, in *pb.Record) error
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
func (p *Pool) Add(ctx context.Context, in *pb.Record) error {

	switch in.RecordType {

	case models.Credentials:
		_, err := p.pool.Exec(ctx, models.QueryAddCredentials, in.UserID, in.Description, in.Metadata, in.Login, in.Password)
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

	case models.Card:
		_, err := p.pool.Exec(ctx, models.QueryAddCard, in.UserID, in.Description, in.Metadata, in.Card)
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

	case models.File:
		_, err := p.pool.Exec(ctx, models.QueryAddFile, in.UserID, in.Description, in.Metadata, in.File)
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

	default:

		return models.ErrInvalidData
	}
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
func (p *Pool) UpdateByID(ctx context.Context, in *pb.Record) error {

	switch in.RecordType {

	case models.Credentials:
		_, err := p.pool.Exec(ctx, models.QueryUpdateCredentials, in.Description, in.Metadata, in.Login, in.Password, in.UserID, in.RecordID)
		if err != nil {
			return err
		}

		return nil

	case models.Card:
		_, err := p.pool.Exec(ctx, models.QueryUpdateCard, in.Description, in.Metadata, in.Card, in.UserID, in.RecordID)
		if err != nil {
			return err
		}

		return nil

	case models.File:
		_, err := p.pool.Exec(ctx, models.QueryUpdateFile, in.Description, in.Metadata, in.File, in.UserID, in.RecordID)
		if err != nil {
			return err
		}

		return nil

	default:

		return models.ErrInvalidData
	}
}

// DeleteByID
func (p *Pool) DeleteByID(ctx context.Context, in *pb.Record) error {

	query := "update " + in.RecordType + " set del_flag=true where user_id=$1 and record_id = $2"

	_, err := p.pool.Exec(ctx, query, in.UserID, in.RecordID)
	if err != nil {
		return err
	}

	return nil
}
