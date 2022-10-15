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

type SpaceGroupHandler struct {
	SpaceGroupUsecase domain.SpaceGroupUsecaseContract
}

func NewSpaceGroupHandler(router *gin.Engine, usecase domain.SpaceGroupUsecaseContract) {
	handler := &SpaceGroupHandler{
		SpaceGroupUsecase: usecase,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
	}

	r := router.Group("/", middleware.Auth, middleware.Authorization(permitted))

	r.GET("/space-groups", handler.FindSpaceGroups)
	r.PUT("/space-groups", handler.UpsertSpaceGroup)
	r.DELETE("/space-groups/:id", handler.DeleteSpaceGroup)
	r.GET("/space-groups/:id", handler.FindSpaceGroup)

	r.POST("/spaces/:spaceGroupId", handler.CreateSpace)
	r.GET("/spaces/:id", handler.FindSpace)
	r.DELETE("/spaces/:id", handler.DeleteSpace)
	r.PATCH("/spaces/:id", handler.UpdateSpace)
}

func (h *SpaceGroupHandler) UpsertSpaceGroup(c *gin.Context) {
	var request domain.SpaceGroupUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.SpaceGroupUsecase.UpsertSpaceGroup(ctx, c.GetString("branch_uuid"), &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space group upsert successfully", result)
}

func (h *SpaceGroupHandler) FindSpaceGroup(c *gin.Context) {
	ctx := context.Background()
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	result, httpCode, err := h.SpaceGroupUsecase.FindSpaceGroup(ctx, branchId, id, withTrashed)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space group found", result)
}

func (h *SpaceGroupHandler) DeleteSpaceGroup(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	result, httpCode, err := h.SpaceGroupUsecase.DeleteSpaceGroup(ctx, c.GetString("branch_uuid"), id)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space group delete successfully", result)
}

func (h *SpaceGroupHandler) FindSpaceGroups(c *gin.Context) {
	ctx := context.Background()
	branchId := c.GetString("branch_uuid")

	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	result, httpCode, err := h.SpaceGroupUsecase.FindSpaceGroups(ctx, branchId, withTrashed)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space groups found", result)
}

// space

func (h *SpaceGroupHandler) CreateSpace(c *gin.Context) {
	var request domain.SpaceCreateRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	spaceGroupId := c.Param("spaceGroupId")
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.SpaceGroupUsecase.CreateSpace(ctx, branchId, spaceGroupId, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space created successfully", result)
}

func (h *SpaceGroupHandler) FindSpace(c *gin.Context) {
	ctx := context.Background()
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	result, httpCode, err := h.SpaceGroupUsecase.FindSpace(ctx, branchId, id, withTrashed)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space found", result)
}

func (h *SpaceGroupHandler) DeleteSpace(c *gin.Context) {
	ctx := context.Background()
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	result, httpCode, err := h.SpaceGroupUsecase.DeleteSpace(ctx, branchId, id)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space delete successfully", result)
}

func (h *SpaceGroupHandler) UpdateSpace(c *gin.Context) {
	var request domain.SpaceUpdateRequest
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	err := c.BindJSON(&request)
	if err != nil {
		http_response.JSON(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.SpaceGroupUsecase.UpdateSpace(ctx, branchId, id, &request)
	if err != nil {
		http_response.JSON(c, httpCode, err.Error(), nil)
		return
	}

	http_response.JSON(c, httpCode, "Space update successfully", result)
}
