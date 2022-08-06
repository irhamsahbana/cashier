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

func (usecase *menuCategoryUsecase) FindMenuCategory(c context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindMenuCategory(ctx, id, withTrashed)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) FindMenuCategories(c context.Context, withTrashed bool) ([]domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindMenuCategories(ctx, withTrashed)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) DeleteMenuCategory(c context.Context, id string, forceDelete bool) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.DeleteMenuCategory(ctx, id, forceDelete)
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

// Menu

func (usecase *menuCategoryUsecase) CreateMenu(c context.Context, menuCategoryId string, request *domain.MenuCreateRequestResponse) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	var menucategory domain.MenuCategory

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
		return &menucategory, http.StatusInternalServerError, err
	}

	return res, http.StatusCreated, nil
}

func (usecase *menuCategoryUsecase) FindMenu(c context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindMenu(ctx, id, withTrashed)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

func (usecase *menuCategoryUsecase) DeleteMenu(c context.Context, id string, forceDelete bool) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.DeleteMenu(ctx, id, forceDelete)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}