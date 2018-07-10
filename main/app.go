package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/src-d/go-kallax.v1"

	//"fmt"
	//"net/url"
	//"time"
	"github.com/apapacy/fitness-release-2/model"

	"github.com/apapacy/fitness-release-2/routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5433/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	userStore := model.NewUserStore(db)
	err = userStore.Insert(&model.User{
		ID:       kallax.NewULID(),
		Username: "john",
		Email:    "john@doe.me",
		Password: "1234bunnies",
	})
	if err != nil {
		panic(err)
	}

	routes.GetRouter().Run() // listen and serve on 0.0.0.0:8080
}
