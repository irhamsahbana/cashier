package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepositoryContract
	userRoleRepo   domain.UserRoleRepositoryContract
	tokenRepo      domain.TokenRepositoryContract
	contextTimeout time.Duration
}

func NewUserUsecase(
	u domain.UserRepositoryContract,
	ur domain.UserRoleRepositoryContract,
	t domain.TokenRepositoryContract,
	timeout time.Duration) domain.UserUsecaseContract {
	return &userUsecase{
		userRepo:       u,
		userRoleRepo:   ur,
		tokenRepo:      t,
		contextTimeout: timeout,
	}
}
