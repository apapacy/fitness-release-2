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
		Code: sql.NullInt64{234, true},
		A2:   sql.NullString{"23", true},
		A3:   sql.NullString{"", false},
		CountryTranslations: CountryTranslations{
			Locale: sql.NullString{"ua", true},
		},
	}
	res, err := dbc.GetDB().Exec("delete from countries")
	fmt.Println(res)
	fmt.Println(err)
	country.Insert(dbc.GetDB())
	s := Countries{}
	dbc.Select(dbc.GetDB(), &s)
}
