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
	country1 := Country{
		Code: sql.NullInt64{1, true},
		A2:   sql.NullString{"2", true},
		A3:   sql.NullString{"3", true},
		CountryTranslations: CountryTranslations{
			Locale: sql.NullString{"ua", true},
		},
	}
	res, err := dbc.GetDB().Exec("delete from city_translations;delete from city;delete from country_translations;delete from country;")
	fmt.Println(res)
	fmt.Println(err)
	res, err = dbc.Insert(dbc.GetDB(), &country1)
	if err != nil {
		panic(err)
	}
	country1Translarions := CountryTranslations{
		Fullname: sql.NullString{"wewe", true},
		Name:     sql.NullString{"wewe", true},
		Locale:   sql.NullString{"ua", true},
		Id:       country1.Id,
	}
	dbc.Insert(dbc.GetDB(), &country1Translarions)
	//fmt.Println("=========================================")
	//fmt.Println(country)
	city := City{
		Country: &country1,
	}
	fmt.Println("=========================================")
	dbc.Insert(dbc.GetDB(), &city)
	fmt.Println(city)
	country2 := Country{
		Code:    sql.NullInt64{2, true},
		A2:      sql.NullString{"3", true},
		A3:      sql.NullString{"4", true},
		Capital: &city,
		CountryTranslations: CountryTranslations{
			Locale: sql.NullString{"ua", true},
		},
	}
	res, err = dbc.Insert(dbc.GetDB(), &country2)
	if err != nil {
		panic(err)
	}
	countries := []Country{}
	_, err = dbc.Select(dbc.GetDB(), &countries)
	if err != nil {
		panic(err)
	}
	fmt.Println("444444444444444444444444444444444444")
	fmt.Println(countries)
	fmt.Println(*(*countries[0].Capital).Country)
	fmt.Println(*(*countries[1].Capital).Country)
}
