package usecase

import (
	"context"
	"lucy/cashier/domain"
)

func (u *waiterUsecase) FindWaiter(ctx context.Context, id string) (*domain.WaiterResponse, int, error) {
	return nil, 200, nil
}