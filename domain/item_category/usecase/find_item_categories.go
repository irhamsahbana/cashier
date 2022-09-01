package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) FindItemCategories(c context.Context, withTrashed bool) ([]domain.ItemCategoryFindAllResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemCategories(ctx, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.ItemCategoryFindAllResponse

	// mc -> item category, m -> item
	for _, mc := range result {
		var data domain.ItemCategoryFindAllResponse
		var items []domain.ItemFindAllResponse

		data.UUID = mc.UUID
		data.Name = mc.Name
		data.CreatedAt = time.UnixMicro(mc.CreatedAt).UTC()
		if mc.UpdatedAt != nil {
			dataUpdatedAt := time.UnixMicro(*mc.UpdatedAt).UTC()
			data.UpdatedAt = &dataUpdatedAt
		}
		if mc.DeletedAt != nil {
			dataDeletedAt := time.UnixMicro(*mc.DeletedAt).UTC()
			data.DeletedAt = &dataDeletedAt
		}

		if len(mc.Items) > 0 {
			for _, m := range mc.Items {
				var dataItem domain.ItemFindAllResponse

				dataItem.UUID = m.UUID
				dataItem.MainUUID = m.MainUUID
				dataItem.Name = m.Name
				dataItem.Price = m.Price
				dataItem.Description = m.Description
				dataItem.Label = m.Label
				dataItem.Public = m.Public
				dataItem.ImageUrl = m.ImageUrl
				dataItem.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
				if m.UpdatedAt != nil {
					dataUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
					dataItem.UpdatedAt = &dataUpdatedAt
				}
				if m.DeletedAt != nil {
					dataDeletedAt := time.UnixMicro(*m.DeletedAt).UTC()
					dataItem.DeletedAt = &dataDeletedAt
				}

				items = append(items, dataItem)
			}
		}

		data.Items = items
		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}
