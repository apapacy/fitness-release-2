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

type Country struct {
	kallax.Model `table:"countries" pk:"id"`
	kallax.Timestamps
	ID kallax.ULID
	Code int  `unique:"true"`
	A2 string `unique:"true"`
	A3 string `unique:"true"`
	Translations []CountryTransaltions
	Locale string `kallax:"-"`
	Name string `kallax:"-"`
	Fullname string `kallax:"-"`
}

type CountryTransaltions struct {
	kallax.Model `table:"countries_translations" pk:"id"`
	kallax.Timestamps
	ID  kallax.ULID
	Locale string
	Name string
	Fullname string
}
