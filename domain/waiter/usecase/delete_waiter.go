package usecase

import (
	"context"
	"lucy/cashier/lib/dto"
	"net/http"
	"time"
)

func (u *waiterUsecase) DeleteWaiter(c context.Context, id string) (*dto.WaiterResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.waiterRepo.DeleteWaiter(ctx, id)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusOK, nil
	}

	var resp dto.WaiterResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.LastActive != nil {
		respLastActive := time.UnixMicro(*result.LastActive).UTC()
		resp.LastActive = &respLastActive
	}
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
