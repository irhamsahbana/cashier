package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *menuCategoryUsecase) FindMenu(c context.Context, id string, withTrashed bool) (*domain.MenuResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.FindMenu(ctx, id, withTrashed)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	menu := result.Menus[0]
	var resp domain.MenuResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public	= menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt).UTC()
	if menu.UpdatedAt != nil {
		menuUpdatedAt := time.UnixMicro(*menu.UpdatedAt).UTC()
		resp.UpdatedAt = &menuUpdatedAt
	}
	if menu.DeletedAt != nil {
		menuDeletedAt := time.UnixMicro(*menu.DeletedAt).UTC()
		resp.DeletedAt = &menuDeletedAt
	}

	return &resp, code, nil
}