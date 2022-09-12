package usecase

import (
	"context"
	"lucy/cashier/domain"
	"time"
)

func (u *employeeShiftUsecase) History(c context.Context, branchId string, limit, offset int) ([]domain.EmployeeShiftResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.employeeShiftRepo.History(ctx, branchId, limit, offset)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.EmployeeShiftResponse
	for _, r := range result {
		var e domain.EmployeeShiftResponse

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

		var supporters []domain.EmployeeShiftSupporterResponse
		for _, s := range r.Supporters {
			var supporter domain.EmployeeShiftSupporterResponse

			supporter.UUID = s.UUID
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

		resp = append(resp, e)
	}

	return resp, code, err
}
