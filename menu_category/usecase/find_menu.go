package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *menuCategoryUsecase) FindMenu(c context.Context, id string, withTrashed bool) (*domain.MenuFindResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, err := u.menuCategoryRepo.FindMenu(ctx, id, withTrashed)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	menu := result.Menus[0]
	var resp domain.MenuFindResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public	= menu.Public
	resp.CreatedAt = time.UnixMicro(result.CreatedAt)
	if menu.UpdatedAt != nil {
		menuUpdatedAt := time.UnixMicro(*menu.UpdatedAt)
		resp.UpdatedAt = &menuUpdatedAt
	}
	if menu.DeletedAt != nil {
		menuDeletedAt := time.UnixMicro(*menu.DeletedAt)
		resp.DeletedAt = &menuDeletedAt
	}

	return &resp, http.StatusOK, nil
}