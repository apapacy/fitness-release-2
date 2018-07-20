package model

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/apapacy/fitness-release-2/dbc"
	// "github.com/apapacy/fitness-release-2/mo"
	// _ "github.com/lib/pq"
	//"github.com/apapacy/fitness-release-2/model"
	//"net/url"
	//"time"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

func TestCountryInsert(t *testing.T) {
	country := Countries{
		Code: sql.NullInt64{1, true},
		A2:   sql.NullString{"2", true},
		A3:   sql.NullString{"4", true},
		CountryTranslations: CountryTranslations{
			Locale: sql.NullString{"ua", true},
		},
	}
	res, err := dbc.GetDB().Exec("delete from countries")
	fmt.Println(res)
	fmt.Println(err)
	dbc.Insert(dbc.GetDB(), &country)
	fmt.Println("=========================================")
	fmt.Println(country)
	dbc.Select(dbc.GetDB(), &Countries{})
}
