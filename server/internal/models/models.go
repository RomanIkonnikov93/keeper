// Package models contains project data structures.
package models

import (
	"errors"
	"time"
)

// TimeOut for ctx.
const TimeOut = time.Second * 5

// Repository errors.
var (
	ErrConflict = errors.New("conflict on insert")
	ErrDelFlag  = errors.New("url is deleted")
)
