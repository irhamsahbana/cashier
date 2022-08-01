package usecase

import (
	"context"
	"errors"
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

func (usecase *menuCategoryUsecase) UpdateMenuCategory(c context.Context, id string, data *domain.MenuCategoryUpdateRequest) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res := &domain.MenuCategory{}

	if data.Name == "" {
		return res, http.StatusUnprocessableEntity, errors.New("Name needed")
	}

	res, err := usecase.menuCategoryRepo.UpdateMenuCategory(ctx, id, data)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) CreateMenu(c context.Context, menuCategoryId string, request *domain.MenuCreateRequestResponse) (*domain.MenuCreateRequestResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	menuData := domain.Menu{
		UUID: request.UUID,
		Name: request.Name,
		Price: request.Price,
		Description: request.Description,
		Label: request.Label,
		Public: request.Public,
		CreatedAt: request.CreatedAt,
	}

	res, err := usecase.menuCategoryRepo.InsertMenu(ctx, menuCategoryId, &menuData)
	if err != nil {
		return request, http.StatusInternalServerError, err
	}

	resp := &domain.MenuCreateRequestResponse{
		UUID: res.UUID,
		Name: res.Name,
		Price: res.Price,
		Description: res.Description,
		Label: res.Label,
		Public: res.Public,
		CreatedAt: res.CreatedAt,
	}
	return resp, http.StatusCreated, nil
}