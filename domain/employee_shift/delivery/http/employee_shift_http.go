package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
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

	r.POST("/employee-shifts/entry-cash/:employee_shift_uuid", handler.InsertEntryCash)
}

func (h *EmployeeShiftHandler) ClockIn(c *gin.Context) {
	var request dto.EmployeeShiftClockInRequest
	branchId := c.GetString("branch_uuid")

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := ValidateClockInRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.ClockIn(ctx, branchId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to clock in", result)
}

func (h *EmployeeShiftHandler) ClockOut(c *gin.Context) {
	var request dto.EmployeeShiftClockOutRequest
	branchId := c.GetString("branch_uuid")

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := validateClockOutRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.ClockOut(ctx, branchId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to clock out", result)
}

func (h *EmployeeShiftHandler) History(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "0")

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

	ctx := context.Background()
	result, httpCode, err := h.EmployeeShiftUsecase.History(ctx, branchId, limitInt, pageInt)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "success to get employee shift history", result)
}

func (h *EmployeeShiftHandler) Active(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, code, err := h.EmployeeShiftUsecase.Active(ctx, branchId)
	if err != nil {
		http_response.JSON(c, code, err.Error(), nil)
		return
	}

	http_response.JSON(c, code, "success to get active employee shift", result)
}

func (h *EmployeeShiftHandler) InsertEntryCash(c *gin.Context) {
	var request dto.CashEntryInsertRequest
	branchId := c.GetString("branch_uuid")
	shiftId := c.Param("employee_shift_uuid")

	if err := c.BindJSON(&request); err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	if errMsg := ValidateCashEntryInsertRequest(&request); len(errMsg) > 0 {
		http_response.JSON(c, http.StatusUnprocessableEntity, errMsg, nil)
		return
	}

	ctx := context.Background()
	result, code, err := h.EmployeeShiftUsecase.InsertEntryCash(ctx, branchId, shiftId, &request)
	if err != nil {
		http_response.JSON(c, code, err.Error(), nil)
		return
	}

	http_response.JSON(c, http.StatusOK, "success to insert entry cash", result)
}
