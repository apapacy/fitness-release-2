package db

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/lib/pq"
	//"net/url"
	//"time"
	"github.com/oklog/ulid"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Country struct {
	Id           ulid.ULID
	Code         sql.NullInt64  `unique:"true"`
	A2           sql.NullString `unique:"true"`
	A3           sql.NullString `unique:"true"`
	Translations []CountryTransaltions
	CountryTransaltions
}

type CountryTransaltions struct {
	Locale   sql.NullString
	Name     sql.NullString
	Fullname sql.NullString
}

func Save(db *sql.DB, record interface{}) int {
	table := reflect.TypeOf(record).String()
	fmt.Println("***>>>" + table)
	fmt.Println(record)

	v := reflect.Indirect(reflect.ValueOf(record))

	for i := 0; i < v.Type().NumField(); i++ {
		fmt.Println("-------------")
		fmt.Println(v.Type().Field(i).Name)
		fmt.Println(v.Type().Field(i).Anonymous)
		fmt.Println(v.Field(i))
		if v.Field(i).Type() == reflect.TypeOf((*sql.NullInt64)(nil)).Elem() {
			fmt.Println(v.Field(i).FieldByName("Value"))
		}

	}

	return 1
}

func (this Country) Save(db *sql.DB) {
	Save(db, this)
}

func init() {
	country := Country{
		Id: ULID(),
		// Code: sql.NullInt64{123, true},
		CountryTransaltions: CountryTransaltions{
			Locale: sql.NullString{"ua", true},
		},
	}
	country.Save(GetDB())
	fmt.Println("111111111111111111111111")
	fmt.Println(country.Code.Value())
}
