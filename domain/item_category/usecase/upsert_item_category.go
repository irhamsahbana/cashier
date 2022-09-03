package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *itemCategoryUsecase) UpsertItemCategory(c context.Context, branchId string, req *domain.ItemCategoryUpsertRequest) (*domain.ItemCategoryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var itemcategory domain.ItemCategory
	itemcategory.UUID = req.UUID
	itemcategory.Name = req.Name
	itemcategory.CreatedAt = createdAt.UTC().UnixMicro()

	result, code, err := u.itemCategoryRepo.UpsertItemCategory(ctx, branchId, &itemcategory)
	if err != nil {
		return nil, code, err
	}

	var resp domain.ItemCategoryResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}

	return &resp, http.StatusOK, nil
}
