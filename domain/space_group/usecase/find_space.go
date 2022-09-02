package usecase

import (
	"context"
	"lucy/cashier/domain"
	"time"
)

func (u *spaceGroupUsecase) FindSpace(c context.Context, branchId, id string, withTrashed bool) (*domain.SpaceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.spaceGroupRepo.FindSpace(ctx, branchId, id, withTrashed)
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

	return &resp, code, nil
}
