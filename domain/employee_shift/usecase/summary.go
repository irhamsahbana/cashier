package usecase

import (
	"context"
	"lucy/cashier/dto"
)

func (u *employeeShiftUsecase) Summary(c context.Context, branchId string, shiftIds []string) dto.EmployeeShiftSummaryResponse {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	summary := dto.EmployeeShiftSummaryResponse{}

	return summary
}
