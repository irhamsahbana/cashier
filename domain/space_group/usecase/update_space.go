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
	if code == http.StatusNotFound {
		return nil, http.StatusOK, err
	}

	var resp domain.SpaceResponse
	for _, space := range result.Spaces {
		if space.UUID != id {
			continue
		}

		resp.UUID = space.UUID
		resp.Number = space.Number
		resp.Occupied = space.Occupied
		resp.Description = space.Description
		resp.CreatedAt = time.UnixMicro(space.CreatedAt).UTC()
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

	return &resp, http.StatusOK, nil
}
