package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

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