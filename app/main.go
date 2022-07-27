package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lucy/cashier/bootstrap"
	_menuCategoryHttp "lucy/cashier/menu_category/delivery/http"
	_menuCategoryRepo "lucy/cashier/menu_category/repository/mongo"
	_menuCategoryUsecase "lucy/cashier/menu_category/usecase"
)

func main() {
	defer func() {
		err := bootstrap.App.Maria.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	router.Use(cors.Default())

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))

	menuCategoryRepo := _menuCategoryRepo.NewMenuCategoryMongoRepository(*mongoDatabase)
	menuCategoryUsecase := _menuCategoryUsecase.NewMenuCategoryUsecase(menuCategoryRepo, timeoutContext)
	_menuCategoryHttp.NewMenuCategoryHandler(router, menuCategoryUsecase)


	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(router.Run(appPort))
}