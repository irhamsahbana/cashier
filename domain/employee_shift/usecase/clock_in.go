package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"net/http"
	"time"
)

func (u *employeeShiftUsecase) ClockIn(ctx context.Context, branchId string, req *dto.EmployeeShiftClockInRequest) (*dto.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var data domain.EmployeeShiftClockInData
	DTOtoDomain_ClockIn(&data, req)

	result, code, err := u.employeeShiftRepo.ClockIn(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.EmployeeShiftResponse
	DomainToDTO_ClockIn(&resp, result)

	return &resp, http.StatusOK, nil
}

func DTOtoDomain_ClockIn(data *domain.EmployeeShiftClockInData, req *dto.EmployeeShiftClockInRequest) {
	data.UUID = req.UUID
	data.SupportingUUID = req.SupportingUUID
	data.StartCash = req.StartCash
	startTime, _ := time.Parse(time.RFC3339Nano, req.StartTime)
	data.StartTime = startTime.UnixMicro()
	data.UserUUID = req.UserUUID
}

func DomainToDTO_ClockIn(resp *dto.EmployeeShiftResponse, result *domain.EmployeeShift) {
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.UserUUID = result.UserUUID
	resp.StartTime = time.UnixMicro(result.StartTime).UTC()
	resp.StartCash = result.StartCash

	// supporters
	supporters := []dto.EmployeeShiftSupporterResponse{}
	for _, supporter := range result.Supporters {
		var s dto.EmployeeShiftSupporterResponse
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

	// cash entries
	cashEntries := []dto.CashEntryResponse{}
	for _, cashEntry := range result.CashEntries {
		var c dto.CashEntryResponse
		c.Username = cashEntry.Username
		c.Description = cashEntry.Description
		c.Expense = cashEntry.Expense
		c.Value = cashEntry.Value
		c.CreatedAt = time.UnixMicro(cashEntry.CreatedAt).UTC()

		cashEntries = append(cashEntries, c)
	}
	resp.CashEntries = cashEntries

	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		updatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &updatedAt
	}
	if result.DeletedAt != nil {
		deletedAt := time.UnixMicro(*result.DeletedAt).UTC()
		resp.DeletedAt = &deletedAt
	}
}
