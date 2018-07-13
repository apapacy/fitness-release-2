package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid"
	"fmt"
	"math/rand"
	//"net/url"
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

func ULID() {
	t := time.Now()
	fmt.Println(t)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
}

func init() {
	ULID()
}