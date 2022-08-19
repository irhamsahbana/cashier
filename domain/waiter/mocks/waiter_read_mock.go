package mocks

import (
	"context"
	"lucy/cashier/domain"
)

func (mock *MockWaiterRepository) UpsertWaiter(ctx context.Context, data *domain.Waiter) (*domain.Waiter, int, error) {
	return nil, 200, nil
}