package usecase_test

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var MockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(MockMenuCategoryRepository, time.Duration(5) * time.Second)
var ctx =  context.Background()

var normalUpsertRequest = domain.MenuCategoryUpsertRequest{
	UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
	Name: "Coffee Base",
	CreatedAt: "2022-08-13T04:06:16.312789Z",
}

func TestUpsertMenuCategory_normalCase(t *testing.T) {
	createdAtString :=  normalUpsertRequest.CreatedAt
	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UnixMicro()

	// prepare for mocking
	menuCategory := domain.MenuCategory{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: createdAtUnix,
	}

	MockMenuCategoryRepository.On("UpsertMenuCategory", ctx, &menuCategory).Return(&menuCategory, http.StatusOK, nil)
	// -- prepare for mocking

	// testing usecase with fake request
	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &normalUpsertRequest)
	// -- testing usecase with fake request

	// assertion section
	assert.NotNil(t, resp)
	assert.Equal(t, code, http.StatusOK)
	assert.Nil(t, err)

	assert.Equal(t, "74c4a96b-b19c-4c32-9b94-d13f533144fe", resp.UUID)
	assert.Equal(t, "Coffee Base", resp.Name)
	assert.Equal(t, createdAt, resp.CreatedAt)
}

func TestUpsertMenuCategory_errorWhenUuidIsEmptyString(t *testing.T) {
	request := normalUpsertRequest
	request.UUID = ""

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusUnprocessableEntity, code)
	assert.NotNil(t, err)

	assert.EqualError(t, err, "invalid UUID length: 0")
}

func TestUpsertMenuCategory_errorWhenCreatedAtNotValidRFC3999(t *testing.T) {
	request := normalUpsertRequest
	request.CreatedAt = "2022-08-13T04:06:13.dsadaZ"

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

	assert.Nil(t, resp)
	assert.Equal(t, http.StatusUnprocessableEntity, code)
	assert.NotNil(t, err)
}