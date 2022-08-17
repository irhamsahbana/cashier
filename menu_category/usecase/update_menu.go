package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *menuCategoryUsecase) UpdateMenu (c context.Context, id string, req *domain.MenuUpdateRequest) (*domain.MenuResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if req.Name == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("name is required")
	}

	if req.Label == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("label is required")
	}

	data := domain.Menu{
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
		Label: req.Label,
		Public: req.Public,
	}

	result, code, err := u.menuCategoryRepo.UpdateMenu(ctx, id, &data)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	menu := result.Menus[0]

	var resp domain.MenuResponse
	resp.UUID = id
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.ImageUrl = menu.ImageUrl
	resp.CreatedAt =time.UnixMicro(menu.CreatedAt).UTC()
	if menu.UpdatedAt != nil {
		menuUpdatedAt := time.UnixMicro(*menu.UpdatedAt).UTC()
		resp.UpdatedAt = &menuUpdatedAt
	}
	if menu.DeletedAt != nil {
		menuDeletedAt := time.UnixMicro(*menu.DeletedAt).UTC()
		resp.DeletedAt = &menuDeletedAt
	}

	return &resp, http.StatusOK, nil
}