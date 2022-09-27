package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/validator"
)

func (u *itemCategoryUsecase) UpsertItemCategoryAndModifiers(c context.Context, branchId string, req *dto.ItemCategoryUpsertRequest) (*dto.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.ItemCategory
	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.Name = req.Name

	var modifierGroups []domain.ModifierGroup
	for _, modifierGroup := range req.ModifierGroups {
		var modifierGroupItem domain.ModifierGroup

		if err := validator.IsUUID(modifierGroup.UUID); err != nil {
			return nil, http.StatusUnprocessableEntity, err
		}

		if modifierGroup.Name == "" {
			return nil, http.StatusUnprocessableEntity, errors.New("modifier group name is required")
		}

		modifierGroupItem.UUID = modifierGroup.UUID
		modifierGroupItem.Name = modifierGroup.Name
		modifierGroupItem.Quantity = modifierGroup.Quantity
		modifierGroupItem.Single = modifierGroup.Single
		modifierGroupItem.Required = modifierGroup.Required

		if modifierGroup.Modifiers != nil {
			var modifiers []domain.Modifier

			for _, modifier := range modifierGroup.Modifiers {
				var modifierItem domain.Modifier

				modifierItem.UUID = modifier.UUID
				modifierItem.Name = modifier.Name
				modifierItem.Price = modifier.Price
				modifierItem.CreatedAt = time.Now().UnixMicro()

				if err := validator.IsUUID(modifier.UUID); err != nil {
					return nil, http.StatusUnprocessableEntity, err
				}

				if modifier.Name == "" {
					return nil, http.StatusUnprocessableEntity, errors.New("modifier name is required")
				}

				modifiers = append(modifiers, modifierItem)
			}

			modifierGroupItem.Modifiers = modifiers
		}

		modifierGroups = append(modifierGroups, modifierGroupItem)
	}
	data.ModifierGroups = modifierGroups

	result, code, err := u.itemCategoryRepo.UpsertItemCategoryAndModifiers(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.ItemCategoryResponse
	itemCategoryDomainToDTO_UpsertItemCategoryAndModifiers(&resp, result)

	return &resp, http.StatusOK, nil
}

func itemCategoryDomainToDTO_UpsertItemCategoryAndModifiers(resp *dto.ItemCategoryResponse, result *domain.ItemCategory) {
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()

	var modifierGroupsResp []dto.ModifierGroupResponse
	for _, modifierGroup := range result.ModifierGroups {
		var modifierGroupItem dto.ModifierGroupResponse

		modifierGroupItem.UUID = modifierGroup.UUID
		modifierGroupItem.Name = modifierGroup.Name
		modifierGroupItem.Quantity = modifierGroup.Quantity
		modifierGroupItem.Single = modifierGroup.Single
		modifierGroupItem.Required = modifierGroup.Required

		var modifiers []dto.ModifierResponse
		for _, modifier := range modifierGroup.Modifiers {
			var modifierItem dto.ModifierResponse

			modifierItem.UUID = modifier.UUID
			modifierItem.Name = modifier.Name
			modifierItem.Price = modifier.Price
			modifierItem.CreatedAt = time.UnixMicro(modifier.CreatedAt).UTC()

			modifiers = append(modifiers, modifierItem)
		}
		modifierGroupItem.Modifiers = modifiers

		if len(modifierGroupItem.Modifiers) == 0 {
			modifierGroupItem.Modifiers = make([]dto.ModifierResponse, 0)
		}

		modifierGroupsResp = append(modifierGroupsResp, modifierGroupItem)
	}
	resp.ModifierGroups = modifierGroupsResp

	if len(result.ModifierGroups) == 0 {
		resp.ModifierGroups = []dto.ModifierGroupResponse{}
	}

	if len(result.Items) == 0 {
		resp.Items = make([]dto.ItemResponse, 0)
	}

	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
}
