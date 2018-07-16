package main

import (
	"github.com/apapacy/fitness-release-2/routes"
)

func main() {
	routes.GetRouter().Run() // listen and serve on 0.0.0.0:8080
}
