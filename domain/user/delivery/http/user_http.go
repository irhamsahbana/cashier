package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/http_response"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"lucy/cashier/lib/middleware"
	"net/http"

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

	authorized := router.Group("/", middleware.Auth)
	authorized.GET("/user-infos", handler.Profile)
	authorized.GET("/auth/logout", handler.Logout)

	router.POST("auth/login", handler.Login)
	router.GET("auth/refresh-token", handler.RefreshToken)
}

func (h *UserHandler) Login(c *gin.Context) {
	var request dto.UserLoginRequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.Login(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Authenticated", result)
}

func (h *UserHandler) Profile(c *gin.Context) {
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	result, httpcode, err := h.UserUsecase.UserBranchInfo(ctx, userId, true)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http_response.ReturnResponse(c, http.StatusNotFound, err.Error(), nil)
			return
		}

		http_response.ReturnResponse(c, httpcode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpcode, "Profile", result)
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	accessToken := c.GetHeader("X-ACCESS-TOKEN")
	refreshToken := c.GetHeader("X-REFRESH-TOKEN")

	_, err := jwthandler.ValidateToken(accessToken)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
		} else {
			http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}
	}

	claimsRT, err := jwthandler.ValidateToken(refreshToken)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.UserUsecase.RefreshToken(ctx, accessToken, refreshToken, claimsRT.UserUUID)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Token refreshed", result)
}

func (h *UserHandler) Logout(c *gin.Context) {
	AT := c.GetString("access_token")
	userId := c.GetString("user_uuid")

	ctx := context.Background()
	_, httpcode, err := h.UserUsecase.Logout(ctx, userId, AT)
	if err != nil {
		http_response.ReturnResponse(c, httpcode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpcode, "Logout", nil)
}
