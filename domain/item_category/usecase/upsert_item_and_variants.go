package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *itemCategoryUsecase) UpsertItemAndVariants(c context.Context, branchId, itemCategoryId string, req *dto.ItemAndVariantsUpsertRequest) (*dto.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var data domain.Item
	DTOtoDomain_UpsertItemAndVariants(&data, req)

	result, code, err := u.itemCategoryRepo.UpsertItemAndVariants(ctx, branchId, itemCategoryId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.ItemResponse
	DomainToDTO_UpsertItemAndVariants(&resp, result, req)

	return &resp, code, nil
}

func DTOtoDomain_UpsertItemAndVariants(data *domain.Item, req *dto.ItemAndVariantsUpsertRequest) {
	data.UUID = req.UUID
	data.Name = req.Name
	data.Price = req.Price
	data.Label = req.Label
	data.Description = req.Description
	data.Public = req.Public
	data.ImagePath = req.ImagePath
	data.CreatedAt = time.Now().UnixMicro()

	variants := []domain.Variant{}
	for _, v := range req.Variants {
		var variant domain.Variant

		variant.UUID = v.UUID
		variant.Label = v.Label
		variant.Price = v.Price
		variant.Public = v.Public
		variant.ImagePath = v.ImagePath
		variant.CreatedAt = time.Now().UnixMicro()
		variants = append(variants, variant)
	}
	data.Variants = variants
}

func DomainToDTO_UpsertItemAndVariants(resp *dto.ItemResponse, result *domain.ItemCategory, req *dto.ItemAndVariantsUpsertRequest) {
	for _, item := range result.Items {
		if item.UUID != req.UUID {
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

		variants := []dto.VariantResponse{}
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
	}
}
