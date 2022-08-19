package mocks

import (
	"context"
	"lucy/cashier/domain"

	"github.com/stretchr/testify/mock"
)

type MockWaiterRepository struct {
	mock.Mock
}

func (mock *MockWaiterRepository) FindWaiter(ctx context.Context, id string, withTrashed bool) (*domain.Waiter, int, error) {
	args := mock.Called(context.Background(), id, withTrashed)

	var entity *domain.Waiter
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.Waiter)
		entity = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entity, code, args.Error(2)
}

func (mock *MockWaiterRepository) DeleteWaiter(ctx context.Context, id string) (*domain.Waiter, int, error) {
	args := mock.Called(context.Background(), id)

	var entity *domain.Waiter
	var code int

	entityArg := args.Get(0)
	codeArg := args.Get(1)

	if entityArg != nil {
		assertion, _ := entityArg.(*domain.Waiter)
		entity = assertion
	}

	if codeArg != nil {
		assertion, _ := codeArg.(int)
		code = assertion
	}

	return entity, code, args.Error(2)
}
