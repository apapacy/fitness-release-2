package model

import (
	"github.com/oklog/ulid"

	//"fmt"
	//"net/url"
	//"time"

	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type City struct {
	ID           ulid.ULID
	Country      Countries
	Translations []CityTransaltions
	Locale       string
	Name         string
	Fullname     string
}

type CityTransaltions struct {
	Locale   string
	Name     string
	Fullname string
}
