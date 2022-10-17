package usecase

import (
	"context"
	"lucy/cashier/dto"
)

func (u *employeeShiftUsecase) InsertEntryCash(c context.Context, branchId, shiftId string, req *dto.CashEntryInsertRequest) ([]dto.CashEntryResponse, int, error) {
	return nil, 200, nil
}
