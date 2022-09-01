package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lucy/cashier/bootstrap"
	"lucy/cashier/domain"

	_itemCategoryHttp "lucy/cashier/domain/item_category/delivery/http"
	_itemCategoryRepo "lucy/cashier/domain/item_category/repository/mongo"
	_itemCategoryUsecase "lucy/cashier/domain/item_category/usecase"

	_userRoleHttp "lucy/cashier/domain/user_role/delivery/http"
	_userRoleRepo "lucy/cashier/domain/user_role/repository/mongo"
	_userRoleUsecase "lucy/cashier/domain/user_role/usecase"

	_userHttp "lucy/cashier/domain/user/delivery/http"
	_userRepo "lucy/cashier/domain/user/repository/mongo"
	_userUsecase "lucy/cashier/domain/user/usecase"

	_spaceGroupHttp "lucy/cashier/domain/space_group/delivery/http"
	_spaceGroupRepo "lucy/cashier/domain/space_group/repository/mongo"
	_spaceGroupUsecase "lucy/cashier/domain/space_group/usecase"

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

	tokenRepo := _tokenRepo.NewTokenMongoRepository(*mongoDatabase, domain.TokenableType_USER)
	userRepo := _userRepo.NewUserMongoRepository(*mongoDatabase)
	userRoleRepo := _userRoleRepo.NewUserRoleMongoRepository(*mongoDatabase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, userRoleRepo, tokenRepo, timeoutContext)
	_userHttp.NewUserHandler(router, userUsecase)

	userRoleUsecase := _userRoleUsecase.NewUserRoleUsecase(userRoleRepo, timeoutContext)
	_userRoleHttp.NewUserRoleHandler(router, userRoleUsecase)

	itemCategoryRepo := _itemCategoryRepo.NewItemCategoryMongoRepository(*mongoDatabase)
	itemCategoryUsecase := _itemCategoryUsecase.NewItemCategoryUsecase(itemCategoryRepo, timeoutContext)
	_itemCategoryHttp.NewItemCategoryHandler(router, itemCategoryUsecase)

	waiterRepo := _waiterRepo.NewWaiterMongoRepository(*mongoDatabase)
	waiterUsecase := _waiterUsecase.NewWaiterUsecase(waiterRepo, timeoutContext)
	_waiterHttp.NewWaiterHandler(router, waiterUsecase)

	spaceGroupRepo := _spaceGroupRepo.NewSpaceGroupMongoRepository(*mongoDatabase)
	spaceGroupUsecase := _spaceGroupUsecase.NewSpaceGroupUsecase(spaceGroupRepo, timeoutContext)
	_spaceGroupHttp.NewSpaceGroupHandler(router, spaceGroupUsecase)

	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(router.Run(appPort))
}
