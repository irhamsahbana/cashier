package usecase

import (
	"lucy/cashier/domain"
	"time"
)

func NewOrderUsecase(repo domain.OrderRepositoryContract, timeout time.Duration) domain.OrderUsecaseContract {
	return &orderUsecase{
		orderRepo:      repo,
		contextTimeout: timeout,
	}
}

type orderUsecase struct {
	orderRepo      domain.OrderRepositoryContract
	contextTimeout time.Duration
}
