package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *employeeShiftUsecase) InsertEntryCash(c context.Context, branchId, shiftId string, req *dto.CashEntryInsertRequest) ([]dto.CashEntryResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var data domain.CashEntry
	DTOtoDomain_insertEntryCash(&data, req)

	result, code, err := u.employeeShiftRepo.InsertEntryCash(ctx, branchId, shiftId, &data)
	if err != nil {
		return nil, code, err
	}

	resp := DomainToDTO_InsertEntryCash(result)

	return resp, code, nil
}

func DTOtoDomain_insertEntryCash(data *domain.CashEntry, req *dto.CashEntryInsertRequest) {
	data.Username = req.Username
	data.Description = req.Description
	data.Expense = req.Expense
	data.Value = req.Value
	createdAt, _ := time.Parse(time.RFC3339, req.CreatedAt)
	data.CreatedAt = createdAt.UnixMicro()
}

func DomainToDTO_InsertEntryCash(result []domain.CashEntry) []dto.CashEntryResponse {
	resp := []dto.CashEntryResponse{}

	for _, r := range result {
		var e dto.CashEntryResponse

		e.Username = r.Username
		e.Description = r.Description
		e.Expense = r.Expense
		e.Value = r.Value
		e.CreatedAt = time.UnixMicro(r.CreatedAt).UTC()

		resp = append(resp, e)
	}

	return resp
}
