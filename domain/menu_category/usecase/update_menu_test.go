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

func TestUpdateMenu(t *testing.T) {
	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
	updatedAtUnix := updatedAt.UTC().UnixMicro()

	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
	deletedAtUnix := deletedAt.UTC().UnixMicro()

	menu := domain.Menu{
		UUID: "fdbacc80-5729-4c26-9557-fbbfbacd630a",
		Name: "Cappucino",
		Price: 23000,
		Label: "Coffee",
		Public: true,
		CreatedAt: createdAtUnix,
		UpdatedAt: &updatedAtUnix,
	}

	menuCategory := domain.MenuCategory{
		UUID: "73d29512-a416-4362-b5ba-688e330b477b",
		Name: "Coffee Base",
		Menus: []domain.Menu{menu},
	}

	request := domain.MenuUpdateRequest{
		Name: menu.Name,
		Price: menu.Price,
		Label: menu.Label,
		Public: menu.Public,
	}

	dataInput := domain.Menu{
		Name: menu.Name,
		Price: menu.Price,
		Description: menu.Description,
		Label: menu.Label,
		Public: menu.Public,
	}

	mockMenuCategoryRepository.Mock.On("UpdateMenu", ctx, menu.UUID, &dataInput).Return(&menuCategory, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menu.UUID, &request)

	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	t.Run("should return error when name is empty", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		emptyNameRequest := request
		emptyNameRequest.Name = ""

		resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menu.UUID, &emptyNameRequest)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.NotNil(t, err)
	})

	t.Run("should return error when label is empty", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		emptyLabelRequest := request
		emptyLabelRequest.Label = ""

		resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menu.UUID, &emptyLabelRequest)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.NotNil(t, err)
	})

	t.Run("should return error if something wrong with repository", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository,timeoutContext)

		mockMenuCategoryRepository.On("UpdateMenu", ctx, menuCategory.UUID, &dataInput).Return(nil, http.StatusInternalServerError, errors.New("something wrong"))

		resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menuCategory.UUID, &request)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, code)
		assert.NotNil(t, err)
	})

	t.Run("should return deleted_at field if resource has been deleted", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		menuDeleted := menu
		menuDeleted.DeletedAt = &deletedAtUnix

		menuCategory := domain.MenuCategory{
			UUID: "73d29512-a416-4362-b5ba-688e330b477b",
			Name: "Coffee Base",
			Menus: []domain.Menu{menuDeleted},
		}

		mockMenuCategoryRepository.Mock.On("UpdateMenu", ctx, menu.UUID, &dataInput).Return(&menuCategory, http.StatusOK, nil)

		resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menu.UUID, &request)

		assert.NotNil(t, resp)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)
	})

	t.Run("should return nothing when nothing to update and response code must be 200", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		mockMenuCategoryRepository.Mock.On("UpdateMenu", ctx, menu.UUID, &dataInput).Return(nil, http.StatusNotFound, nil)

		resp, code, err := testMenuCategoryUsecase.UpdateMenu(ctx, menu.UUID, &request)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)

	})
}