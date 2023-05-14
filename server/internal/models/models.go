// Package models contains project data structures.
package models

import (
	"errors"
	"time"
)

const (
	TimeOut    = time.Second * 5
	TimeFormat = "2006-01-02 15:04:05 +0000 UTC"
)

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
	Card        = "users_cards"
	File        = "users_files"
)

type Record struct {
	RecordID    int32
	RecordType  string
	Description string
	Metadata    string
	Login       string
	Password    string
	Card        string
	File        []byte
	CreatedAt   string
}

// Query for repository.
var (
	QueryAddCredentials = `insert into users_credentials (user_id,description,metadata,user_login,user_password) values ($1, $2, $3, $4, $5)`
	QueryAddCard        = `insert into users_cards (user_id,description,metadata,user_card) values ($1, $2, $3, $4)`
	QueryAddFile        = `insert into users_files (user_id,description,metadata,user_file) values ($1, $2, $3, $4)`

	QueryGetCredentials = `select description, metadata, user_login, user_password, del_flag, created_at from users_credentials where user_id = $1 and record_id = $2`
	QueryGetCard        = `select description, metadata, user_card, del_flag, created_at from users_cards where user_id = $1 and record_id = $2`
	QueryGetFile        = `select description, metadata, user_file, del_flag, created_at from users_files where user_id = $1 and record_id = $2`

	QueryGetAllCredentials = `select record_id, description, metadata, user_login, user_password, created_at from users_credentials where user_id = $1 and del_flag = false`
	QueryGetAllCard        = `select record_id, description, metadata, user_card, created_at from users_cards where user_id = $1 and del_flag = false`
	QueryGetAllFile        = `select record_id, description, metadata, user_file, created_at from users_files where user_id = $1 and del_flag = false`

	QueryCheckChangesCredentials = `select record_id, description, metadata, user_login, user_password, created_at from users_credentials where user_id = $1 and del_flag = false and created_at > $2`
	QueryCheckChangesCard        = `select record_id, description, metadata, user_card, created_at from users_cards where user_id = $1 and del_flag = false and created_at > $2`
)
