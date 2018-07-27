package model

import (
	"database/sql"
	//"time"

	_ "github.com/lib/pq"
	//"net/url"
	//"time"
	// "github.com/oklog/ulid"
	"github.com/satori/go.uuid"

	"github.com/apapacy/fitness-release-2/dbc"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Country struct {
	Id           uuid.UUID `dbc:"pk,auto"`
	Code         sql.NullInt64
	A2           sql.NullString
	A3           sql.NullString
	Translations []CountryTranslations
	CountryTranslations
}

type CountryTranslations struct {
	Id       uuid.UUID      `dbc:"translation"`
	Locale   sql.NullString `dbc:"locale,translation"`
	Name     sql.NullString `dbc:"translation"`
	Fullname sql.NullString `dbc:"translation"`
	dbc.Timestamp
}

// https://gist.github.com/drewolson/4771479
