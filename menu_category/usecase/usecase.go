package usecase

import (
	"time"

	"lucy/cashier/domain"
)

type menuCategoryUsecase struct {
	menuCategoryRepo	domain.MenuCategoryRepositoryContract
	contextTimeout		time.Duration
}

func NewMenuCategoryUsecase(repo domain.MenuCategoryRepositoryContract, timeout time.Duration) domain.MenuCategoryUsecaseContract {
	return &menuCategoryUsecase{
		menuCategoryRepo: repo,
		contextTimeout: timeout,
	}
}