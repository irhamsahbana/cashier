package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type userRoleUsecase struct {
	userRoleRepo   domain.UserRoleRepositoryContract
	contextTimeout time.Duration
}

func NewUserRoleUsecase(u domain.UserRoleRepositoryContract, timeout time.Duration) domain.UserRoleUsecaseContract {
	return &userRoleUsecase{
		userRoleRepo:   u,
		contextTimeout: timeout,
	}
}
