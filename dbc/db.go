package dbc

import (
	"database/sql"
	"fmt"
	"math/rand"

	// _ "github.com/lib/pq"
	"github.com/oklog/ulid"

	//"net/url"
	"encoding/hex"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error
	if db != nil {
		return db
	}
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost:5433/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func ULID(t time.Time) ulid.ULID {
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newUlid := ulid.MustNew(ulid.Timestamp(t), entropy)
	return newUlid
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

type structFields struct {
	name  string
	ftype reflect.Type
	tag   reflect.StructTag
	value reflect.Value
}

func plainFields(v reflect.Value) []structFields {
	fields := []structFields{}
	for i := 0; i < v.Type().NumField(); i++ {
		field := structFields{
			name:  v.Type().Field(i).Name,
			ftype: v.Type().Field(i).Type,
			tag:   v.Type().Field(i).Tag,
			value: v.FieldByName(v.Type().Field(i).Name),
		}
		if v.Type().Field(i).Anonymous {
			fields = append(fields, plainFields(v.FieldByName(v.Type().Field(i).Name))...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

// Insert insert new recorde in database
func Insert(db *sql.DB, record interface{}) int {
	now := time.Now()
	table := underscore(reflect.TypeOf(record).String())
	sql := "insert into \"" + table + "\" ("
	places := ""
	p := 1
	values := []interface{}{}
	v := reflect.Indirect(reflect.ValueOf(record))
	fields := plainFields(v)
	for _, field := range fields {
		if field.name == "Translations" || field.name == "Locale" {
			continue
		}
		tag := field.tag.Get("dbc")
		match, _ := regexp.MatchString("translation", tag)
		if match {
			continue
		}
		value := field.value
		if sql[len(sql)-1] != '(' {
			sql += ","
			places += ","
		}
		sql += "\"" + underscore(field.name) + "\""
		places += "$" + strconv.Itoa(p)
		p++
		if field.name == "CreatedAt" || field.name == "UpdatedAt" {
			values = append(values, now)
		} else if field.ftype.String() == "ulid.ULID" {
			ulid, _ := ULID(now).MarshalBinary()
			uuid := UUID(ulid)
			values = append(values, uuid)
		} else {
			values = append(values, value.Interface())
		}
	}
	sql += ") values (" + places + ")"
	result, err := db.Exec(sql, values...)
	fmt.Println(result)
	fmt.Println(err)
	return 1
}
