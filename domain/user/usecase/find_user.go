package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *userUsecase) FindUser(c context.Context, id string, withTrashed bool) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.userRepo.FindUserBy(ctx, "uuid", id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UserResponse
	resp.UUID = result.UUID
	resp.Name = result.Name
	respCreatedAt :=  time.UnixMicro(result.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt
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
