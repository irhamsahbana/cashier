package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) UpdateSpace(c context.Context, branchId, id string, req *domain.SpaceUpdateRequest) (*domain.SpaceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if req.Number < 1 {
		return nil, http.StatusUnprocessableEntity, errors.New("number is required and must greather than 0")
	}

	data := domain.Space{
		Number:      req.Number,
		Occupied:    req.Occupied,
		Description: req.Description,
	}

	result, code, err := u.spaceGroupRepo.UpdateSpace(ctx, branchId, id, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.SpaceResponse
	for _, space := range result.Spaces {
		if space.UUID == id {
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

	return &resp, http.StatusOK, nil
}
