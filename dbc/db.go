package dbc

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid"
	"fmt"
	"math/rand"
	//"net/url"
	"time"
	"encoding/hex"
	"reflect"
	"strconv"
	"strings"
	"regexp"
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
	fmt.Println(t)
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

func plainFields(v reflect.Value) []string {
	fields := []string{}
	for i := 0; i < v.Type().NumField(); i++ {
		if v.Type().Field(i).Name == "Translations" {
			continue
		}
		if v.Type().Field(i).Anonymous {
			fields = append(fields, plainFields(v.FieldByName(v.Type().Field(i).Name))...)
		} else {
			fields = append(fields, v.Type().Field(i).Name)
		}
	}
	return fields
}

func Insert(db *sql.DB, record interface{}) int {
	now := time.Now()
	table := underscore(reflect.TypeOf(record).String())
	sql := "insert into \"" + table + "\" ("
	places := ""
	p := 1
	values := []interface{}{}
	v := reflect.Indirect(reflect.ValueOf(record))
	t := reflect.TypeOf(record)
	var field reflect.StructField
	var value reflect.Value
	fields := plainFields(v)
	for _, name := range fields {
		if name == "Translations" || name == "Locale" {
			continue
		}
		field, _ = t.FieldByName(name)
		tag := field.Tag.Get("dbc")
		fmt.Println(tag)
		match, _ := regexp.MatchString("translation", tag)
		if (match) {
			continue
		}
		value = v.FieldByName(name)
		if sql[len(sql)-1] != '(' {
			sql += ","
			places += ","
		}
		sql += "\"" + underscore(name) + "\""
		places += "$" + strconv.Itoa(p)
		p++
		if name == "CreatedAt" || name == "UpdatedAt" {
			values = append(values, now)
		} else if field.Type.String() == "ulid.ULID" {
			ulid, _ := ULID(now).MarshalBinary()
			uuid := UUID(ulid)
			values = append(values, uuid)
		} else {
			values = append(values, value.Interface())
		}
	}
	sql += ") values (" + places + ")"
	result, err := db.Exec(sql, values...)
	fmt.Println(result.RowsAffected())
	fmt.Println(err)
	return 1
}

