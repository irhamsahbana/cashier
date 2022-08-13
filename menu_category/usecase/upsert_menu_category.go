package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *menuCategoryUsecase) UpsertMenuCategory(c context.Context, req *domain.MenuCategoryUpsertRequest) (*domain.MenuCategoryResponse, int, error) {
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

	var resp domain.MenuCategoryResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}

	return &resp, http.StatusOK, nil
}