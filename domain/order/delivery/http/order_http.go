package http

import (
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecaseContract
}

func NewOrderHandler(router *gin.Engine, usecase domain.OrderUsecaseContract) {
	handler := &orderHandler{
		orderUsecase: usecase,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
		middleware.UserRole_ADMIN_CASHIER,
		middleware.UserRole_CASHIER,
	}

	r := router.Group("/", middleware.Auth, middleware.Authorization(permitted))

	r.PUT("/order-groups", handler.UpsertActiveOrder)
}

func (h *orderHandler) UpsertActiveOrder(ctx *gin.Context) {
	var request dto.OrderGroupUpsertRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := ctx.GetString("branch_uuid")

	result, httpCode, err := h.orderUsecase.UpsertActiveOrder(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(ctx, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(ctx, http.StatusOK, "success to upsert active order", result)
}
