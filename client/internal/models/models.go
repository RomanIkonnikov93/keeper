// Package models contains project data structures.
package models

import (
	"time"
)

var (
	BuildVersion = "1.0"
)

// Possible actions with a Record.
const (
	Add    = "add"
	Get    = "get"
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
	Login       string
	Password    string
	Token       string
	LastChanges string
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
	CreatedAt   string
	ActionType  string
}

// Storage struct for temporary storage of all user records (not including binary data).
type Storage struct {
	Credentials []Record
	Cards       []Record
}
