package usecase

import (
	"context"
	"errors"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *employeeShiftUsecase) History(c context.Context, branchId string, limit, page int) ([]dto.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, code, err := u.employeeShiftRepo.History(ctx, branchId, limit, page)
	if err != nil {
		return nil, code, err
	}

	resp := DomainToDTO_History(result)

	for i, r := range resp {
		shiftIds := []string{r.UUID}

		for _, s := range r.Supporters {
			shiftIds = append(shiftIds, s.UUID)
		}
		fmt.Println("shift ids => ", shiftIds)
		summary, code, err := u.employeeShiftRepo.Summary(ctx, branchId, shiftIds)
		if err != nil {
			return nil, code, errors.New(fmt.Sprintf("failed to get summary for shift %s => %s", r.UUID, err.Error()))
		}

		if summary != nil {
			resp[i].Summary = *summary
		}
	}

	return resp, code, err
}

func DomainToDTO_History(result []domain.EmployeeShift) []dto.EmployeeShiftResponse {
	resp := []dto.EmployeeShiftResponse{}
	for _, r := range result {
		var e dto.EmployeeShiftResponse

		e.UUID = r.UUID
		e.BranchUUID = r.BranchUUID
		e.UserUUID = r.UserUUID
		e.StartTime = time.UnixMicro(r.StartTime).UTC()
		e.StartCash = r.StartCash
		if r.EndTime != nil {
			endTime := time.UnixMicro(*r.EndTime).UTC()
			e.EndTime = &endTime
		}
		e.EndCash = r.EndCash
		e.CreatedAt = time.UnixMicro(r.CreatedAt).UTC()
		if r.UpdatedAt != nil {
			updatedAt := time.UnixMicro(*r.UpdatedAt).UTC()
			e.UpdatedAt = &updatedAt
		}
		if r.DeletedAt != nil {
			deletedAt := time.UnixMicro(*r.DeletedAt).UTC()
			e.DeletedAt = &deletedAt
		}

		// supporters
		supporters := []dto.EmployeeShiftSupporterResponse{}
		for _, s := range r.Supporters {
			var supporter dto.EmployeeShiftSupporterResponse

			supporter.UUID = s.UUID
			supporter.UserUUID = s.UserUUID
			supporter.StartTime = time.UnixMicro(s.StartTime).UTC()
			if s.EndTime != nil {
				endTime := time.UnixMicro(*s.EndTime).UTC()
				supporter.EndTime = &endTime
			}
			supporter.CreatedAt = time.UnixMicro(s.CreatedAt).UTC()
			if s.UpdatedAt != nil {
				updatedAt := time.UnixMicro(*s.UpdatedAt).UTC()
				supporter.UpdatedAt = &updatedAt
			}
			if s.DeletedAt != nil {
				deletedAt := time.UnixMicro(*s.DeletedAt).UTC()
				supporter.DeletedAt = &deletedAt
			}

			supporters = append(supporters, supporter)
		}
		e.Supporters = supporters

		// cash entries
		cashEntries := []dto.CashEntryResponse{}
		for _, c := range r.CashEntries {
			var cashEntry dto.CashEntryResponse

			cashEntry.Username = c.Username
			cashEntry.Description = c.Description
			cashEntry.Expense = c.Expense
			cashEntry.Value = c.Value

			cashEntries = append(cashEntries, cashEntry)
		}
		e.CashEntries = cashEntries

		// summary
		summary := dto.EmployeeShiftSummaryResponse{}
		summary.TotalRefunds = 0
		summary.Orders = []dto.EmployeeShiftSummaryOrder{}
		summary.Payments = []dto.EmployeeShiftSummaryPayment{
			{
				UUID:  "981fddcb-8e10-42ba-a77a-850ae0169c56",
				Qty:   12,
				Total: 120000,
			},
		}
		e.Summary = summary

		resp = append(resp, e)
	}

	return resp
}
