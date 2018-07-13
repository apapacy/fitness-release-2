package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/src-d/go-kallax.v1"

	"fmt"
	//"net/url"
	//"time"
	"github.com/apapacy/fitness-release-2/model"
	dbn "github.com/apapacy/fitness-release-2/db"

	"github.com/apapacy/fitness-release-2/routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5433/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	userStore := model.NewUserStore(db)
	user := model.User{
		ID:       kallax.NewULID(),
		Username: "john",
		Email:    "john@doe.me",
		Password: "1234bunnies",
	}
	err = userStore.Insert(&user)
	if err != nil {
		panic(err)
	}
	id :=  user.ID
	time := uint64(id[5]) | uint64(id[4])<<8 |
		uint64(id[3])<<16 | uint64(id[2])<<24 |
		uint64(id[1])<<32 | uint64(id[0])<<40
	fmt.Print(time)
	dbn.ULID()
	routes.GetRouter().Run() // listen and serve on 0.0.0.0:8080
}
