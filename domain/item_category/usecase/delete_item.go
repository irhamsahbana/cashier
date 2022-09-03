package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) DeleteItem(c context.Context, branchId, id string) (*domain.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.DeleteItem(ctx, branchId, id)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	item := result.Items[0]

	var resp domain.ItemResponse
	resp.UUID = item.UUID
	resp.Name = item.Name
	resp.Price = item.Price
	resp.Description = item.Description
	resp.Label = item.Label
	resp.Public = item.Public
	resp.CreatedAt = time.UnixMicro(item.CreatedAt).UTC()
	if item.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*item.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if item.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*item.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}
