package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type userUsecase struct {
	userRepo           domain.UserRepositoryContract
	userRoleRepo       domain.UserRoleRepositoryContract
	tokenRepo          domain.TokenRepositoryContract
	companyRepo        domain.CompanyRepositoryContract
	branchDiscountRepo domain.BranchDiscountRepositoryContract
	contextTimeout     time.Duration
}

func NewUserUsecase(
	u domain.UserRepositoryContract,
	ur domain.UserRoleRepositoryContract,
	t domain.TokenRepositoryContract,
	c domain.CompanyRepositoryContract,
	bd domain.BranchDiscountRepositoryContract,
	timeout time.Duration) domain.UserUsecaseContract {
	return &userUsecase{
		userRepo:           u,
		userRoleRepo:       ur,
		tokenRepo:          t,
		companyRepo:        c,
		branchDiscountRepo: bd,
		contextTimeout:     timeout,
	}
}
