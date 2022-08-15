package usecase_test

import (
	"errors"
	"testing"

	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"time"

	"lucy/cashier/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindMenuCategories_normalCase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	createdAtString :=  normalUpsertRequest.CreatedAt
	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	menuCategories := []domain.MenuCategory{
		{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			BranchUUID: "56ede4f9-91a6-4189-a13d-c9cc3e767d53",
			Name: "Coffee Base",
			CreatedAt: createdAtUnix,
		},
		{
			UUID: "09f04bb0-2bd7-4af0-9dc5-c59af7f938de",
			BranchUUID: "5ece4124-e588-404f-a068-1c3c503768fe",
			Name: "Juice",
			CreatedAt: createdAtUnix,
			UpdatedAt: &createdAtUnix,
		},
	}

	mockMenuCategoryRepository.On("FindMenuCategories", ctx, false).Return(menuCategories, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.FindMenuCategories(ctx, false)

	assert.NotNil(t, resp)
	assert.Equal(t, code, http.StatusOK)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(resp))

	assert.Equal(t, createdAt, resp[0].CreatedAt)
}

func TestFindMenuCategories_ErrorWhenSomethingWrongWithMarshalingOrDatabase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	mockMenuCategoryRepository.On("FindMenuCategories", ctx, false).Return(nil, http.StatusInternalServerError, errors.New("mongo: request timeout"))

	resp, code, err := testMenuCategoryUsecase.FindMenuCategories(ctx, false)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, code)
	assert.NotNil(t, err)
}