package usecase

import (
	"context"
	"net/http"
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

func (u *menuCategoryUsecase) FindMenuCategories(c context.Context, withTrashed bool) ([]domain.MenuCategoryFindAllResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.FindMenuCategories(ctx, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.MenuCategoryFindAllResponse

	for _, mc := range result {
		var data domain.MenuCategoryFindAllResponse

		data.UUID = mc.UUID
		data.Name = mc.Name
		data.CreatedAt = time.UnixMicro(mc.CreatedAt)
		if mc.UpdatedAt != nil {
			dataUpdatedAt := time.UnixMicro(*mc.UpdatedAt)
			data.UpdatedAt = &dataUpdatedAt
		}
		if mc.DeletedAt != nil {
			dataDeletedAt := time.UnixMicro(*mc.DeletedAt)
			data.DeletedAt = &dataDeletedAt
		}

		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}

func (u *menuCategoryUsecase) FindMenuCategory(c context.Context, id string, withTrashed bool) (*domain.MenuCategoryFindResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.FindMenuCategory(ctx, id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp domain.MenuCategoryFindResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}
	if result.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*result.DeletedAt)
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}

// Menu

func (u *menuCategoryUsecase) FindMenu(c context.Context, id string, withTrashed bool) (*domain.MenuFindResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, err := u.menuCategoryRepo.FindMenu(ctx, id, withTrashed)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	menu := result.Menus[0]
	var resp domain.MenuFindResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public	= menu.Public
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if menu.UpdatedAt != nil {
		menuUpdatedAt := time.UnixMicro(*menu.UpdatedAt)
		resp.UpdatedAt = &menuUpdatedAt
	}
	if menu.DeletedAt != nil {
		menuDeletedAt := time.UnixMicro(*menu.DeletedAt)
		resp.DeletedAt = &menuDeletedAt
	}

	return &resp, http.StatusOK, nil
}
