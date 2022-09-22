package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/dto"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *waiterUsecase) UpsertWaiter(c context.Context, branchId string, req *dto.WaiterUpsertrequest) (*dto.WaiterResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if req.Name == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("Name field is required")
	}

	var data domain.Waiter
	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.Name = req.Name
	data.CreatedAt = createdAt.UTC().UnixMicro()

	result, code, err := u.waiterRepo.UpsertWaiter(ctx, &data)
	if err != nil {
		return nil, code, err
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
