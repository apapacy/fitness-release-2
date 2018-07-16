package dbc

import (
	"time"
)

// Timestamp - common  fiedls in database table
type Timestamp struct {
	CreatedAt time.Time `dbc:"required"`
	UpdatedAt time.Time `dbc:"required"`
}

// Paranoid common timestamp fiedls in database table with
type Paranoid struct {
	DeletedAt time.Time
}