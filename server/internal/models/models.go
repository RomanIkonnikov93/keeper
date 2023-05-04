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
	ErrConflict    = errors.New("conflict on insert")
	ErrDelFlag     = errors.New("record is deleted")
	ErrNotExist    = errors.New("not exist")
	ErrInvalidData = errors.New("invalid data")
	ErrNotValid    = errors.New("token not valid")
)

// Users data types.
const (
	Credentials = "users_credentials"
	File        = "users_files"
	Card        = "users_cards"
)

// Query for repository.
var (
	QueryAddCredentials = `insert into users_credentials (user_id,description,metadata,user_login,user_password) values ($1, $2, $3, $4, $5)`
	QueryAddCard        = `insert into users_cards (user_id,description,metadata,user_card) values ($1, $2, $3, $4)`
	QueryAddFile        = `insert into users_files (user_id,description,metadata,user_file) values ($1, $2, $3, $4)`

	QueryUpdateCredentials = `update users_credentials set description = $1, metadata = $2, user_login = $3, user_password= $4 where user_id = $5 and record_id = $6`
	QueryUpdateCard        = `update users_cards set description = $1, metadata = $2, user_card = $3 where user_id = $4 and record_id = $5`
	QueryUpdateFile        = `update users_files set description = $1, metadata = $2, user_file = $3 where user_id = $4 and record_id = $5`
)
