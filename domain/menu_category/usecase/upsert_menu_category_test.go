package usecase_test

import (
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/domain/menu_category/mocks"
	"lucy/cashier/domain/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpsertMenuCategory(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

	var request = domain.MenuCategoryUpsertRequest{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: createdAtString,
	}

	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
	updatedAtUnix := updatedAt.UTC().UnixMicro()

	menuCategory := domain.MenuCategory{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "Coffee Base",
		CreatedAt: createdAtUnix,
	}

	mockMenuCategoryRepository.On("UpsertMenuCategory", ctx, &menuCategory).Return(&menuCategory, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

	assert.NotNil(t, resp)
	assert.Equal(t, code, http.StatusOK)
	assert.Nil(t, err)

	assert.Equal(t, "74c4a96b-b19c-4c32-9b94-d13f533144fe", resp.UUID)
	assert.Equal(t, "Coffee Base", resp.Name)
	assert.Equal(t, createdAt, resp.CreatedAt)

	t.Run("should convert unix time in database to date.Time", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

		var request = domain.MenuCategoryUpsertRequest{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			Name: "Coffee Base",
			CreatedAt: createdAtString,
		}

		menuCategoryInput := domain.MenuCategory{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			Name: "Coffee Base",
			CreatedAt: createdAtUnix,
		}

		menuCategoryOutput := domain.MenuCategory{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			Name: "Coffee Base",
			CreatedAt: createdAtUnix,
			UpdatedAt: &updatedAtUnix,
		}

		mockMenuCategoryRepository.On("UpsertMenuCategory", ctx, &menuCategoryInput).Return(&menuCategoryOutput, http.StatusOK, nil)

		resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

		assert.Equal(t, &updatedAt, resp.UpdatedAt)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)
	})

	t.Run("should return error when uuid is empty string", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

		var request = domain.MenuCategoryUpsertRequest{
			Name: "Coffee Base",
			CreatedAt: createdAtString,
		}

		resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.NotNil(t, err)

		assert.EqualError(t, err, "invalid UUID length: 0")
	})

	t.Run("should return error when created_at field not a valid RFC3999", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

		var request = domain.MenuCategoryUpsertRequest{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			Name: "Coffee Base",
			CreatedAt: "2022-08-13T04:06:13.dsadaZ",
		}

		resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.NotNil(t, err)
	})

	t.Run("should return error if something wrong with repository", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

		var request = domain.MenuCategoryUpsertRequest{
			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
			Name: "Coffee Base",
			CreatedAt: createdAtString,
		}

		mockMenuCategoryRepository.On("UpsertMenuCategory", ctx, &menuCategory).Return(nil, http.StatusInternalServerError, errors.New("something wrong"))

		resp, code, err := testMenuCategoryUsecase.UpsertMenuCategory(ctx, &request)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, code)
		assert.NotNil(t, err)
	})
}