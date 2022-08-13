package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
)

func (u *menuCategoryUsecase) FindMenuCategories(c context.Context, withTrashed bool) ([]domain.MenuCategoryFindAllResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.menuCategoryRepo.FindMenuCategories(ctx, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.MenuCategoryFindAllResponse

	// mc -> menu category, m -> menu
	for _, mc := range result {
		var data domain.MenuCategoryFindAllResponse
		var menus []domain.MenuFindAllResponse

		data.UUID = mc.UUID
		data.Name = mc.Name
		data.CreatedAt = time.UnixMicro(mc.CreatedAt)
		if mc.UpdatedAt != nil {
			dataUpdatedAt := time.UnixMicro(*mc.UpdatedAt)
			data.UpdatedAt = &dataUpdatedAt
		}
		if mc.DeletedAt != nil {
			dataDeletedAt := time.UnixMicro(*mc.DeletedAt)
			data.DeletedAt = &dataDeletedAt
		}

		if len(mc.Menus) > 0 {
			for _, m := range mc.Menus {
				var dataMenu domain.MenuFindAllResponse

				dataMenu.UUID = m.UUID
				dataMenu.MainUUID = m.MainUUID
				dataMenu.Name = m.Name
				dataMenu.Price = m.Price
				dataMenu.Description = m.Description
				dataMenu.Label = m.Label
				dataMenu.Public = m.Public
				dataMenu.ImageUrl = m.ImageUrl
				dataMenu.CreatedAt = time.UnixMicro(m.CreatedAt).UTC()
				if m.UpdatedAt != nil {
					dataUpdatedAt := time.UnixMicro(*m.UpdatedAt).UTC()
					dataMenu.UpdatedAt = &dataUpdatedAt
				}
				if m.DeletedAt != nil {
					dataDeletedAt := time.UnixMicro(*m.DeletedAt).UTC()
					dataMenu.DeletedAt = &dataDeletedAt
				}

				menus = append(menus, dataMenu)
			}
		}

		data.Menus = menus

		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}