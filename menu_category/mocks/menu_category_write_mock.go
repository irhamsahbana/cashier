package mocks

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
)


func(mock *MockMenuCategoryRepository) UpsertMenuCategory(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, int, error) {
	args := mock.Called(context.Background(), data)

	var entitiy *domain.MenuCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.MenuCategory)
		entitiy = assertion
	}
	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entitiy, code, args.Error(2)
}

func(mock *MockMenuCategoryRepository) DeleteMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	args := mock.Called(context.Background(), id)

	var entity *domain.MenuCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.MenuCategory)
		entity = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entity, code, args.Error(2)
}

func(mock *MockMenuCategoryRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) UpdateMenu(ctx context.Context, id string, data *domain.Menu) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}

func(mock *MockMenuCategoryRepository) DeleteMenu(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	return &domain.MenuCategory{}, http.StatusOK, nil
}