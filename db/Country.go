package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"reflect"
	//"net/url"
	//"time"
	"github.com/oklog/ulid"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Country struct {
	Id           ulid.ULID
	Code         int    `unique:"true"`
	A2           string `unique:"true"`
	A3           string `unique:"true"`
	Translations []CountryTransaltions
	CountryTransaltions
}

type CountryTransaltions struct {
	Locale   string
	Name     string
	Fullname string
}


func Save(db *sql.DB, record interface{}) int {
	table := reflect.TypeOf(record).String()
	fmt.Println("***>>>" + table)
	fmt.Println(record)

	v := reflect.ValueOf(record)

    values := make([]interface{}, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        values[i] = v.Field(i).Interface()
		fmt.Println(values[i])
	}

	return 1;
}

func (this Country) Save(db *sql.DB) {
	Save(db, this)
}

func init() {
	country := Country{
		Id: ULID(),
		CountryTransaltions: CountryTransaltions{
			Locale: "ua",
		},
	}
	country.Save(GetDB())
}