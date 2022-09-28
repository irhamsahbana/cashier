package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *itemCategoryUsecase) UpsertItemAndVariants(c context.Context, branchId, itemCategoryId string, req *dto.ItemAndVariantsUpsertRequest) (*dto.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.Item
	data.UUID = req.UUID
	data.Name = req.Name
	data.Price = req.Price
	data.Label = req.Label
	data.Description = req.Description
	data.Public = req.Public
	data.ImagePath = req.ImagePath
	data.Description = req.Description
	data.CreatedAt = time.Now().UnixMicro()

	var variants []domain.Variant
	for _, v := range req.Variants {
		var variant domain.Variant

		if err := validator.IsUUID(v.UUID); err != nil {
			return nil, http.StatusUnprocessableEntity, err
		}

		variant.UUID = v.UUID
		variant.Label = v.Label
		variant.Price = v.Price
		variant.Public = v.Public
		variant.ImagePath = v.ImagePath
		variant.CreatedAt = time.Now().UnixMicro()
		variants = append(variants, variant)
	}
	data.Variants = variants

	result, code, err := u.itemCategoryRepo.UpsertItemAndVariants(ctx, branchId, itemCategoryId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.ItemResponse
	itemCategoryDomainToDTO_UpsertItemAndVariants(&resp, result)

	return &resp, code, nil
}

func itemCategoryDomainToDTO_UpsertItemAndVariants(resp *dto.ItemResponse, result *domain.ItemCategory) {
	for _, item := range result.Items {
		if item.UUID != resp.UUID {
			continue
		}

		resp.UUID = item.UUID
		resp.Name = item.Name
		resp.Price = item.Price
		resp.Label = item.Label
		resp.Public = item.Public
		resp.ImagePath = item.ImagePath
		resp.Description = item.Description
		resp.CreatedAt = time.UnixMicro(item.CreatedAt)
		if item.UpdatedAt != nil {
			respUpdatedAt := time.UnixMicro(*item.UpdatedAt).UTC()
			resp.UpdatedAt = &respUpdatedAt
		}

		var variants []dto.VariantResponse
		for _, v := range item.Variants {
			var variant dto.VariantResponse

			variant.UUID = v.UUID
			variant.Price = v.Price
			variant.Label = v.Label
			variant.Public = v.Public
			variant.ImagePath = v.ImagePath
			variant.CreatedAt = time.UnixMicro(v.CreatedAt).UTC()

			variants = append(variants, variant)
		}
		resp.Variants = variants

		if len(resp.Variants) == 0 {
			resp.Variants = make([]dto.VariantResponse, 0)
		}
	}
}
