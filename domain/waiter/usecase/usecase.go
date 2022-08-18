package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type waiterUsecase struct {
	waiterRepo domain.WaiterRepositoryContract
	contextTimeout time.Duration
}

func NewWaiterUsecase(repo domain.WaiterRepositoryContract, timeout time.Duration) domain.WaiterUsecaseContract {
	return &waiterUsecase{
		waiterRepo: repo,
		contextTimeout: timeout,
	}
}