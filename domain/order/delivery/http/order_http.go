package http

import (
	"context"
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
	r.GET("/order-groups", handler.FindActiveOrders)
	r.DELETE("/order-groups/:id", handler.DeleteActiveOrder)
}

func (h *orderHandler) UpsertActiveOrder(c *gin.Context) {
	var request dto.OrderGroupUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	errMsg := validateUpserOrderRequest(&request)
	if len(errMsg) > 0 {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.UpsertActiveOrder(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, http.StatusOK, "success to upsert active order", result)
}

func (h *orderHandler) FindActiveOrders(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.FindActiveOrders(ctx, branchId)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, http.StatusOK, "success to find active orders", result)
}

func (h *orderHandler) DeleteActiveOrder(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	orderId := c.Param("id")

	var request dto.OrderGroupDeleteRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.DeleteActiveOrder(ctx, branchId, orderId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, http.StatusOK, "success to delete active order", result)
}
