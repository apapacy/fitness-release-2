//go:generate kallax gen
package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	//"fmt"
	//"net/url"
	//"time"

	"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
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

type User struct {
	kallax.Model `table:"users" pk:"id"`
	ID           kallax.ULID
	Username     string
	Email        string
	Password     string
}

