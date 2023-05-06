// Package repository contains storage options for links and short links.
package repository

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository interface for repository methods.
type Repository interface {
	Add(ctx context.Context, record *pb.Record) error
	Get(ctx context.Context, in *pb.Record) (*pb.Record, error)
	GetAllByType(ctx context.Context, in *pb.Record) ([]models.Record, error)
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
func (p *Pool) Get(ctx context.Context, in *pb.Record) (*pb.Record, error) {

	var (
		createdAt time.Time
		flag      bool
	)

	switch in.RecordType {

	case models.Credentials:
		rows, err := p.pool.Query(ctx, models.QueryGetCredentials, in.UserID, in.RecordID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			if err = rows.Scan(&in.Description, &in.Metadata, &in.Login, &in.Password, &flag, &createdAt); err != nil {
				return nil, err
			}
		}
		if flag {
			return nil, models.ErrDelFlag
		}

		in.CreatedAt = createdAt.String()
		in.UserID = ""

		return in, nil

	case models.Card:
		rows, err := p.pool.Query(ctx, models.QueryGetCard, in.UserID, in.RecordID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			if err = rows.Scan(&in.Description, &in.Metadata, &in.Card, &flag, &createdAt); err != nil {
				return nil, err
			}
		}
		if flag {
			return nil, models.ErrDelFlag
		}

		in.CreatedAt = createdAt.String()
		in.UserID = ""

		return in, nil

	case models.File:
		rows, err := p.pool.Query(ctx, models.QueryGetFile, in.UserID, in.RecordID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			if err = rows.Scan(&in.Description, &in.Metadata, &in.File, &flag, &createdAt); err != nil {
				return nil, err
			}
		}
		if flag {
			return nil, models.ErrDelFlag
		}

		in.CreatedAt = createdAt.String()
		in.UserID = ""

		return in, nil

	default:

		return nil, models.ErrInvalidData
	}
}

// GetAllByType
func (p *Pool) GetAllByType(ctx context.Context, in *pb.Record) ([]models.Record, error) {

	var (
		createdAt time.Time
	)

	out := make([]models.Record, 0)

	switch in.RecordType {

	case models.Credentials:
		rows, err := p.pool.Query(ctx, models.QueryGetAllCredentials, in.UserID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			if err = rows.Scan(&in.RecordID, &in.Description, &in.Metadata, &in.Login, &in.Password, &createdAt); err != nil {
				return nil, err
			}

			record := models.Record{
				RecordID:    in.RecordID,
				RecordType:  in.RecordType,
				Description: in.Description,
				Metadata:    in.Metadata,
				Login:       in.Login,
				Password:    in.Password,
				CreatedAt:   createdAt.String(),
			}

			out = append(out, record)
		}

		if len(out) < 1 {
			return nil, models.ErrNotExist
		}

		return out, nil

	case models.Card:
		rows, err := p.pool.Query(ctx, models.QueryGetAllCard, in.UserID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			if err = rows.Scan(&in.RecordID, &in.Description, &in.Metadata, &in.Card, &createdAt); err != nil {
				return nil, err
			}

			record := models.Record{
				RecordID:    in.RecordID,
				RecordType:  in.RecordType,
				Description: in.Description,
				Metadata:    in.Metadata,
				Card:        in.Card,
				CreatedAt:   createdAt.String(),
			}

			out = append(out, record)
		}

		if len(out) < 1 {
			return nil, models.ErrNotExist
		}

		return out, nil

	case models.File:
		rows, err := p.pool.Query(ctx, models.QueryGetAllFile, in.UserID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			if err = rows.Scan(&in.RecordID, &in.Description, &in.Metadata, &in.File, &createdAt); err != nil {
				return nil, err
			}

			record := models.Record{
				RecordID:    in.RecordID,
				RecordType:  in.RecordType,
				Description: in.Description,
				Metadata:    in.Metadata,
				File:        in.File,
				CreatedAt:   createdAt.String(),
			}

			out = append(out, record)
		}

		if len(out) < 1 {
			return nil, models.ErrNotExist
		}

		return out, nil

	default:

		return nil, models.ErrInvalidData
	}
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
