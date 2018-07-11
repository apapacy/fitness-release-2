package model

import (
	_ "github.com/lib/pq"
	//"fmt"
	//"net/url"
	//"time"

	"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type City struct {
	kallax.Model `table:"cities" pk:"id"`
	kallax.Timestamps
	ID           kallax.ULID
	Country      Country `fk:",inverse"`
	Translations []CityTransaltions
	Locale       string `kallax:"-"`
	Name         string `kallax:"-"`
	Fullname     string `kallax:"-"`
}

type CityTransaltions struct {
	kallax.Model `table:"cities_translations" pk:"id"`
	kallax.Timestamps
	ID       kallax.ULID
	Locale   string
	Name     string
	Fullname string
}
