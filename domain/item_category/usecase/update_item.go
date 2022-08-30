package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *itemCategoryUsecase) UpdateItem(c context.Context, id string, req *domain.ItemUpdateRequest) (*domain.ItemResponse, int, error) {
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

	result, code, err := u.itemCategoryRepo.UpdateItem(ctx, id, &data)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	item := result.Items[0]

	var resp domain.ItemResponse
	resp.UUID = id
	resp.Name = item.Name
	resp.Price = item.Price
	resp.Description = item.Description
	resp.Label = item.Label
	resp.Public = item.Public
	resp.ImageUrl = item.ImageUrl
	resp.CreatedAt = time.UnixMicro(item.CreatedAt).UTC()
	if item.UpdatedAt != nil {
		itemUpdatedAt := time.UnixMicro(*item.UpdatedAt).UTC()
		resp.UpdatedAt = &itemUpdatedAt
	}
	if item.DeletedAt != nil {
		itemDeletedAt := time.UnixMicro(*item.DeletedAt).UTC()
		resp.DeletedAt = &itemDeletedAt
	}

	return &resp, http.StatusOK, nil
}
