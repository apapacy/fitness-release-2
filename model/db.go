//go:generate kallax gen
package model

import (
	_ "github.com/lib/pq"
	//"fmt"
	//"net/url"
	//"time"

	"gopkg.in/src-d/go-kallax.v1"
	//"gopkg.in/src-d/go-kallax.v1/tests/fixtures"
	//"gopkg.in/src-d/go-kallax.v1/types"
)

type User struct {
	kallax.Model `table:"users" pk:"id"`
	ID           kallax.ULID
	Username     string
	Email        string
	Password     string
}
