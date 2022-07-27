package http

import (
	"context"
	"net/http"

	"lucy/cashier/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Message string `json:"message"`
}

type MenuCategoryHandler struct {
	MenuCategoryUsecase domain.MenuCategoryUsecaseContract
}

func NewMenuCategoryHandler(router *gin.Engine, usecase domain.MenuCategoryUsecaseContract) {
	handler := &MenuCategoryHandler{
		MenuCategoryUsecase: usecase,
	}

	router.POST("/menu-categories", handler.CreateMenuCategory)
	router.GET("/menu-categories/:id", handler.FindMenuCategory)
}

func (handler *MenuCategoryHandler) CreateMenuCategory(c *gin.Context) {
	var request domain.MenuCategory
	var	err error

	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&request); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	result, err, httpCode := handler.MenuCategoryUsecase.CreateMenuCategory(ctx, &request)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpCode, result)
}

func (handler *MenuCategoryHandler) FindMenuCategory(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, err, httpCode := handler.MenuCategoryUsecase.FindMenuCategory(ctx, id)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpCode, result)
}

func isRequestValid(request *domain.MenuCategory) (bool, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return false, err
	}
	return true, nil
}