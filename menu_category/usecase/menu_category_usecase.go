package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
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

func (usecase *menuCategoryUsecase) CreateMenuCategory(c context.Context, data *domain.MenuCategory) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	if  err := validator.IsUUID(data.UUID); err != nil {
		return nil, http.StatusBadRequest, err
	}

	res, err := usecase.menuCategoryRepo.InsertMenuCategory(ctx, data)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusCreated, nil
}

func (usecase *menuCategoryUsecase) FindMenuCategory(c context.Context, id string) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindMenuCategory(ctx, id)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) DeleteMenuCategory(c context.Context, id string) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.DeleteMenuCategory(ctx, id)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) UpdateMenuCategory(c context.Context, id string, data *domain.MenuCategoryUpdatePayload) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.UpdateMenuCategory(ctx, id, data)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}