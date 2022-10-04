package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/validator"
	"net/http"
	"net/mail"
	"time"
)

func (u *userUsecase) UpsertCustomer(c context.Context, req *dto.CustomerUpserRequest) (*dto.CustomerResponse, int, error) {
	err := validateUpsertCustomerRequest(req)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	// find user role with name customer
	userRole, code, err := u.userRoleRepo.FindUserRoleByName(c, "Customer", false)
	if err != nil {
		return nil, code, err
	}

	var data domain.User
	data.RoleUUID = userRole.UUID
	userDTOtoDomain_UpsertCustomer(&data, req)
	result, code, err := u.userRepo.UpsertUser(c, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.CustomerResponse
	userDomainToDTO_UpsertCustomer(&resp, result)

	return &resp, code, nil
}

func validateUpsertCustomerRequest(req *dto.CustomerUpserRequest) error {
	if err := validator.IsUUID(req.UUID); err != nil {
		return errors.New("invalid uuid")
	}

	if err := validator.IsUUID(req.BranchUUID); err != nil {
		return errors.New("invalid branch uuid")
	}

	if req.Name == "" || len(req.Name) > 100 || len(req.Name) < 3 {
		return errors.New("name is required and max length is 100 and min length is 3")
	}

	if _, err := time.Parse(time.RFC3339Nano, req.Dob); err != nil {
		return errors.New("invalid date of birth")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("invalid email")
	}

	if _, err := time.Parse(time.RFC3339Nano, req.CreatedAt); err != nil {
		return errors.New("invalid created at")
	}

	return nil
}

func userDTOtoDomain_UpsertCustomer(data *domain.User, req *dto.CustomerUpserRequest) {
	data.UUID = req.UUID
	data.BranchUUID = req.BranchUUID
	data.Name = req.Name
	data.Email = req.Email
	data.Phone = &req.Phone
	data.Address = &req.Address
	createdAt, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)
	data.CreatedAt = createdAt.UnixMicro()
}

func userDomainToDTO_UpsertCustomer(resp *dto.CustomerResponse, data *domain.User) {
	resp.UUID = data.UUID
	resp.BranchUUID = data.BranchUUID
	resp.Name = data.Name
	resp.Email = data.Email
	resp.Phone = *data.Phone
	resp.Address = *data.Address
	resp.CreatedAt = time.UnixMicro(data.CreatedAt)
	if data.UpdatedAt != nil {
		updatedAt := time.UnixMicro(*data.UpdatedAt)
		resp.UpdatedAt = &updatedAt
	}
	if data.DeletedAt != nil {
		deletedAt := time.UnixMicro(*data.DeletedAt)
		resp.DeletedAt = &deletedAt
	}
}
