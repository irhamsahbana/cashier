package mongo

import (
	"context"
	"lucy/cashier/domain"
)

func (r *orderRepository) InsertRefund(ctx context.Context, branchId string, data *domain.Refund) ([]domain.Refund, int, error) {
	panic("implement me")
}
