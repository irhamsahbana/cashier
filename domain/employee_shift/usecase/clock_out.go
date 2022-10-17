package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"net/http"
	"time"
)

func (u *employeeShiftUsecase) ClockOut(ctx context.Context, branchId string, req *dto.EmployeeShiftClockOutRequest) (*dto.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	endTime, _ := time.Parse(time.RFC3339Nano, req.EndTime)

	var data domain.EmployeeShiftClockOutData
	data.UUID = req.UUID
	data.EndCash = req.EndCash
	data.EndTime = endTime.UnixMicro()

	result, code, err := u.employeeShiftRepo.ClockOut(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.EmployeeShiftResponse
	DomainToDTO_ClockOut(&resp, result)

	return &resp, http.StatusOK, nil
}

func DomainToDTO_ClockOut(resp *dto.EmployeeShiftResponse, result *domain.EmployeeShift) {
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
