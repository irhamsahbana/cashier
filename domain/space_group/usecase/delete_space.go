package usecase

import (
	"context"
	"lucy/cashier/domain"
	"time"
)

func (u *spaceGroupUsecase) DeleteSpace(c context.Context, branchId, id string) (*domain.SpaceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.spaceGroupRepo.DeleteSpace(ctx, branchId, id)
	if err != nil {
		return nil, code, err
	}

	var resp domain.SpaceResponse
	for _, space := range result.Spaces {
		if space.UUID == id {
			resp.UUID = space.UUID
			resp.Number = space.Number
			resp.Description = space.Description
			resp.CreatedAt = time.UnixMicro(space.CreatedAt).UTC()
			if space.UpdatedAt != nil {
				respUpdatedAt := time.UnixMicro(*space.UpdatedAt).UTC()
				resp.UpdatedAt = &respUpdatedAt
			}
			if space.DeletedAt != nil {
				respDeletedAt := time.UnixMicro(*space.DeletedAt).UTC()
				resp.DeletedAt = &respDeletedAt
			}
		}
	}

	return &resp, code, nil
}
