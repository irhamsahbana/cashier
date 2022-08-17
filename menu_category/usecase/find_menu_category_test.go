package usecase_test

import (
	"errors"
	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"lucy/cashier/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindMenuCategory_normalCase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
	updatedAtUnix := updatedAt.UTC().UnixMicro()

	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
	deletedAtUnix := deletedAt.UTC().UnixMicro()

	menuCategory := domain.MenuCategory{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: createdAtUnix,
	}

	mockMenuCategoryRepository.On("FindMenuCategory", ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false).Return(&menuCategory, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.FindMenuCategory(ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false)

	assert.NotNil(t, resp)
	assert.Equal(t, code, http.StatusOK)
	assert.Nil(t, err)

	t.Run("should convert from unix time to time.Time", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		menuWithUpdatedAndDeleted := menuCategory

		menuWithUpdatedAndDeleted.UpdatedAt = &updatedAtUnix
		menuWithUpdatedAndDeleted.DeletedAt = &deletedAtUnix

		mockMenuCategoryRepository.On("FindMenuCategory", ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false).Return(&menuWithUpdatedAndDeleted, http.StatusOK, nil)

		resp, code, err := testMenuCategoryUsecase.FindMenuCategory(ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false)

		assert.NotNil(t, resp)
		assert.Equal(t, code, http.StatusOK)
		assert.Nil(t, err)
	})
}

func TestFindMenuCategory_ErrorWhenMenuCategoryDoesntExists(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	mockMenuCategoryRepository.On("FindMenuCategory", ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false).Return(nil, http.StatusNotFound, errors.New("Menu category not found"))

	resp, code, err := testMenuCategoryUsecase.FindMenuCategory(ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusNotFound, code)
	assert.NotNil(t, err)

	assert.Error(t, err, "Menu category not found")
}

func TestFindMenuCategory_ErrorWhenSomethingWrongWithMarshalingOrDatabase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	mockMenuCategoryRepository.On("FindMenuCategory", ctx, "74c4a96b-b19c-4c32-9b94-d13f533144fe", false).Return(nil, http.StatusInternalServerError, errors.New("mongo: request timeout"))

	resp, code, err := testMenuCategoryUsecase.FindMenuCategory(ctx,"74c4a96b-b19c-4c32-9b94-d13f533144fe", false)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, code)
	assert.NotNil(t, err)
}