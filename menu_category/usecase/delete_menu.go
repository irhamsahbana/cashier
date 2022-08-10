package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *menuCategoryUsecase) DeleteMenu(c context.Context, id string) (*domain.MenuDeleteResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.DeleteMenu(ctx, id)
	if err != nil {
		return nil, code, err
	}

	menu := result.Menus[0]

	var resp domain.MenuDeleteResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt)
	if menu.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*menu.UpdatedAt)
		resp.UpdatedAt = &respUpdatedAt
	}
	if menu.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*menu.DeletedAt)
		resp.DeletedAt = respDeletedAt
	}

	return &resp, http.StatusOK, nil
}