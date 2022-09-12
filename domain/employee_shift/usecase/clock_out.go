package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *employeeShiftUsecase) ClockOut(ctx context.Context, branchId string, req *domain.EmployeeShiftClockOutRequest) (*domain.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if err := validateClockOutRequest(req); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	endTime, _ := time.Parse(time.RFC3339Nano, req.EndTime)

	var data domain.EmployeeShiftClockOutData
	data.UUID = req.UUID
	data.EndCash = req.EndCash
	data.EndTime = endTime.UnixMicro()

	result, code, err := u.employeeShiftRepo.ClockOut(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.EmployeeShiftResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.UserUUID = result.UserUUID
	resp.StartTime = time.UnixMicro(result.StartTime).UTC()
	resp.StartCash = result.StartCash
	if result.EndTime != nil {
		endTime := time.UnixMicro(*result.EndTime).UTC()
		resp.EndTime = &endTime
	}
	if result.EndCash != nil {
		resp.EndCash = result.EndCash
	}

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

func validateClockOutRequest(req *domain.EmployeeShiftClockOutRequest) error {
	if err := validator.IsUUID(req.UUID); err != nil {
		return errors.New("invalid uuid field")
	}

	if _, err := time.Parse(time.RFC3339Nano, req.EndTime); err != nil {
		return errors.New("invalid end time field")
	}

	if req.EndCash != nil {
		if *req.EndCash < 0 {
			return errors.New("invalid end cash field")
		}
	}

	return nil
}
