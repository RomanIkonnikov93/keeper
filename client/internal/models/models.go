package models

import (
	"errors"
	"time"
)

var (
	BuildVersion = "1.0"
)

var (
	ErrNotExist    = errors.New("not exist")
	ErrMaxFileSize = errors.New("file size exceeded")
)

// Possible actions with a Record.
const (
	Add    = "add"
	Get    = "get"
	GetAll = "getAll"
	Update = "update"
	Delete = "delete"
)

// Ticker for func tui.StartScanningChanges.
const Ticker = time.Second * 10

// DefaultLastChangesTime and layout for time.Parse.
const DefaultLastChangesTime = "2006-01-02 15:04:05 +0000 UTC"

// Users data types.
const (
	Credentials = "users_credentials"
	Card        = "users_cards"
	File        = "users_files"
)

// Auth struct for user authorization and authentication.
type Auth struct {
	Login           string
	Password        string
	Token           string
	LastChangesTime string
}

// Record to store input user data before sending it to the server.
type Record struct {
	RecordID    int32
	RecordType  string
	Description string
	Metadata    string
	Login       string
	Password    string
	Card        string
	File        []byte
	FilePath    string
	CreatedAt   string
	ActionType  string
}

// Storage struct for temporary storage of all user records (not including binary data).
type Storage struct {
	Credentials map[int32]Record
	Cards       map[int32]Record
	FileInfo    map[int32]Record
}
