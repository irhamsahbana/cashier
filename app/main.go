package main

import (
	"log"

	"lucy/cashier/bootstrap"
	route "lucy/cashier/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		err := bootstrap.App.Maria.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	route.NewHttpRoutes(gin.Default())
}
