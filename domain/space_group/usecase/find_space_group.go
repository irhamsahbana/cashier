package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) FindSpaceGroup(c context.Context, branchId, id string, withTrashed bool) (*domain.SpaceGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.spaceGroupRepo.FindSpaceGroup(ctx, branchId, id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp domain.SpaceGroupResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Code = result.Code
	resp.Shape = result.Shape
	resp.Pax = result.Pax
	resp.Floor = result.Floor
	resp.Smooking = result.Smooking
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
