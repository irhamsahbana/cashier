package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"
	"strconv"
	"time"

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

	r.POST("/invoices", handler.InsertInvoice)
	r.GET("/invoices", handler.FindInvoiceHistories)
	r.POST("/invoice-refunds", handler.MakeRefund)

}

func (h *orderHandler) UpsertActiveOrder(c *gin.Context) {
	var request dto.OrderGroupUpsertRequest
	branchId := c.GetString("branch_uuid")

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := validateUpserOrderRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.UpsertActiveOrder(ctx, branchId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to upsert active order", result)
}

func (h *orderHandler) FindActiveOrders(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.FindActiveOrders(ctx, branchId)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to find active orders", result)
}

func (h *orderHandler) DeleteActiveOrder(c *gin.Context) {
	var request dto.OrderGroupDeleteRequest
	branchId := c.GetString("branch_uuid")
	orderId := c.Param("id")

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.DeleteActiveOrder(ctx, branchId, orderId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to delete active order", result)
}

func (h *orderHandler) InsertInvoice(c *gin.Context) {
	var request dto.InvoiceInsertRequest
	branchId := c.GetString("branch_uuid")

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := ValidateInsertInvoiceRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.InsertInvoice(ctx, branchId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to insert invoice", result)
}

func (h *orderHandler) FindInvoiceHistories(c *gin.Context) {
	// cursor pagination
	limitQ := c.DefaultQuery("limit", "15")
	cursor := c.DefaultQuery("cursor", "")
	direction := c.DefaultQuery("direction", "next")
	sortType := c.DefaultQuery("sort_type", "desc")

	// filter
	branchId := c.GetString("branch_uuid")
	from := c.DefaultQuery("from", "")
	to := c.DefaultQuery("to", "")

	if direction != "prev" {
		direction = "next"
	}

	// direction make only next, not prev
	direction = "next"

	if sortType != "asc" {
		sortType = "desc"
	}

	// limit
	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, "limit must be integer", nil)
		return
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", limit)
	ctx = context.WithValue(ctx, "cursor", cursor)
	ctx = context.WithValue(ctx, "direction", direction)
	ctx = context.WithValue(ctx, "sort_type", sortType)

	ctx = context.WithValue(ctx, "branch_uuid", branchId)
	ctx = context.WithValue(ctx, "from", from)
	ctx = context.WithValue(ctx, "to", to)

	result, nextCur, prevCur, httpCode, err := h.orderUsecase.FindInvoiceHistories(ctx)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}
	var nextCurTime, prevCurTime *time.Time

	if nextCur != nil {
		nextCurInt64 := *nextCur
		nextCurTimeSource := time.UnixMicro(nextCurInt64)
		nextCurTime = &nextCurTimeSource
	}

	if prevCur != nil {
		prevCurInt64 := *prevCur
		prevCurTimeSource := time.UnixMicro(prevCurInt64)
		prevCurTime = &prevCurTimeSource
	}

	pagination := struct {
		NextCursor any `json:"next_cursor"`
		PrevCursor any `json:"prev_cursor"`
	}{
		NextCursor: nextCurTime,
		PrevCursor: prevCurTime,
	}

	meta := struct {
		Pagination any `json:"pagination"`
	}{
		Pagination: pagination,
	}

	http_response.JSON(c, http.StatusOK, "success to find invoice histories", result, meta)
}

func (h *orderHandler) MakeRefund(c *gin.Context) {
	var request dto.RefundInsertRequest
	branchId := c.GetString("branch_uuid")

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := ValidateInsertRefundRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.orderUsecase.InsertRefund(ctx, branchId, request.InvoiceUUID, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to make refund", result)
}
