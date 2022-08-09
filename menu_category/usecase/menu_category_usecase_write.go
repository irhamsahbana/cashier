package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *menuCategoryUsecase) UpsertMenuCategory(c context.Context, req *domain.MenuCategoryUpsertRequest) (*domain.MenuCategoryUpsertResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()


	if  err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var menucategory domain.MenuCategory
	menucategory.UUID = req.UUID
	menucategory.Name = req.Name
	menucategory.CreatedAt = createdAt.UnixMicro()

	result, code, err := u.menuCategoryRepo.UpsertMenuCategory(ctx, &menucategory)
	if err != nil {
		return nil, code, err
	}

	var resp domain.MenuCategoryUpsertResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}

	return &resp, http.StatusOK, nil
}

func (u *menuCategoryUsecase) DeleteMenuCategory(c context.Context, id string) (*domain.MenuCategoryDeleteResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.DeleteMenuCategory(ctx, id)
	if err != nil {
		return nil, code, err
	}

	var resp domain.MenuCategoryDeleteResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}
	if result.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*result.DeletedAt)
		resp.DeletedAt = respDeletedAt
	}

	return &resp, http.StatusOK, nil
}

// Menu

func (u *menuCategoryUsecase) CreateMenu(c context.Context, menuCategoryId string, req *domain.MenuCreateRequest) (*domain.MenuCreateResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if  err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	data := domain.Menu{
		UUID: req.UUID,
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
		Label: req.Label,
		Public: req.Public,
		CreatedAt: createdAt.UnixMicro(),
	}

	result, code, err := u.menuCategoryRepo.InsertMenu(ctx, menuCategoryId, &data)
	if err != nil {
		return nil, code, err
	}

	menu := result.Menus[0]

	var resp domain.MenuCreateResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt)

	return &resp, http.StatusCreated, nil
}

func (u *menuCategoryUsecase) DeleteMenu(c context.Context, id string) (*domain.MenuDeleteResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.DeleteMenu(ctx, id)
	if err != nil {
		return nil, code, err
	}

	menu := result.Menus[0]

	var resp domain.MenuDeleteResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt)
	if menu.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*menu.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}
	if menu.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*menu.DeletedAt)
		resp.DeletedAt = respDeletedAt
	}

	return &resp, http.StatusOK, nil
}