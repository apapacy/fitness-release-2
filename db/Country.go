package db

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	//"net/url"
	//"time"
	"github.com/oklog/ulid"
	//"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type Countries struct {
	Id           ulid.ULID
	Code         sql.NullInt64  `unique:"true"`
	A2           sql.NullString `unique:"true"`
	A3           sql.NullString `unique:"true"`
	Translations []CountryTransaltions
	CountryTransaltions
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CountryTransaltions struct {
	Locale   sql.NullString
	Name     sql.NullString
	Fullname sql.NullString
}

func UUID(u []byte) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

func underscore(s string) string {
	name := ""
	for i := 0; i < len(s); i = i + 1 {
		if strings.ContainsAny(s[i:i+1], ".*") {
			name = ""
		} else if strings.ToLower(s[i:i+1]) != s[i:i+1] {
			if name == "" {
				name = strings.ToLower(s[i : i+1])
			} else {
				name += "_" + strings.ToLower(s[i:i+1])
			}
		} else {
			name += s[i : i+1]
		}
	}
	return name
}

func Save(db *sql.DB, record interface{}) int {
	table := underscore(reflect.TypeOf(record).String())
	sql := "insert into \"" + table + "\" ("
	places := ""
	p := 1
	values := []interface{}{}

	v := reflect.Indirect(reflect.ValueOf(record))

	for i := 0; i < v.Type().NumField(); i++ {
		fmt.Println(v.Type().Field(i).Type.String())
		if v.Type().Field(i).Name == "Translations" || v.Type().Field(i).Anonymous {
			continue
		}
		if sql[len(sql)-1] != '(' {
			sql += ","
			places += ","
		}
		sql += "\"" + underscore(v.Type().Field(i).Name) + "\""
		places += "$" + strconv.Itoa(p)
		p++
		if v.Type().Field(i).Type.String() == "ulid.ULID" {
			ulid, _ := v.FieldByName(v.Type().Field(i).Name).Interface().(ulid.ULID).MarshalBinary()
			uuid := UUID(ulid)
			values = append(values, uuid)
		} else {
			values = append(values, v.FieldByName(v.Type().Field(i).Name).Interface())
		}
	}
	sql += ") values (" + places + ")"
	fmt.Println("==================================")
	fmt.Println(sql)
	fmt.Println(places)
	fmt.Println(values)
	fmt.Println("==================================")
	result, err := db.Exec(sql, values...)
	fmt.Println(result)
	fmt.Println(err)
	return 1
}

func (this Countries) Save(db *sql.DB) {
	Save(db, this)
}

func init() {
	country := Countries{
		Id:   ULID(),
		Code: sql.NullInt64{123, true},
		A2:   sql.NullString{"12", true},
		A3:   sql.NullString{"123", true},
		CountryTransaltions: CountryTransaltions{
			Locale: sql.NullString{"ua", true},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	country.Save(GetDB())
	fmt.Println("111111111111111111111111")
	fmt.Println(country.Code.Value())
}

// https://gist.github.com/drewolson/4771479
