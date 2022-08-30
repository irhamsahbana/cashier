package http

import (
	"context"
	"net/http"
	"strconv"

	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"

	"github.com/gin-gonic/gin"
)

type ItemCategoryHandler struct {
	ItemCategoryUsecase domain.ItemCategoryUsecaseContract
}

func NewItemCategoryHandler(router *gin.Engine, usecase domain.ItemCategoryUsecaseContract) {
	handler := &ItemCategoryHandler{
		ItemCategoryUsecase: usecase,
	}

	router.PUT("/item-categories", handler.UpsertItemCategory)
	router.GET("/item-categories", handler.FindItemCategories)
	router.GET("/item-categories/:id", handler.FindItemCategory)
	router.DELETE("/item-categories/:id", handler.DeleteItemCategory)

	router.POST("items/:itemCategoryId", handler.CreateItem)
	router.PATCH("items/:id", handler.UpdateItem)
	router.GET("items/:id", handler.FindItem)
	router.DELETE("items/:id", handler.DeleteItem)
}

func (h *ItemCategoryHandler) UpsertItemCategory(c *gin.Context) {
	var request domain.ItemCategoryUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.UpsertItemCategory(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item category upsert successfully", result)
}

func (h *ItemCategoryHandler) FindItemCategory(c *gin.Context) {
	id := c.Param("id")
	trashed := c.Query("with_trashed")

	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemCategory(ctx, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) FindItemCategories(c *gin.Context) {
	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemCategories(ctx, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) DeleteItemCategory(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.DeleteItemCategory(ctx, id)
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

	itemCategoryId := c.Param("itemCategoryId")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.CreateItem(ctx, itemCategoryId, &request)
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

	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.UpdateItem(ctx, id, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item updated successfully", result)
}

func (h *ItemCategoryHandler) FindItem(c *gin.Context) {
	id := c.Param("id")
	trashed := c.Query("with_trashed")

	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItem(ctx, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *ItemCategoryHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.DeleteItem(ctx, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item deleted successfully", result)
}
