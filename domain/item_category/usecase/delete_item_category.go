package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/dto"
)

func (u *itemCategoryUsecase) DeleteItemCategory(c context.Context, branchId, id string) (*dto.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.DeleteItemCategory(ctx, branchId, id)
	if err != nil {
		if code == http.StatusNotFound {
			return nil, http.StatusOK, nil
		}

		return nil, code, err
	}

	var resp dto.ItemCategoryResponse
	itemCategoryDomainToDTO_DeleteItemCategory(&resp, result)

	return &resp, http.StatusOK, nil
}

func itemCategoryDomainToDTO_DeleteItemCategory(resp *dto.ItemCategoryResponse, result *domain.ItemCategory) {
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}

	// items
	var items []dto.ItemResponse
	if len(result.Items) > 0 {
		for _, m := range result.Items {
			var dataItem dto.ItemResponse

			dataItem.UUID = m.UUID
			dataItem.Name = m.Name
			dataItem.Price = m.Price
			dataItem.Description = m.Description
			dataItem.Label = m.Label
			dataItem.Public = m.Public
			dataItem.ImagePath = m.ImagePath
			for _, v := range m.Variants {
				var dataVariant dto.VariantResponse

				dataVariant.UUID = v.UUID
				dataVariant.Label = v.Label
				dataVariant.Price = v.Price
				dataVariant.Public = v.Public
				dataVariant.ImagePath = v.ImagePath
				dataItem.Variants = append(dataItem.Variants, dataVariant)
			}
			dataItem.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
			if m.UpdatedAt != nil {
				dataUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
				dataItem.UpdatedAt = &dataUpdatedAt
			}

			items = append(items, dataItem)
		}

		resp.Items = items
	}

	// modifier groups
	var modifierGroups []dto.ModifierGroupResponse
	if len(result.ModifierGroups) > 0 {
		for _, mg := range result.ModifierGroups {
			var dataModifierGroup dto.ModifierGroupResponse

			dataModifierGroup.UUID = mg.UUID
			dataModifierGroup.Name = mg.Name

			if len(mg.Modifiers) > 0 {
				for _, m := range mg.Modifiers {
					var dataModifier dto.ModifierResponse

					dataModifier.UUID = m.UUID
					dataModifier.Name = m.Name
					dataModifier.Price = m.Price
					dataModifier.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
					if m.UpdatedAt != nil {
						dataUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
						dataModifier.UpdatedAt = &dataUpdatedAt
					}

					dataModifierGroup.Modifiers = append(dataModifierGroup.Modifiers, dataModifier)
				}
			}

			modifierGroups = append(modifierGroups, dataModifierGroup)
		}

		resp.ModifierGroups = modifierGroups
	}
}
