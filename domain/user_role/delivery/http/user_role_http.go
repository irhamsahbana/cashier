package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoleHandler struct {
	userRoleUsecase domain.UserRoleUsecaseContract
}

func NewUserRoleHandler(router *gin.Engine, usecase domain.UserRoleUsecaseContract) {
	handler := &UserRoleHandler{
		userRoleUsecase: usecase,
	}

	authorized := router.Group("/", middleware.Auth)
	authorized.PUT("/user-roles", handler.UpsertUserRole)
}

func (h *UserRoleHandler) UpsertUserRole(c *gin.Context) {
	var request domain.UserRoleUpsertrequest
	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.userRoleUsecase.UpsertUserRole(ctx, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		c.Abort()
		return
	}

	http_response.JSON(c, httpCode, "User Role Created", result)
}
