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

	for _, mc := range result {
		var data domain.MenuCategoryFindAllResponse

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

		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}