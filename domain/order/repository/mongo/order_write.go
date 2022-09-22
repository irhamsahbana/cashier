package mongo

import (
	"context"
	"lucy/cashier/domain"
)

func (repo *orderRepository) UpsertOrder(ctx context.Context, data *domain.OrderGroup) (*domain.OrderGroup, int, error) {
	panic("implement me")
}
