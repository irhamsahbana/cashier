package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lucy/cashier/bootstrap"
	"lucy/cashier/domain"

	_menuCategoryHttp "lucy/cashier/domain/menu_category/delivery/http"
	_menuCategoryRepo "lucy/cashier/domain/menu_category/repository/mongo"
	_menuCategoryUsecase "lucy/cashier/domain/menu_category/usecase"

	_userHttp "lucy/cashier/domain/user/delivery/http"
	_userRepo "lucy/cashier/domain/user/repository/mongo"
	_userUsecase "lucy/cashier/domain/user/usecase"

	_waiterHttp "lucy/cashier/domain/waiter/delivery/http"
	_waiterRepo "lucy/cashier/domain/waiter/repository/mongo"
	_waiterUsecase "lucy/cashier/domain/waiter/usecase"

	_tokenRepo "lucy/cashier/domain/token/repository/mongo"
)

func main() {
	defer func() {
		err := bootstrap.App.Maria.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if !bootstrap.App.Config.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))

	tokenRepo := _tokenRepo.NewTokenMongoRepository(*mongoDatabase, domain.UserTokenableType)
	userRepo := _userRepo.NewUserMongoRepository(*mongoDatabase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, tokenRepo, timeoutContext)
	_userHttp.NewUserHandler(router, userUsecase)

	menuCategoryRepo := _menuCategoryRepo.NewMenuCategoryMongoRepository(*mongoDatabase)
	menuCategoryUsecase := _menuCategoryUsecase.NewMenuCategoryUsecase(menuCategoryRepo, timeoutContext)
	_menuCategoryHttp.NewMenuCategoryHandler(router, menuCategoryUsecase)

	waiterRepo := _waiterRepo.NewWaiterMongoRepository(*mongoDatabase)
	waiterUsecase := _waiterUsecase.NewWaiterUsecase(waiterRepo, timeoutContext)
	_waiterHttp.NewWaiterHandler(router, waiterUsecase)

	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(router.Run(appPort))
}
