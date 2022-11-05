package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *orderUsecase) InsertRefund(c context.Context, branchId, invoiceId string, req *dto.RefundInsertRequest) (*dto.InvoiceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var data domain.InvoiceRefund
	DTOtoDomain_InsertRefund(&data, req)
	refunds := []domain.OrderRefundData{}
	for _, refund := range req.OrderRefunds {
		var rData domain.OrderRefundData
		rData.OrderGroupUUID = refund.OrderGroupUUID
		rData.OrderUUID = refund.OrderUUID
		rData.Qty = refund.Qty

		refunds = append(refunds, rData)
	}

	result, code, err := u.orderRepo.InsertRefund(ctx, branchId, invoiceId, &data, refunds)
	if err != nil {
		return nil, code, err
	}

	var resp dto.InvoiceResponse

	DomainToDTO_InsertInvoice(&resp, result)

	return &resp, code, nil
}

// DTO to Domain
func DTOtoDomain_InsertRefund(data *domain.InvoiceRefund, req *dto.RefundInsertRequest) {
	createdAt, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)

	data.UUID = req.UUID
	data.Total = req.Total

	data.EmployeeShift.EmployeeShiftUUID = req.EmployeeShift.EmployeeShiftUUID
	data.EmployeeShift.UserUUID = req.EmployeeShift.UserUUID
	data.EmployeeShift.Name = req.EmployeeShift.Name

	data.CreatedAt = createdAt.UnixMicro()
}
