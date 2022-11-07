package usecase

import (
	"context"
	"errors"
	"lucy/cashier/dto"
	"net/http"
	"time"
)

func (u *orderUsecase) DeleteActiveOrder(c context.Context, branchId, orderId string, req *dto.OrderGroupDeleteRequest) (*dto.OrderGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if req.CancelReason == "" {
		return nil, http.StatusUnprocessableEntity, errors.New("reason is required")
	}

	deletedAt, err := time.Parse(time.RFC3339Nano, req.DeletedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, errors.New("deleted_at is invalid")
	}

	var resp dto.OrderGroupResponse
	result, code, err := u.orderRepo.DeleteActiveOrder(ctx, branchId, orderId, req.CancelReason, deletedAt.UnixMicro())
	if err != nil {
		if code == http.StatusNotFound {
			return nil, http.StatusOK, err
		}
		return &resp, code, err
	}

	DomainToDTO_UpsertActiveOrder(&resp, result)

	return &resp, code, nil
}
