package http

import (
	"context"
	"net/http"

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
	r.DELETE("/item-categories/:id", handler.DeleteItemCategory)

	r.PUT("items/:itemCategoryId", handler.UpsertItemAndVariants)
	r.GET("items/:id", handler.FindItemAndVariants)
	r.DELETE("items/:id", handler.DeleteItemAndVariants)
	/*
		r.GET("/item-categories/:id", handler.FindItemCategory)

		r.PATCH("items/:id", handler.UpdateItem)

	*/
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
	result, httpCode, err := h.ItemCategoryUsecase.UpsertItemCategoryAndModifiers(ctx, branchId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item category upsert successfully", result)
}

/*
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
*/

func (h *ItemCategoryHandler) FindItemCategories(c *gin.Context) {
	branchId := c.GetString("branch_uuid")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemCategories(ctx, branchId)
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

// Item and its variants

func (h *ItemCategoryHandler) UpsertItemAndVariants(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	itemCategoryId := c.Param("itemCategoryId")

	var request domain.ItemAndVariantsUpsertRequest

	err := c.BindJSON(&request)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.UpsertItemAndVariants(ctx, branchId, itemCategoryId, &request)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item and variants upsert successfully", result)
}

func (h *ItemCategoryHandler) FindItemAndVariants(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.FindItemAndVariants(ctx, branchId, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item and variants upsert successfully", result)
}

func (h *ItemCategoryHandler) DeleteItemAndVariants(c *gin.Context) {
	branchId := c.GetString("branch_uuid")
	id := c.Param("id")

	ctx := context.Background()
	result, httpCode, err := h.ItemCategoryUsecase.DeleteItemAndVariants(ctx, branchId, id)
	if err != nil {
		http_response.ReturnResponse(c, httpCode, err.Error(), nil)
		return
	}

	http_response.ReturnResponse(c, httpCode, "Item and variants deleted successfully", result)
}
