package model

import (
  "testing"
	_ "github.com/lib/pq"
	"fmt"
	//"net/url"
	//"time"

	"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

func TestCountryInsert(t *testing.T) {
  countryStore := NewCountryStore(GetDB())
  country := Country{
    ID: kallax.NewULID(),
    Code: 36,
    A2: "AU",
    A3: "AUS",
    Translations: []CountryTransaltions{{
        ID: kallax.NewULID(),
        Locale: "ua",
        Name: "Австрія",
        Fullname: "Австрія",
      },
    },
  }

  err := countryStore.Insert(&country)
  if err != nil {
          fmt.Println(err)
  }

}


func TestCountrySelect(t *testing.T) {
  countryStore := NewCountryStore(GetDB())
  rs, err := countryStore.Find(NewCountryQuery().WithTranslations(kallax.Eq(Schema.CountryTransaltions.Locale, "ua")).FindByA2("AU"))
  if err != nil {
          fmt.Println(err)
  }

  for rs.Next() {
          country, err := rs.Get()
          fmt.Println(country)
          if err != nil {
                  fmt.Println(err)
          }
  }

}
