package dbc

import (
	"database/sql"
	"fmt"
	"math/rand"

	// _ "github.com/lib/pq"

	"github.com/oklog/ulid"
	"github.com/satori/go.uuid"

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
	addr  interface{}
}

func plainFields(v *reflect.Value) []structFields {
	fields := []structFields{}
	for i := 0; i < v.Type().NumField(); i++ {
		vp := v.FieldByName(v.Type().Field(i).Name)
		match, _ := regexp.MatchString("(^|,)re1f(,|$)", v.Type().Field(i).Tag.Get("dbc"))
		if v.Type().Field(i).Anonymous || match {
			fields = append(fields, plainFields(&vp)...)
		} else {
			field := structFields{
				name:  v.Type().Field(i).Name,
				ftype: v.Type().Field(i).Type,
				tag:   v.Type().Field(i).Tag,
				value: v.FieldByName(v.Type().Field(i).Name),
				addr:  &vp,
			}
			fields = append(fields, field)
		}
	}
	return fields
}

// Insert insert new recorde in database
func Insert(db *sql.DB, record interface{}) int {
	now := time.Now()
	table := underscore(reflect.TypeOf(record).String())
	isTranslations, _ := regexp.MatchString("_translations$", table)
	v := reflect.ValueOf(record).Elem()
	sql := "insert into \"" + table + "\" ("
	places := ""
	p := 1
	values := []interface{}{}
	fields := plainFields(&v)
	for _, field := range fields {
		if field.name == "Translations" {
			continue
		}
		tag := field.tag.Get("dbc")
		if !isTranslations {
			if field.name == "Locale" {
				continue
			}
			match, _ := regexp.MatchString("(^|,)translation(,|$)", tag)
			if match {
				continue
			}
		}
		if sql[len(sql)-1] != '(' {
			sql += ","
			places += ","
		}
		fmt.Println("//////////// " + tag)
		ref, _ := regexp.MatchString("(^|,)ref(,|$)", tag)
		fmt.Println(ref)
		if ref {
			sql += "\"" + underscore(field.name) + "_id\""
		} else {
			sql += "\"" + underscore(field.name) + "\""
		}
		places += "$" + strconv.Itoa(p)
		p++
		if field.name == "CreatedAt" || field.name == "UpdatedAt" {
			v.FieldByName(field.name).Set(reflect.ValueOf(now))
			values = append(values, now)
		} else if field.ftype.String() == "uuid.UUID" {
			match, _ := regexp.MatchString("(^|,)auto(,|$)", tag)
			if match {
				uid, _ := uuid.NewV1()
				fmt.Println(field.name)
				v.FieldByName(field.name).Set(reflect.ValueOf(uid))
			}
			values = append(values, field.value.Interface())
		} else {
			if ref {
				values = append(values, v.FieldByName(field.name).FieldByName("Id").Interface())
				fmt.Println("7777777777777777777777777777777777777")
				fmt.Println(v.FieldByName(field.name).FieldByName("Id").Interface())
			} else {
				values = append(values, field.value.Interface())
			}
		}
	}
	sql += ") values (" + places + ")"
	_, err := db.Exec(sql, values...)
	fmt.Println("+++++++++++++++++++++++++++++++++++++")
	fmt.Println(err)
	fmt.Println(sql)
	fmt.Println(values)
	return 1
}

func Select(db *sql.DB, records interface{}) {
	// now := time.Now()
	item := reflect.TypeOf(records).Elem().Elem()
	returnsPtr := reflect.ValueOf(records)
	returns := returnsPtr.Elem()

	//returns := reflect.New(reflect.ArrayOf(0, item)).Elem()
	table := underscore(item.String())
	r := reflect.New(item).Elem()
	fmt.Println("qqqqqqqqqqqqqqqqqqqq")
	fmt.Println(r)
	translations_table := " left join \"" + table + "_translations\" on \"" + table + "\".\"id\"=\"" + table + "_translations\".\"id\""
	from := table
	sql := "select "
	values := []interface{}{}
	//v := reflect.ValueOf(r)
	fields := plainFields(&r)
	for _, field := range fields {
		if field.name == "Translations" {
			from = table + translations_table
			continue
		}
		if sql[len(sql)-1] != ' ' {
			sql += ","
		}
		tag := field.tag.Get("dbc")
		match, _ := regexp.MatchString("(^|,)translation(,|$)", tag)
		if match {
			sql += "\"" + table + "_translations\".\"" + underscore(field.name) + "\""
		} else {
			sql += "\"" + table + "\".\"" + underscore(field.name) + "\""
		}
		fmt.Println("eeeeeeeeeeeeeeeeeeeeeeee")
		fmt.Println(r.FieldByName(field.name).Addr())
		values = append(values, r.FieldByName(field.name).Addr().Interface())
	}

	sql += " from " + from
	fmt.Println(sql)
	row, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		for row.Next() {
			for _, ref := range values {
				r := reflect.ValueOf(ref).Elem()
				r.Set(reflect.Zero(r.Type()))
			}
			row.Scan(values...)
			fmt.Println(r)
			fmt.Println("11111111111111")
			returns.Set(reflect.Append(returns, r))

			//returns = append(returns, r)
			fmt.Println(returns)
			fmt.Println("222222222222")
		}
	}
}
