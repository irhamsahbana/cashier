package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *employeeShiftUsecase) ClockIn(ctx context.Context, branchId string, req *domain.EmployeeShiftClockInRequest) (*domain.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if err := validateClockInRequest(req); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	startTime, _ := time.Parse(time.RFC3339Nano, req.StartTime)

	var data domain.EmployeeShiftClockInData
	data.UUID = req.UUID
	data.SupportingUUID = req.SupportingUUID
	data.StartCash = req.StartCash
	data.StartTime = startTime.UnixMicro()
	data.UserUUID = req.UserUUID

	result, code, err := u.employeeShiftRepo.ClockIn(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.EmployeeShiftResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.UserUUID = result.UserUUID
	resp.StartTime = time.UnixMicro(result.StartTime).UTC()
	resp.StartCash = result.StartCash

	var supporters []domain.EmployeeShiftSupporterResponse
	for _, supporter := range result.Supporters {
		var s domain.EmployeeShiftSupporterResponse
		s.UUID = supporter.UUID
		s.StartTime = time.UnixMicro(supporter.StartTime).UTC()
		s.CreatedAt = time.UnixMicro(supporter.CreatedAt).UTC()
		if supporter.EndTime != nil {
			endTime := time.UnixMicro(*supporter.EndTime).UTC()
			s.EndTime = &endTime
		}
		if supporter.UpdatedAt != nil {
			updatedAt := time.UnixMicro(*supporter.UpdatedAt).UTC()
			s.UpdatedAt = &updatedAt
		}
		if supporter.DeletedAt != nil {
			deletedAt := time.UnixMicro(*supporter.DeletedAt).UTC()
			s.DeletedAt = &deletedAt
		}

		supporters = append(supporters, s)
	}
	resp.Supporters = supporters

	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		updatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &updatedAt
	}
	if result.DeletedAt != nil {
		deletedAt := time.UnixMicro(*result.DeletedAt).UTC()
		resp.DeletedAt = &deletedAt
	}

	return &resp, http.StatusOK, nil
}

func validateClockInRequest(req *domain.EmployeeShiftClockInRequest) error {
	if err := validator.IsUUID(req.UUID); err != nil {
		return errors.New("invalid uuid field")
	}

	if err := validator.IsUUID(req.UserUUID); err != nil {
		return errors.New("invalid user_uuid field")
	}

	if req.SupportingUUID != nil {
		if err := validator.IsUUID(*req.SupportingUUID); err != nil {
			return errors.New("supporting_uuid field is not valid")
		}
	}

	if req.StartCash != nil && req.SupportingUUID != nil {
		return errors.New("start_cash and supporting_uuid field cannot be set at the same time")
	}

	if req.SupportingUUID == nil && req.StartCash == nil {
		return errors.New("start_cash field is required if supporting_uuid is null")
	}

	_, err := time.Parse(time.RFC3339Nano, req.StartTime)
	if err != nil {
		return err
	}

	return nil
}
