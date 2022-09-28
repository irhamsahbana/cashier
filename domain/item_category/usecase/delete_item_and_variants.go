package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/dto"
)

func (u *itemCategoryUsecase) DeleteItemAndVariants(c context.Context, branchId, id string) (*dto.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.DeleteItemAndVariants(ctx, branchId, id)
	if err != nil {
		if code == http.StatusNotFound {
			return nil, http.StatusOK, nil
		}

		return nil, code, err
	}

	var resp dto.ItemResponse
	itemCategoryDomainToDTO_DeleteItemAndVariants(&resp, result)

	return &resp, http.StatusOK, nil
}

func itemCategoryDomainToDTO_DeleteItemAndVariants(resp *dto.ItemResponse, result *domain.ItemCategory) {
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

		var variants []dto.VariantResponse
		for _, v := range item.Variants {
			var dataVariant dto.VariantResponse

			dataVariant.UUID = v.UUID
			dataVariant.Label = v.Label
			dataVariant.Price = v.Price
			dataVariant.Public = v.Public
			dataVariant.ImagePath = v.ImagePath
			variants = append(variants, dataVariant)
		}
		resp.Variants = variants
	}
}
