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
	ErrDelFlag  = errors.New("record is deleted")
	ErrNotExist = errors.New("not exist")
)

// Users data types.
const (
	Credentials = "users_credentials"
	File        = "users_files"
	Card        = "users_cards"
)
