package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *itemCategoryUsecase) UpsertItemCategoryAndModifiers(c context.Context, branchId string, req *domain.ItemCategoryUpsertRequest) (*domain.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.ItemCategory
	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.Name = req.Name

	if req.ModifierGroups != nil {
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
			modifierGroupItem.Condition = modifierGroup.Condition
			modifierGroupItem.Single = modifierGroup.Single
			modifierGroupItem.Required = modifierGroup.Required

			if modifierGroup.Condition != nil {
				modifierGroupItem.Condition = modifierGroup.Condition

				if err := validateCondition(*modifierGroup.Condition); err != nil {
					return nil, http.StatusUnprocessableEntity, err
				}
			}

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
	}

	result, code, err := u.itemCategoryRepo.UpsertItemCategoryAndModifiers(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.ItemCategoryResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.ModifierGroups != nil {
		var modifierGroups []domain.ModifierGroupResponse

		for _, modifierGroup := range result.ModifierGroups {
			var modifierGroupItem domain.ModifierGroupResponse

			modifierGroupItem.UUID = modifierGroup.UUID
			modifierGroupItem.Name = modifierGroup.Name
			modifierGroupItem.Quantity = modifierGroup.Quantity
			modifierGroupItem.Condition = modifierGroup.Condition
			modifierGroupItem.Single = modifierGroup.Single
			modifierGroupItem.Required = modifierGroup.Required

			if modifierGroup.Modifiers != nil {
				var modifiers []domain.ModifierResponse

				for _, modifier := range modifierGroup.Modifiers {
					var modifierItem domain.ModifierResponse

					modifierItem.UUID = modifier.UUID
					modifierItem.Name = modifier.Name
					modifierItem.Price = modifier.Price
					modifierItem.CreatedAt = time.UnixMicro(modifier.CreatedAt).UTC()

					modifiers = append(modifiers, modifierItem)
				}

				modifierGroupItem.Modifiers = modifiers
			}

			modifierGroups = append(modifierGroups, modifierGroupItem)
		}

		resp.ModifierGroups = modifierGroups
	}

	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}

	return &resp, http.StatusOK, nil
}

func validateCondition(condition domain.ModifierGroupCondition) error {
	switch condition {
	case domain.ModifierGroupCondition_MIN,
		domain.ModifierGroupCondition_MAX,
		domain.ModifierGroupCondition_EQUAL:
		return nil
	default:
		return errors.New(fmt.Sprintf("invalid condition %s, choose one: %s, %s, %s", condition, domain.ModifierGroupCondition_MIN, domain.ModifierGroupCondition_MAX, domain.ModifierGroupCondition_EQUAL))
	}
}
