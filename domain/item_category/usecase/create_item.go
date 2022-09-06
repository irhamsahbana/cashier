package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *itemCategoryUsecase) CreateItem(c context.Context, branchId, itemCategoryId string, req *domain.ItemCreateRequest) (*domain.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	data := domain.Item{
		UUID:        req.UUID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Public:      req.Public,
		CreatedAt:   createdAt.UnixMicro(),
	}

	result, code, err := u.itemCategoryRepo.InsertItem(ctx, branchId, itemCategoryId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.ItemResponse
	for _, item := range result.Items {
		if item.UUID != req.UUID {
			continue
		}

		resp.UUID = item.UUID
		resp.MainUUID = item.MainUUID
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
	}

	return &resp, http.StatusCreated, nil
}
