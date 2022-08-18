package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *menuCategoryUsecase) DeleteMenu(c context.Context, id string) (*domain.MenuResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.DeleteMenu(ctx, id)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	menu := result.Menus[0]

	var resp domain.MenuResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt).UTC()
	if menu.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*menu.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if menu.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*menu.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}