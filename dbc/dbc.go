package dbc

import (
	"database/sql"
	"fmt"

	"encoding/hex"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
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

func plainFields(vp *reflect.Value) []structFields {
	fields := []structFields{}
	v := vp.Elem()
	for i := 0; i < v.Type().NumField(); i++ {
		vp := v.FieldByName(v.Type().Field(i).Name).Addr()
		match, _ := regexp.MatchString("(^|,)re1f(,|$)", v.Type().Field(i).Tag.Get("dbc"))
		if v.Type().Field(i).Anonymous || match {
			fields = append(fields, plainFields(&vp)...)
		} else {
			field := structFields{
				name:  v.Type().Field(i).Name,
				ftype: v.Type().Field(i).Type,
				tag:   v.Type().Field(i).Tag,
				value: v.FieldByName(v.Type().Field(i).Name),
				addr:  v.FieldByName(v.Type().Field(i).Name).Addr().Interface(),
			}
			fields = append(fields, field)
		}
	}
	return fields
}

// Insert insert new recorde in database
func Insert(db *sql.DB, record interface{}) (sql.Result, error) {
	now := time.Now()
	table := underscore(reflect.TypeOf(record).String())
	isTranslations, _ := regexp.MatchString("_translations$", table)
	vp := reflect.ValueOf(record)
	v := vp.Elem()
	sql := "insert into \"" + table + "\" ("
	places := ""
	p := 1
	values := []interface{}{}
	fields := plainFields(&vp)
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
		ref, _ := regexp.MatchString("(^|,)ref(,|$)", tag)
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
				v.FieldByName(field.name).Set(reflect.ValueOf(uid))
			}
			values = append(values, field.value.Interface())
		} else {
			if ref {
				if v.FieldByName(field.name).IsNil() {
					values = append(values, nil)
				} else {
					values = append(values, v.FieldByName(field.name).Elem().FieldByName("Id").Interface())
				}
			} else {
				values = append(values, field.value.Interface())
			}
		}
	}
	sql += ") values (" + places + ")"
	result, err := db.Exec(sql, values...)
	fmt.Println(sql)
	return result, err
}

func Select(db *sql.DB, records interface{}) (*sql.Rows, error) {
	returns := reflect.ValueOf(records).Elem()
	element := reflect.TypeOf(records).Elem().Elem()
	table := underscore(element.String())
	fields := []structFields{}
	values := []interface{}{}
	tables := ""
	sqlFields := ""
	prepareSelect(records, "", table, &tables, &sqlFields, &fields, &values, 0)
	// newElement := newElementPtr.Elem()
	sql := "select " + sqlFields + " from " + tables
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			fields := []structFields{}
			values := []interface{}{}
			tables := ""
			sqlFields := ""
			newElementPtr := prepareSelect(records, "", table, &tables, &sqlFields, &fields, &values, 0)
			newElement := newElementPtr.Elem()
			rows.Scan(values...)
			fmt.Println(values)

			returns.Set(reflect.Append(returns, newElement))
		}
	}
	fmt.Println()

	return rows, err
}

func prepareSelect(records interface{}, prefix string, alias string, tables *string, sqlFields *string, allFields *[]structFields, allValues *[]interface{}, level int) reflect.Value {
	fmt.Println("-------------------------------------")
	fmt.Println(prefix)
	element := reflect.TypeOf(records).Elem().Elem()
	newElementPtr := reflect.New(element)
	//newElement := newElementPtr.Elem()
	var prefixed string
	if prefix != "" {
		prefixed = prefix + "__"
	}
	table := underscore(element.String())
	translationsTable := " left join \"" + table + "_translations\" \"" + prefixed + alias + "_translations\" on \"" + prefixed + alias + "\".\"id\"=\"" + prefixed + alias + "_translations\".\"id\""
	var from string
	if level == 0 {
		from = table
	} else {
		from = " left join \"" + table + "\" \"" + prefixed + alias + "\" on \"" + prefixed + alias + "\".\"id\"=\"" + prefix + "\".\"" + alias + "_id\" "
	}
	// values := []interface{}{}
	fields := plainFields(&newElementPtr)
	for _, field := range fields {
		if field.name == "Translations" {
			from = from + translationsTable
			break
		}
	}
	*tables += " " + from
	for _, field := range fields {
		if field.name == "Translations" {
			continue
		}
		tag := field.tag.Get("dbc")
		match, _ := regexp.MatchString("(^|,)ref(,|$)", tag)
		if match {
			if level < 4 {
				element := prepareSelect(field.addr, prefixed+alias, underscore(field.name), tables, sqlFields, allFields, allValues, level+1)
				field.value.Set(element)
			}
			continue
		}
		if len(*sqlFields) != 0 {
			*sqlFields += ","
		}
		match, _ = regexp.MatchString("(^|,)translation(,|$)", tag)
		if match {
			*sqlFields += "\"" + prefixed + alias + "_translations\".\"" + underscore(field.name) + "\" as \"" + prefixed + alias + "_translations_" + underscore(field.name) + "\""
		} else {
			*sqlFields += "\"" + prefixed + alias + "\".\"" + underscore(field.name) + "\" as \"" + prefixed + alias + "_" + underscore(field.name) + "\""
		}
		*allValues = append(*allValues, field.addr)
	}
	return newElementPtr
}
