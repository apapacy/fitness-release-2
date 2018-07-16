package dbc

import (
	"time"
)

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Paranoid struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}