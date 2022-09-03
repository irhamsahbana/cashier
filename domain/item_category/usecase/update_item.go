package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *itemCategoryUsecase) UpdateItem(c context.Context, branchId, id string, req *domain.ItemUpdateRequest) (*domain.ItemResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if req.Name == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("name is required")
	}

	if req.Label == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("label is required")
	}

	data := domain.Item{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Label:       req.Label,
		Public:      req.Public,
	}

	result, code, err := u.itemCategoryRepo.UpdateItem(ctx, branchId, id, &data)
	if err != nil {
		return nil, code, err
	}
	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	var resp domain.ItemResponse
	for _, item := range result.Items {
		if item.UUID != id {
			continue
		}

		resp.UUID = item.UUID
		resp.MainUUID = item.MainUUID
		resp.Name = item.Name
		resp.Price = item.Price
		resp.Description = item.Description
		resp.Label = item.Label
		resp.Public = item.Public
		resp.ImageUrl = item.ImageUrl
		resp.CreatedAt = time.UnixMicro(item.CreatedAt).UTC()
		if item.UpdatedAt != nil {
			respUpdatedAt := time.UnixMicro(*item.UpdatedAt).UTC()
			resp.UpdatedAt = &respUpdatedAt
		}
		if item.DeletedAt != nil {
			respDeletedAt := time.UnixMicro(*item.DeletedAt).UTC()
			resp.DeletedAt = &respDeletedAt
		}
	}

	return &resp, http.StatusOK, nil
}
