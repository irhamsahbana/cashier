package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) DeleteSpaceGroup(c context.Context, branchId, id string) (*domain.SpaceGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.spaceGroupRepo.DeleteSpaceGroup(ctx, branchId, id)
	if err != nil {
		if code == http.StatusNotFound {
			return nil, http.StatusOK, nil
		}
		return nil, code, err
	}

	var resp domain.SpaceGroupResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.Code = result.Code
	resp.Shape = result.Shape
	resp.Pax = result.Pax
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if result.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*result.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}
