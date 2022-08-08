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

func (usecase *menuCategoryUsecase) UpsertMenuCategory(c context.Context, data *domain.MenuCategoryUpsertRequest) (*domain.MenuCategoryUpsertResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()


	if  err := validator.IsUUID(data.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, data.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var menucategory domain.MenuCategory
	menucategory.UUID = data.UUID
	menucategory.Name = data.Name
	menucategory.CreatedAt = createdAt.UnixMicro()

	result, err := usecase.menuCategoryRepo.UpsertMenuCategory(ctx, &menucategory)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var resp domain.MenuCategoryUpsertResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	respUpdatedAt := time.UnixMicro(*result.UpdatedAt)
	resp.UpdatedAt = &respUpdatedAt

	return &resp, http.StatusOK, nil
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

func (usecase *menuCategoryUsecase) FindMenuCategory(c context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err := usecase.menuCategoryRepo.FindMenuCategory(ctx, id, withTrashed)
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