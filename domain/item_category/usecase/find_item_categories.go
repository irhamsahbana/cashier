package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/dto"
)

func (u *itemCategoryUsecase) FindItemCategories(c context.Context, branchId string) ([]dto.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.itemCategoryRepo.FindItemCategories(ctx, branchId)
	if err != nil {
		return nil, code, err
	}

	var resp []dto.ItemCategoryResponse
	resp = itemCategoryDomainToDTO_FindItemCategories(resp, result)

	return resp, http.StatusOK, nil
}

func itemCategoryDomainToDTO_FindItemCategories(resp []dto.ItemCategoryResponse, result []domain.ItemCategory) []dto.ItemCategoryResponse {
	for _, mc := range result {
		var data dto.ItemCategoryResponse
		var items []dto.ItemResponse
		var modifierGroups []dto.ModifierGroupResponse

		data.UUID = mc.UUID
		data.BranchUUID = mc.BranchUUID
		data.Name = mc.Name
		data.CreatedAt = time.UnixMicro(mc.CreatedAt).UTC()
		if mc.UpdatedAt != nil {
			respUpdatedAt := time.UnixMicro(*mc.UpdatedAt).UTC()
			data.UpdatedAt = &respUpdatedAt
		}

		// items
		for _, m := range mc.Items {
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

			if len(dataItem.Variants) == 0 {
				dataItem.Variants = make([]dto.VariantResponse, 0)
			}

			dataItem.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
			if m.UpdatedAt != nil {
				respUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
				dataItem.UpdatedAt = &respUpdatedAt
			}

			items = append(items, dataItem)
		}
		// items, when empty make it empty array
		data.Items = items
		if len(data.Items) == 0 {
			data.Items = make([]dto.ItemResponse, 0)
		}

		// modifier groups
		for _, mg := range mc.ModifierGroups {
			var dataModifierGroup dto.ModifierGroupResponse
			dataModifierGroup.UUID = mg.UUID
			dataModifierGroup.Name = mg.Name
			dataModifierGroup.MaxQty = mg.MaxQty
			dataModifierGroup.MinQty = mg.MinQty

			// modifiers
			for _, m := range mg.Modifiers {
				var dataModifier dto.ModifierResponse

				dataModifier.UUID = m.UUID
				dataModifier.Name = m.Name
				dataModifier.Price = m.Price
				dataModifier.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
				if m.UpdatedAt != nil {
					respUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
					dataModifier.UpdatedAt = &respUpdatedAt
				}

				dataModifierGroup.Modifiers = append(dataModifierGroup.Modifiers, dataModifier)
			}
			// modifiers, when empty make it empty array
			if len(dataModifierGroup.Modifiers) == 0 {
				dataModifierGroup.Modifiers = make([]dto.ModifierResponse, 0)
			}

			dataModifierGroup.CreatedAt = time.UnixMicro(mg.CreatedAt).UTC()
			if mg.UpdatedAt != nil {
				respUpdatedAt := time.UnixMicro(*mg.UpdatedAt).UTC()
				dataModifierGroup.UpdatedAt = &respUpdatedAt
			}

			modifierGroups = append(modifierGroups, dataModifierGroup)
		}
		// modifier groups, when empty make it empty array
		data.ModifierGroups = modifierGroups
		if len(modifierGroups) == 0 {
			data.ModifierGroups = make([]dto.ModifierGroupResponse, 0)
		}

		resp = append(resp, data)
	}

	// if not have item category ( include item & modifier group not found) so make empty array
	if len(resp) == 0 {
		resp = make([]dto.ItemCategoryResponse, 0)
	}

	return resp
}
