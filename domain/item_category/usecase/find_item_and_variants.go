package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *itemCategoryUsecase) FindItemAndVariants(c context.Context, branchId, id string) (*domain.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemAndVariants(ctx, branchId, id)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	var resp domain.ItemResponse

	for _, item := range result.Items {
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

		var variants []domain.VariantResponse
		for _, v := range item.Variants {
			var dataVariant domain.VariantResponse

			dataVariant.UUID = v.UUID
			dataVariant.Label = v.Label
			dataVariant.Price = v.Price
			dataVariant.Public = v.Public
			dataVariant.ImagePath = v.ImagePath
			variants = append(variants, dataVariant)
		}
		resp.Variants = variants

		return &resp, http.StatusOK, nil
	}

	return &resp, http.StatusOK, nil
}
