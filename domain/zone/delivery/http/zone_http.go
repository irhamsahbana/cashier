package http

import (
	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ZoneHandler struct {
	ZoneUsecase domain.ZoneUsecaseContract
}

func NewZoneHandler(router *gin.Engine, usecase domain.ZoneUsecaseContract) {
	handler := &ZoneHandler{
		ZoneUsecase: usecase,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
	}

	r := router.Group("/", middleware.Auth, middleware.Authorization(permitted))

	r.PUT("/zones", handler.UpsertZones)
	r.GET("/zones", handler.GetZones)
}

func (handler *ZoneHandler) UpsertZones(ctx *gin.Context) {
	var request domain.ZoneUpsertRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := ctx.GetString("branch_uuid")

	result, httpCode, err := handler.ZoneUsecase.UpsertZones(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(ctx, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(ctx, http.StatusOK, "success to upsert zones", result)
}

func (handler *ZoneHandler) GetZones(ctx *gin.Context) {
	branchId := ctx.GetString("branch_uuid")

	result, httpCode, err := handler.ZoneUsecase.Zones(ctx, branchId)
	if err != nil {
		http_response.ReturnResponse(ctx, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(ctx, http.StatusOK, "success to get zones", result)
}
