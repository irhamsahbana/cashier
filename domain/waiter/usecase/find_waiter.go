package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *waiterUsecase) FindWaiter(c context.Context, id string, withTrashed bool) (*domain.WaiterResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.waiterRepo.FindWaiter(ctx, id, withTrashed)
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