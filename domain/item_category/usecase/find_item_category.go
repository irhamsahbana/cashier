package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) FindItemCategory(c context.Context, id string, withTrashed bool) (*domain.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemCategory(ctx, id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp domain.ItemCategoryResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if result.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*result.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}
