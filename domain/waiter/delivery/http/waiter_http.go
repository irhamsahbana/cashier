package http

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WaiterHandler struct {
	WaiterUsecase	domain.WaiterUsecaseContract
}

func NewWaiterHandler(router *gin.Engine, usecase domain.WaiterUsecaseContract) {
	handler := &WaiterHandler{
		WaiterUsecase: usecase,
	}

	router.PUT("waiters", handler.UpsertWaiter)
	router.GET("waiters/:id", handler.FindWaiter)
	router.DELETE("waiters/:id", handler.DeleteWaiter)
}

func (h *WaiterHandler) UpsertWaiter(c *gin.Context) {
	var request domain.WaiterUpsertrequest

	if err := c.BindJSON(&request); err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpcode, err := h.WaiterUsecase.UpsertWaiter(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpcode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpcode, "Waiter upsert successfully", result)
}

func (h *WaiterHandler) FindWaiter(c *gin.Context) {
	id := c.Param("id")
	trashed := c.Query("with_trashed")

	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.WaiterUsecase.FindWaiter(ctx, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *WaiterHandler) DeleteWaiter(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.WaiterUsecase.DeleteWaiter(ctx, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Waiter deleted successfully", result)
}