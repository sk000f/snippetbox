package models

import (
	"errors"
	"time"
)

// ErrNoRecord defines an error returned when no matching record is found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is a type representing Snippet values from the DB
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
