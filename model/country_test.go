package model

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	//"net/url"
	//"time"

	"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

func TestCountryInsert(t *testing.T) {
	countryStore := NewCountryStore(GetDB())
	country := Country{
		ID:   kallax.NewULID(),
		Code: 36,
		A2:   "AU",
		A3:   "AUS",
		Translations: []CountryTransaltions{{
			ID:       kallax.NewULID(),
			Locale:   "ua",
			Name:     "Австрія",
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

func TestCityInsert(t *testing.T) {
	countryStore := NewCountryStore(GetDB())
	cityStore := NewCityStore(GetDB())

	country, err := countryStore.FindOne(NewCountryQuery().FindByA2("AU"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(country)

	city := City{
		ID:      kallax.NewULID(),
		Country: *country,
		Translations: []CityTransaltions{{
			ID:       kallax.NewULID(),
			Locale:   "ua",
			Name:     "Пенза",
			Fullname: "Пенза",
		},
		},
	}
	err = cityStore.Insert(&city)
	fmt.Println(err)

}

func TestCitySelect(t *testing.T) {
	cityStore := NewCityStore(GetDB())
	city, err := cityStore.
		FindOne(NewCityQuery().
			WithTranslations(kallax.Eq(Schema.CityTransaltions.Locale, "ua")).
			Where(kallax.Eq(Schema.CityTransaltions.Name, "Пенза")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(city)

}
