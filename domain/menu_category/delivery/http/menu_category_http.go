package http

import (
	"context"
	"net/http"
	"strconv"

	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"

	"github.com/gin-gonic/gin"
)

type MenuCategoryHandler struct {
	MenuCategoryUsecase domain.MenuCategoryUsecaseContract
}

func NewMenuCategoryHandler(router *gin.Engine, usecase domain.MenuCategoryUsecaseContract) {
	handler := &MenuCategoryHandler{
		MenuCategoryUsecase: usecase,
	}

	router.PUT("/menu-categories", handler.UpsertMenuCategory)
	router.GET("/menu-categories", handler.FindMenuCategories)
	router.GET("/menu-categories/:id", handler.FindMenuCategory)
	router.DELETE("/menu-categories/:id", handler.DeleteMenuCategory)

	router.POST("menus/:menuCategoryId", handler.CreateMenu)
	router.PATCH("menus/:id", handler.UpdateMenu)
	router.GET("menus/:id", handler.FindMenu)
	router.DELETE("menus/:id", handler.DeleteMenu)
}

func (h *MenuCategoryHandler) UpsertMenuCategory(c *gin.Context) {
	var request domain.MenuCategoryUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.UpsertMenuCategory(ctx, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, 200, "Menu category upsert successfully", result)
}

func (h *MenuCategoryHandler) FindMenuCategory(c *gin.Context) {
	id := c.Param("id")
	trashed := c.Query("with_trashed")

	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.FindMenuCategory(ctx, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *MenuCategoryHandler) FindMenuCategories(c *gin.Context) {
	trashed := c.Query("with_trashed")
	withTrashed, _ := strconv.ParseBool(trashed)

	ctx:= context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.FindMenuCategories(ctx, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *MenuCategoryHandler) DeleteMenuCategory(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.DeleteMenuCategory(ctx, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Menu category Deleted successfully", result)
}

// Menu

func (h *MenuCategoryHandler) CreateMenu(c *gin.Context) {
	var request domain.MenuCreateRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	menuCategoryId := c.Param("menuCategoryId")

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.CreateMenu(ctx, menuCategoryId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, 200, "Menu created successfully", result)
}

func (h *MenuCategoryHandler) UpdateMenu(c *gin.Context) {
	var request domain.MenuUpdateRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.UpdateMenu(ctx, id, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Menu updated successfully", result)
}

func (h *MenuCategoryHandler) FindMenu( c *gin.Context) {
	id := c.Param("id")
	trashed := c.Query("with_trashed")

	withTrashed, _ := strconv.ParseBool(trashed)

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.FindMenu(ctx, id, withTrashed)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "OK", result)
}

func (h *MenuCategoryHandler) DeleteMenu(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.MenuCategoryUsecase.DeleteMenu(ctx, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Menu deleted successfully", result)
}