package usecase

import (
	"context"
	"lucy/cashier/domain"
)

func (u *waiterUsecase) UpsertWaiter(ctx context.Context, data *domain.WaiterUpsertrequest) (*domain.WaiterResponse, int, error) {
	return nil, 200, nil
}