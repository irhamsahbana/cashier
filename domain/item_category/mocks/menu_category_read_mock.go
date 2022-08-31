package mocks

import (
	"context"
	"lucy/cashier/domain"

	"github.com/stretchr/testify/mock"
)

type MockItemCategoryRepository struct {
	mock.Mock
}

func (mock *MockItemCategoryRepository) FindItemCategories(ctx context.Context, withTrashed bool) ([]domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), withTrashed)

	var entities []domain.ItemCategory
	var code int

	entitiesArg := args.Get(0)
	codeArg := args.Get(1)

	if entitiesArg != nil {
		assertion, _ := entitiesArg.([]domain.ItemCategory)
		entities = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entities, code, args.Error(2)
}

func (mock *MockItemCategoryRepository) FindItemCategory(ctx context.Context, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), id, withTrashed)

	var entity *domain.ItemCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.ItemCategory)
		entity = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entity, code, args.Error(2)
}

func (mock *MockItemCategoryRepository) FindItem(ctx context.Context, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), id, withTrashed)

	var entity *domain.ItemCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.ItemCategory)
		entity = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entity, code, args.Error(2)
}
