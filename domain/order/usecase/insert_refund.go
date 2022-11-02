package usecase

import (
	"context"
	"lucy/cashier/dto"
)

func (u *orderUsecase) InsertRefund(ctx context.Context, branchId string, req *dto.RefundInsertRequest) ([]dto.InvoiceRefundResponse, int, error) {
	panic("implement me")
}
