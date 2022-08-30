package usecase

import (
	"time"

	"lucy/cashier/domain"
)

type itemCategoryUsecase struct {
	itemCategoryRepo domain.ItemCategoryRepositoryContract
	contextTimeout   time.Duration
}

func NewItemCategoryUsecase(repo domain.ItemCategoryRepositoryContract, timeout time.Duration) domain.ItemCategoryUsecaseContract {
	return &itemCategoryUsecase{
		itemCategoryRepo: repo,
		contextTimeout:   timeout,
	}
}
