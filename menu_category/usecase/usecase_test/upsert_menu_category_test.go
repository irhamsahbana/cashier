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

func TestUpsertMenuCategory_normalCase(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

	createdAtString :=  normalUpsertRequest.CreatedAt
	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UnixMicro()

	menuCategory := domain.MenuCategory{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: createdAtUnix,
	}

	mockMenuCategoryRepository.On("UpsertMenuCategory", ctx, &menuCategory).Return(&menuCategory, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &normalUpsertRequest)

	assert.NotNil(t, resp)
	assert.Equal(t, code, http.StatusOK)
	assert.Nil(t, err)

	assert.Equal(t, "74c4a96b-b19c-4c32-9b94-d13f533144fe", resp.UUID)
	assert.Equal(t, "Coffee Base", resp.Name)
	assert.Equal(t, createdAt, resp.CreatedAt)
}

func TestUpsertMenuCategory_errorWhenUuidIsEmptyString(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

	request := normalUpsertRequest
	request.UUID = ""

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusUnprocessableEntity, code)
	assert.NotNil(t, err)

	assert.EqualError(t, err, "invalid UUID length: 0")
}

func TestUpsertMenuCategory_errorWhenCreatedAtNotValidRFC3999(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

	request := normalUpsertRequest
	request.CreatedAt = "2022-08-13T04:06:13.dsadaZ"

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusUnprocessableEntity, code)
	assert.NotNil(t, err)
}