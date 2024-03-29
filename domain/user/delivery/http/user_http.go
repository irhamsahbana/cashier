package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/http_response"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"lucy/cashier/lib/middleware"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecaseContract
}

func NewUserHandler(router *gin.Engine, usecase domain.UserUsecaseContract) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	var permittedRoles = []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
		middleware.UserRole_MANAGER,
		middleware.UserRole_ADMIN_CASHIER,
		middleware.UserRole_CASHIER,
	}

	authorized := router.Group("/", middleware.Auth)
	permitted := router.Group("/", middleware.Auth, middleware.Authorization(permittedRoles))
	permitted.PUT("/customers", handler.UpsertCustomer)
	permitted.GET("/customers", handler.FindCustomers)
	authorized.GET("/user-infos", handler.Profile)
	authorized.GET("/auth/logout", handler.Logout)

	router.POST("auth/login", handler.Login)
	router.GET("auth/refresh-token", handler.RefreshToken)
}

func (h *UserHandler) Login(c *gin.Context) {
	var request dto.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.Login(ctx, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Authenticated", result)
}

func (h *UserHandler) Profile(c *gin.Context) {
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	result, httpcode, err := h.UserUsecase.UserBranchInfo(ctx, userId, true)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http_response.JSON(c, http.StatusNotFound, err.Error(), nil)
			return
		}

		http_response.JSON(c, httpcode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpcode, "Profile", result)
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	accessToken := c.GetHeader("X-ACCESS-TOKEN")
	refreshToken := c.GetHeader("X-REFRESH-TOKEN")

	_, err := jwthandler.ValidateToken(accessToken)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
		} else {
			http_response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}
	}

	claimsRT, err := jwthandler.ValidateToken(refreshToken)
	if err != nil {
		http_response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RefreshToken(ctx, accessToken, refreshToken, claimsRT.UserUUID)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Token refreshed", result)
}

func (h *UserHandler) Logout(c *gin.Context) {
	AT := c.GetString("access_token")
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	_, httpcode, err := h.UserUsecase.Logout(ctx, userId, AT)
	if err != nil {
		http_response.JSON(c, httpcode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpcode, "Logout", nil)
}

func (h *UserHandler) UpsertCustomer(c *gin.Context) {
	var request dto.CustomerUpserRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	errMsg := validateUpsertCustomerRequest(&request)
	if len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpcode, err := h.UserUsecase.UpsertCustomer(ctx, branchId, &request)
	if err != nil {
		http_response.JSON(c, httpcode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpcode, "Customer upserted", result)
}

func (h *UserHandler) FindCustomers(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "0")
	withTrashed := c.DefaultQuery("with_trashed", "false")

	// convert limit and page to int
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http_response.JSON(c, http.StatusBadRequest, "limit must be integer", nil)
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http_response.JSON(c, http.StatusBadRequest, "page must be integer", nil)
		return
	}

	// withTrashed boolean
	withTrashedBool, _ := strconv.ParseBool(withTrashed)

	ctx := context.Background()
	result, httpcode, err := h.UserUsecase.FindCustomers(ctx, branchId, limitInt, pageInt, withTrashedBool)
	if err != nil {
		http_response.JSON(c, httpcode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpcode, "Customers found", result)
}
