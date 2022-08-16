package usecase_test

import (
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindMenu(t *testing.T) {
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
		DeletedAt: &deletedAtUnix,
	}

	menuCategory := domain.MenuCategory{
		UUID: "73d29512-a416-4362-b5ba-688e330b477b",
		Name: "Coffee Base",
		Menus: []domain.Menu{menu},
	}

	mockMenuCategoryRepository.Mock.On("FindMenu", ctx, menu.UUID, true).Return(&menuCategory, http.StatusOK, nil)

	resp, code, err := testMenuCategoryUsecase.FindMenu(ctx, menu.UUID, true)

	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	t.Run("should return error when response from repository got wrong", func(t *testing.T) {
		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

		mockMenuCategoryRepository.Mock.On("FindMenu", ctx, menu.UUID, true).Return(nil, http.StatusInternalServerError, errors.New("something wrong"))

		resp, code, err := testMenuCategoryUsecase.FindMenu(ctx, menu.UUID, true)

		assert.Nil(t, resp)
		assert.Equal(t, http.StatusInternalServerError, code)
		assert.NotNil(t, err)
	})
}