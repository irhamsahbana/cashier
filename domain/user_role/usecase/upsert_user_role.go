package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *userRoleUsecase) UpsertUserRole(ctx context.Context, req *domain.UserRoleUpsertrequest) (*domain.UserRoleModel, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
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

	var userRole domain.UserRole
	userRole.UUID = req.UUID
	userRole.Name = req.Name
	userRole.Power = req.Power
	userRole.CreatedAt = createdAt.UTC().UnixMicro()

	result, code, err := u.userRoleRepo.UpsertUserRole(ctx, &userRole)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UserRoleModel
	resp.UUID = result.UUID
	resp.Name = result.Name
	resp.Power = result.Power
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
