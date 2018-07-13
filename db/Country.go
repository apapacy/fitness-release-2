package db

import (
	_ "github.com/lib/pq"
	//"fmt"
	//"net/url"
	//"time"
	"github.com/oklog/ulid"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Country struct {
	Id           ulid.ULID
	Code         int    `unique:"true"`
	A2           string `unique:"true"`
	A3           string `unique:"true"`
	Translations []CountryTransaltions
	CountryTransaltions
}

type CountryTransaltions struct {
	Locale   string
	Name     string
	Fullname string
}
