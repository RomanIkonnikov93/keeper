package repository

import (
	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/conn"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"
)

// Reps struct for NewReps.
type Reps struct {
	Rep     Repository
	UserRep IDRepository
	Ping    Pinger
}

// NewReps creates new repositories and selects their type.
func NewReps(cfg config.Config) (*Reps, error) {

	R := Reps{}

	logger := logging.GetLogger()

	pool := conn.NewConnection(cfg)

	userRep, err := NewPGIDRepository(pool)
	if err != nil {
		return nil, err
	}

	rep, err := NewPGRepository(pool)
	if err != nil {
		return nil, err
	}

	p, err := NewPing(cfg)
	if err != nil {
		return nil, err
	}

	R.Rep = rep
	R.UserRep = userRep
	R.Ping = p

	logger.Println("connection to database")

	return &R, nil
}
