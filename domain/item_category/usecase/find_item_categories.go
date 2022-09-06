package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *itemCategoryUsecase) FindItemCategories(c context.Context, branchId string, withTrashed bool) ([]domain.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemCategories(ctx, branchId, withTrashed)
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

				items = append(items, dataItem)
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
				}

				modifierGroups = append(modifierGroups, dataModifierGroup)
			}

			data.ModifierGroups = modifierGroups
		}

		data.Items = items
		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}
