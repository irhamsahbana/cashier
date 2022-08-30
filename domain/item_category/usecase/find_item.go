package usecase

import (
	"context"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) FindItem(c context.Context, id string, withTrashed bool) (*domain.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItem(ctx, id, withTrashed)
	if err != nil {
		return nil, code, err
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
		itemUpdatedAt := time.UnixMicro(*item.UpdatedAt).UTC()
		resp.UpdatedAt = &itemUpdatedAt
	}
	if item.DeletedAt != nil {
		itemDeletedAt := time.UnixMicro(*item.DeletedAt).UTC()
		resp.DeletedAt = &itemDeletedAt
	}

	return &resp, code, nil
}
