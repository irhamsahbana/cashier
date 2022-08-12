package mocks

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockMenuCategoryRepository struct {
	mock.Mock
}

func(mock *MockMenuCategoryRepository) UpsertMenuCategory(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, int, error) {
	args := mock.Called(ctx, data)

	dataArg := args.Get(1)

	return dataArg.(*domain.MenuCategory), http.StatusOK, args.Error(0)
}

func(mock *MockMenuCategoryRepository) FindMenuCategories(ctx context.Context, withTrashed bool) ([]domain.MenuCategory, int, error) {
	return nil, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) DeleteMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) UpdateMenu(ctx context.Context, id string, data *domain.Menu) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) FindMenu(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) DeleteMenu(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}