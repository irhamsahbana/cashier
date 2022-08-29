package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepositoryContract
	tokenRepo      domain.TokenRepositoryContract
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepositoryContract, t domain.TokenRepositoryContract, timeout time.Duration) domain.UserUsecaseContract {
	return &userUsecase{
		userRepo:       u,
		tokenRepo:      t,
		contextTimeout: timeout,
	}
}
