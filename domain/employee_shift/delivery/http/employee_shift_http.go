package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"

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
