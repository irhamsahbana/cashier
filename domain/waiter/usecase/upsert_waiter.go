package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *waiterUsecase) UpsertWaiter(c context.Context, req *domain.WaiterUpsertrequest) (*domain.WaiterResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if  err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if req.Name == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("Name field is required")
	}

	var waiter domain.Waiter
	waiter.UUID = req.UUID
	waiter.Name = req.Name
	waiter.CreatedAt = createdAt.UTC().UnixMicro()

	result, code, err := u.waiterRepo.UpsertWaiter(ctx, &waiter)
	if err != nil {
		return nil, code, err
	}

	var resp domain.WaiterResponse
	resp.UUID = result.UUID
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