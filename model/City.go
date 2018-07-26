package model

import (
	"github.com/apapacy/fitness-release-2/dbc"
	"github.com/satori/go.uuid"
	//"fmt"
	//"net/url"
	//"time"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type City struct {
	Id           uuid.UUID `dbc:"pk,auto"`
	Country      Country   `dbc:"ref"`
	Translations []CityTransaltions
	CityTransaltions
}

type CityTransaltions struct {
	Locale   string `dbc:"locale"`
	Name     string `dbc:"translation"`
	Fullname string `dbc:"translation"`
	dbc.Timestamp
}
