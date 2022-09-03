package http

import (
	"context"
	"net/http"
	"strconv"

	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"
	"lucy/cashier/lib/middleware"

	"github.com/gin-gonic/gin"
)

type ItemCategoryHandler struct {
	ItemCategoryUsecase domain.ItemCategoryUsecaseContract
}

func NewItemCategoryHandler(router *gin.Engine, usecase domain.ItemCategoryUsecaseContract) {
	handler := &ItemCategoryHandler{
		ItemCategoryUsecase: usecase,
	}

	permitted := []middleware.UserRole{
		middleware.UserRole_OWNER,
		middleware.UserRole_BRANCH_OWNER,
	}

	r := router.Group("/", middleware.Auth, middleware.Authorization(permitted))

	r.PUT("/item-categories", handler.UpsertItemCategory)
	r.GET("/item-categories", handler.FindItemCategories)
	r.GET("/item-categories/:id", handler.FindItemCategory)
	r.DELETE("/item-categories/:id", handler.DeleteItemCategory)

	r.POST("items/:itemCategoryId", handler.CreateItem)
	r.PATCH("items/:id", handler.UpdateItem)
	r.GET("items/:id", handler.FindItem)
	r.DELETE("items/:id", handler.DeleteItem)
}

func (h *ItemCategoryHandler) UpsertItemCategory(c *gin.Context) {
	var request domain.ItemCategoryUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.UpsertItemCategory(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item category upsert successfully", result)
}

func (h *ItemCategoryHandler) FindItemCategory(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")
	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemCategory(ctx, branchId, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) FindItemCategories(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemCategories(ctx, branchId, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) DeleteItemCategory(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.DeleteItemCategory(ctx, branchId, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item category Deleted successfully", result)
}

// Item

func (h *ItemCategoryHandler) CreateItem(c *gin.Context) {
	var request domain.ItemCreateRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := c.GetString("branch_uuid")
	itemCategoryId := c.Param("itemCategoryId")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.CreateItem(ctx, branchId, itemCategoryId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, 200, "Item created successfully", result)
}

func (h *ItemCategoryHandler) UpdateItem(c *gin.Context) {
	var request domain.ItemUpdateRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.UpdateItem(ctx, branchId, id, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item updated successfully", result)
}

func (h *ItemCategoryHandler) FindItem(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")
	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItem(ctx, branchId, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) DeleteItem(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.DeleteItem(ctx, branchId, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item deleted successfully", result)
}
