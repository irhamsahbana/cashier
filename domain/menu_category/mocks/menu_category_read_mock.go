package mocks

import (
	"context"
	"lucy/cashier/domain"

	"github.com/stretchr/testify/mock"
)

type MockMenuCategoryRepository struct {
	mock.Mock
}

func(mock *MockMenuCategoryRepository) FindMenuCategories(ctx context.Context, withTrashed bool) ([]domain.MenuCategory, int, error) {
	args := mock.Called(context.Background(), withTrashed)

	var entities []domain.MenuCategory
	var code int

	entitiesArg := args.Get(0)
	codeArg := args.Get(1)

	if entitiesArg != nil {
		assertion, _ := entitiesArg.([]domain.MenuCategory)
		entities = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entities, code, args.Error(2)
}

func(mock *MockMenuCategoryRepository) FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	args := mock.Called(context.Background(), id, withTrashed)

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

func(mock *MockMenuCategoryRepository) FindMenu(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	args := mock.Called(context.Background(), id, withTrashed)

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