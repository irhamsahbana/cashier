package mocks

import (
	"context"
	"lucy/cashier/domain"
)

func (mock *MockItemCategoryRepository) UpsertItemCategory(ctx context.Context, data *domain.ItemCategory) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), data)

	var entitiy *domain.ItemCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.ItemCategory)
		entitiy = assertion
	}
	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entitiy, code, args.Error(2)
}

func (mock *MockItemCategoryRepository) DeleteItemCategory(ctx context.Context, id string) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), id)

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

func (mock *MockItemCategoryRepository) InsertItem(ctx context.Context, ItemCategoryId string, data *domain.Item) (*domain.ItemCategory, int, error) {
	// return &domain.ItemCategory{}, http.StatusOK, nil

	args := mock.Called(context.Background(), ItemCategoryId, data)

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

func (mock *MockItemCategoryRepository) UpdateItem(ctx context.Context, id string, data *domain.Item) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), id, data)

	var entitiy *domain.ItemCategory
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.ItemCategory)
		entitiy = assertion
	}
	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entitiy, code, args.Error(2)
}

func (mock *MockItemCategoryRepository) DeleteItem(ctx context.Context, id string) (*domain.ItemCategory, int, error) {
	args := mock.Called(context.Background(), id)

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
