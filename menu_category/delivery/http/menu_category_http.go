package http

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/http_response"

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
	router.DELETE("/menu-categories/:id", handler.DeleteMenuCategory)
	router.PATCH("/menu-categories/:id", handler.UpdateMenuCategory)
}

func (handler *MenuCategoryHandler) CreateMenuCategory(c *gin.Context) {
	var request domain.MenuCategory

	err := c.BindJSON(&request)
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
	result, httpCode, err := handler.MenuCategoryUsecase.CreateMenuCategory(ctx, &request)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	res := http_response.Response{
		StatusCode: httpCode,
		Message: "OK",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data: result,
	}

	c.JSON(httpCode, res)
}

func (handler *MenuCategoryHandler) FindMenuCategory(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := handler.MenuCategoryUsecase.FindMenuCategory(ctx, id)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpCode, result)
}

func (handler *MenuCategoryHandler) DeleteMenuCategory(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := handler.MenuCategoryUsecase.DeleteMenuCategory(ctx, id)
	if err != nil {
		c.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpCode, result)
}

func (handler *MenuCategoryHandler) UpdateMenuCategory(c *gin.Context) {
	var request domain.MenuCategory

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&request); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := handler.MenuCategoryUsecase.UpdateMenuCategory(ctx, id, &request)
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