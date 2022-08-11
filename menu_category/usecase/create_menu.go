package usecase

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
)

func (u *menuCategoryUsecase) CreateMenu(c context.Context, menuCategoryId string, req *domain.MenuCreateRequest) (*domain.MenuResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if  err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	data := domain.Menu{
		UUID: req.UUID,
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
		Label: req.Label,
		Public: req.Public,
		CreatedAt: createdAt.UnixMicro(),
	}

	result, code, err := u.menuCategoryRepo.InsertMenu(ctx, menuCategoryId, &data)
	if err != nil {
		return nil, code, err
	}

	menu := result.Menus[0]

	var resp domain.MenuResponse
	resp.UUID = menu.UUID
	resp.Name = menu.Name
	resp.Price = menu.Price
	resp.Description = menu.Description
	resp.Label = menu.Label
	resp.Public = menu.Public
	resp.CreatedAt = time.UnixMicro(menu.CreatedAt)

	return &resp, http.StatusCreated, nil
}

