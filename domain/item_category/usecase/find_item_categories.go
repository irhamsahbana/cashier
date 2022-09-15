package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) FindItemCategories(c context.Context, branchId string) ([]domain.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemCategories(ctx, branchId)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.ItemCategoryResponse

	// mc -> item category, m -> item
	for _, mc := range result {
		var data domain.ItemCategoryResponse
		var items []domain.ItemResponse
		var modifierGroups []domain.ModifierGroupResponse

		data.UUID = mc.UUID
		data.BranchUUID = mc.BranchUUID
		data.Name = mc.Name
		data.CreatedAt = time.UnixMicro(mc.CreatedAt).UTC()
		if mc.UpdatedAt != nil {
			dataUpdatedAt := time.UnixMicro(*mc.UpdatedAt).UTC()
			data.UpdatedAt = &dataUpdatedAt
		}

		if len(mc.Items) > 0 {
			for _, m := range mc.Items {
				var dataItem domain.ItemResponse

				dataItem.UUID = m.UUID
				dataItem.Name = m.Name
				dataItem.Price = m.Price
				dataItem.Description = m.Description
				dataItem.Label = m.Label
				dataItem.Public = m.Public
				dataItem.ImagePath = m.ImagePath

				for _, v := range m.Variants {
					var dataVariant domain.VariantResponse

					dataVariant.UUID = v.UUID
					dataVariant.Label = v.Label
					dataVariant.Price = v.Price
					dataVariant.Public = v.Public
					dataVariant.ImagePath = v.ImagePath
					dataItem.Variants = append(dataItem.Variants, dataVariant)
				}

				if len(dataItem.Variants) == 0 {
					dataItem.Variants = make([]domain.VariantResponse, 0)
				}

				dataItem.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
				if m.UpdatedAt != nil {
					dataUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
					dataItem.UpdatedAt = &dataUpdatedAt
				}

				items = append(items, dataItem)
			}

			if len(items) == 0 {
				items = make([]domain.ItemResponse, 0)
			}
		}

		if len(mc.ModifierGroups) > 0 {
			for _, mg := range mc.ModifierGroups {
				var dataModifierGroup domain.ModifierGroupResponse

				dataModifierGroup.UUID = mg.UUID
				dataModifierGroup.Name = mg.Name
				dataModifierGroup.Quantity = mg.Quantity
				dataModifierGroup.Condition = mg.Condition
				dataModifierGroup.Single = mg.Single
				dataModifierGroup.Required = mg.Required

				if len(mg.Modifiers) > 0 {
					for _, m := range mg.Modifiers {
						var dataModifier domain.ModifierResponse

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

					if len(dataModifierGroup.Modifiers) == 0 {
						dataModifierGroup.Modifiers = make([]domain.ModifierResponse, 0)
					}
				}

				modifierGroups = append(modifierGroups, dataModifierGroup)
			}

			data.ModifierGroups = modifierGroups

			if len(data.ModifierGroups) == 0 {
				data.ModifierGroups = make([]domain.ModifierGroupResponse, 0)
			}
		}

		data.Items = items

		if len(data.Items) == 0 {
			data.Items = make([]domain.ItemResponse, 0)
		}

		resp = append(resp, data)
	}

	if len(resp) == 0 {
		resp = make([]domain.ItemCategoryResponse, 0)
	}

	return resp, http.StatusOK, nil
}
