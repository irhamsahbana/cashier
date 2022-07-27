package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib"
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

func (usecase *menuCategoryUsecase) CreateMenuCategory(c context.Context, data *domain.MenuCategory) (*domain.MenuCategory, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	if ok, err := lib.ValidatorUUID(data.UUID); !ok {
		return nil, err, http.StatusBadRequest
	}

	res, err := usecase.menuCategoryRepo.InsertOne(ctx, data)
	if err != nil {
		return res, err, http.StatusInternalServerError
	}

	return res, nil, http.StatusCreated
}

func (usecase *menuCategoryUsecase) FindMenuCategory(c context.Context, id string) (*domain.MenuCategory, error, int) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindOne(ctx, id)
	if err != nil {
		return res, err, http.StatusInternalServerError
	}

	return res, nil, http.StatusOK
}