package usecase

import (
	"context"
	"lucy/cashier/dto"
)

func (u *orderUsecase) InsertInvoiceRefund(c context.Context, branchId string, req *dto.RefundInsertRequest) {
	// ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	// defer cancel()
	// u.orderRepo.InsertInvoiceRefund(ctx, branchId, req)
}
