package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeShiftHandler struct {
	EmployeeShiftUsecase domain.EmployeeShiftUsecaseContract
}

func NewEmployeeShiftHandler(router *gin.Engine, usecase domain.EmployeeShiftUsecaseContract) {
	handler := &EmployeeShiftHandler{
		EmployeeShiftUsecase: usecase,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
		middleware.UserRole_ADMIN_CASHIER,
		middleware.UserRole_CASHIER,
	}

	r := router.Group("/", middleware.Auth, middleware.Authorization(permitted))

	r.POST("/employee-shifts/clock-in", handler.ClockIn)
	r.PATCH("/employee-shifts/clock-out", handler.ClockOut)
	r.GET("/employee-shifts/history", handler.History)
	r.GET("/employee-shifts/active", handler.Active)
}

func (h *EmployeeShiftHandler) ClockIn(c *gin.Context) {
	var request domain.EmployeeShiftClockInRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.ClockIn(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, http.StatusOK, "success to clock in", result)
}

func (h *EmployeeShiftHandler) ClockOut(c *gin.Context) {
	var request domain.EmployeeShiftClockOutRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.ClockOut(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, http.StatusOK, "success to clock out", result)
}

func (h *EmployeeShiftHandler) History(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "0")

	// convert limit and page to int
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusBadRequest, "limit must be integer", nil)
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusBadRequest, "page must be integer", nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.History(ctx, branchId, limitInt, pageInt)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "success to get employee shift history", result)
}

func (h *EmployeeShiftHandler) Active(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.Active(ctx, branchId)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "success to get active employee shift", result)
}
