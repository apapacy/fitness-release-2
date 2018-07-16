package dbc

import (
	//"fmt"
	"testing"
	"database/sql"

	// _ "github.com/lib/pq"
	"github.com/apapacy/fitness-release-2/model"
	//"net/url"
	//"time"

	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

func TestCountryInsert(t *testing.T) {
	country := model.Countries{
		Code: sql.NullInt64{234, true},
		A2:   sql.NullString{"23", true},
		A3:   sql.NullString{"234", true},
		CountryTranslations: model.CountryTranslations{
			Locale: sql.NullString{"ua", true},
		},
	}

	country.Save(GetDB())
}

