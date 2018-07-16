package model

import (
	"database/sql"
	//"time"

	_ "github.com/lib/pq"
	//"net/url"
	//"time"
	"github.com/oklog/ulid"
	"github.com/apapacy/fitness-release-2/dbc"

	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Countries struct {
	Id           ulid.ULID `dbc:"pk"`
	Code         sql.NullInt64
	A2           sql.NullString
	A3           sql.NullString
	dbc.Timestamp
	Translations []CountryTranslations
	CountryTranslations
}

type CountryTranslations struct {
	Locale   sql.NullString `dbc:"locale"`
	Name     sql.NullString `dbc:"translation"`
	Fullname sql.NullString `dbc:"translation"`
}


func (this Countries) Insert(db *sql.DB) {
	dbc.Insert(db, this)
}


// https://gist.github.com/drewolson/4771479
