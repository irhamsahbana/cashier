package http

import (
	"lucy/cashier/domain"

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

}

func (h *WaiterHandler) FindWaiter(c *gin.Context) {

}

func (h *WaiterHandler) DeleteWaiter(c *gin.Context) {

}