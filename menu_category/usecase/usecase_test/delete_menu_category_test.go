package usecase_test

import (
	"lucy/cashier/domain"
	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteMenuCategory_normalCase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)
	// prepare for mocking
	deletedAt := time.Now().UTC().UnixMicro()

	menuCategory := domain.MenuCategory{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: time.Now().UTC().UnixMicro(),
		DeletedAt: &deletedAt,
	}

	mockMenuCategoryRepository.Mock.On("DeleteMenuCategory", ctx, menuCategory.UUID).Return(&menuCategory, http.StatusOK, nil)
	// -- prepare for mocking

	resp, code, err := testMenuCategoryUsecase.DeleteMenuCategory(ctx, menuCategory.UUID)

	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}

func TestDeleteMenuCategory_successWhenDeleteMenuCategoryThatDoesntExists(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	mockMenuCategoryRepository.Mock.On("DeleteMenuCategory", ctx, normalUpsertRequest.UUID).Return(nil, http.StatusNotFound, nil)

	resp, code, err := testMenuCategoryUsecase.DeleteMenuCategory(ctx, normalUpsertRequest.UUID)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}