package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord defines an error returned when no matching record is found
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials indicates a failed authentication attempt
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail indicates a failed attempt to signup with an exisitng email address
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet is a type representing Snippet values from the DB
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User represents an authenticated system user
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
