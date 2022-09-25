package route

import (
	"fmt"
	"lucy/cashier/bootstrap"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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

	_zoneHttp "lucy/cashier/domain/zone/delivery/http"
	_zoneRepo "lucy/cashier/domain/zone/repository/mongo"
	_zoneUsecase "lucy/cashier/domain/zone/usecase"

	_waiterHttp "lucy/cashier/domain/waiter/delivery/http"
	_waiterRepo "lucy/cashier/domain/waiter/repository/mongo"
	_waiterUsecase "lucy/cashier/domain/waiter/usecase"

	_shiftHttp "lucy/cashier/domain/employee_shift/delivery/http"
	_shiftRepo "lucy/cashier/domain/employee_shift/repository/mongo"
	_shiftUsecase "lucy/cashier/domain/employee_shift/usecase"

	_fileHttp "lucy/cashier/domain/file/delivery/http"
	_fileRepo "lucy/cashier/domain/file/repository/mongo"
	_fileUsecase "lucy/cashier/domain/file/usecase"

	_branchDiscountRepo "lucy/cashier/domain/branch_discount/repository/mongo"
	_companyRepo "lucy/cashier/domain/company/repository/mongo"
	_tokenRepo "lucy/cashier/domain/token/repository/mongo"
)

func NewHttpRoutes(r *gin.Engine) {
	if !bootstrap.App.Config.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	appStorageURL := bootstrap.App.Config.GetString("app.url") + bootstrap.App.Config.GetString("app.static_asssets_url")

	r.Static(bootstrap.App.Config.GetString("app.static_asssets_url"), bootstrap.App.Config.GetString("app.static_assets"))
	r.Use(cors.Default())

	// user
	tokenRepo := _tokenRepo.NewTokenMongoRepository(*mongoDatabase, domain.TokenableType_USER)
	userRepo := _userRepo.NewUserMongoRepository(*mongoDatabase)
	userRoleRepo := _userRoleRepo.NewUserRoleMongoRepository(*mongoDatabase)
	companyRepo := _companyRepo.NewCompanyMongoRepository(*mongoDatabase)
	branchDiscountRepo := _branchDiscountRepo.NewBranchDiscountMongoRepository(*mongoDatabase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, userRoleRepo, tokenRepo, companyRepo, branchDiscountRepo, timeoutContext)
	_userHttp.NewUserHandler(r, userUsecase)

	// user role
	userRoleUsecase := _userRoleUsecase.NewUserRoleUsecase(userRoleRepo, timeoutContext)
	_userRoleHttp.NewUserRoleHandler(r, userRoleUsecase)

	// item category
	itemCategoryRepo := _itemCategoryRepo.NewItemCategoryMongoRepository(*mongoDatabase)
	itemCategoryUsecase := _itemCategoryUsecase.NewItemCategoryUsecase(itemCategoryRepo, timeoutContext)
	_itemCategoryHttp.NewItemCategoryHandler(r, itemCategoryUsecase)

	// waiter
	waiterRepo := _waiterRepo.NewWaiterMongoRepository(*mongoDatabase)
	waiterUsecase := _waiterUsecase.NewWaiterUsecase(waiterRepo, timeoutContext)
	_waiterHttp.NewWaiterHandler(r, waiterUsecase)

	// employee shift
	shiftRepo := _shiftRepo.NewEmployeeShiftMongoRepository(*mongoDatabase)
	shiftUsecase := _shiftUsecase.NewEmployeeShiftUsecase(shiftRepo, timeoutContext)
	_shiftHttp.NewEmployeeShiftHandler(r, shiftUsecase)

	// space group
	spaceGroupRepo := _spaceGroupRepo.NewSpaceGroupMongoRepository(*mongoDatabase)
	spaceGroupUsecase := _spaceGroupUsecase.NewSpaceGroupUsecase(spaceGroupRepo, timeoutContext)
	_spaceGroupHttp.NewSpaceGroupHandler(r, spaceGroupUsecase)

	// zone
	zoneRepo := _zoneRepo.NewZoneMongoRepository(*mongoDatabase)
	zoneUsecase := _zoneUsecase.NewZoneUsecase(zoneRepo, timeoutContext)
	_zoneHttp.NewZoneHandler(r, zoneUsecase)

	// file
	fileRepo := _fileRepo.NewFileMongoRepository(*mongoDatabase)
	fileUsecase := _fileUsecase.NewFileUploadUsecase(fileRepo, timeoutContext)
	_fileHttp.NewFileHandler(r, fileUsecase, appStorageURL)

	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))

	logger.Log(logrus.Fields{
		"port": appPort,
	}).Error(r.Run(appPort))
}
