package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) CreateSpace(c context.Context, branchId, SpaceGroupId string, req *domain.SpaceCreateRequest) (*domain.SpaceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if req.Number < 1 {
		return nil, http.StatusUnprocessableEntity, errors.New("number is required and must greather than 0")
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	data := domain.Space{
		UUID:        req.UUID,
		Number:      req.Number,
		Occupied:    req.Occupied,
		Description: req.Description,
		CreatedAt:   createdAt.UTC().UnixMicro(),
	}

	result, code, err := u.spaceGroupRepo.InsertSpace(ctx, branchId, SpaceGroupId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.SpaceResponse
	for _, space := range result.Spaces {
		if space.UUID == req.UUID {
			resp = domain.SpaceResponse{
				UUID:        space.UUID,
				Number:      space.Number,
				Occupied:    space.Occupied,
				Description: space.Description,
				CreatedAt:   time.UnixMicro(space.CreatedAt).UTC(),
			}

			if space.UpdatedAt != nil {
				unixMicro := *space.UpdatedAt
				updatedAt := time.UnixMicro(unixMicro).UTC()
				resp.UpdatedAt = &updatedAt
			}

			if space.DeletedAt != nil {
				unixMicro := *space.DeletedAt
				deletedAt := time.UnixMicro(unixMicro).UTC()
				resp.DeletedAt = &deletedAt
			}
		}
	}

	return &resp, http.StatusCreated, nil
}
