package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *orderUsecase) InsertInvoice(c context.Context, branchId string, req *dto.InvoiceInsertRequest) (*dto.InvoiceResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var data domain.Invoice
	DTOtoDomain_InsertInvoice(&data, req, branchId)

	result, code, err := u.orderRepo.InsertInvoice(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.InvoiceResponse
	DomainToDTO_InsertInvoice(&resp, result)

	return &resp, code, nil
}

func DTOtoDomain_InsertInvoice(data *domain.Invoice, req *dto.InvoiceInsertRequest, branchId string) {
	createdAt, _ := time.Parse(time.RFC3339, req.CreatedAt)

	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.TotalAmount = req.TotalAmount
	data.TotalTax = req.TotalTax
	data.TotalDiscount = req.TotalDiscount
	data.TotalChange = req.TotalChange
	data.TotalTip = req.TotalTip
	data.Note = req.Note
	data.CreatedAt = createdAt.UnixMicro()
	if req.CompletedAt != nil {
		completedAt, _ := time.Parse(time.RFC3339, *req.CompletedAt)
		completedAtUnixMicro := completedAt.UnixMicro()
		data.CompletedAt = &completedAtUnixMicro
	}

	if req.Customer != nil {
		var customer domain.Customer
		customer.Name = req.Customer.Name
		customer.Phone = req.Customer.Phone
		customer.Address = req.Customer.Address
		data.Customer = &customer
	}

	// payments
	data.Payments = []domain.InvoicePayment{}
	for _, p := range req.Payments {
		var payment domain.InvoicePayment
		createdAt, _ := time.Parse(time.RFC3339Nano, p.CreatedAt)

		payment.UUID = p.UUID
		payment.OrderGroupUUID = p.OrderGroupUUID
		payment.Total = p.Total
		payment.Fee = p.Fee
		payment.CreatedAt = createdAt.UnixMicro()

		payment.PaymentMethod.PaymentMethodUUID = p.PaymentMethod.PaymentMethodUUID
		payment.PaymentMethod.Group = p.PaymentMethod.Group
		payment.PaymentMethod.Name = p.PaymentMethod.Name

		payment.PaymentMethod.Fee.Fixed = p.PaymentMethod.Fee.Fixed
		payment.PaymentMethod.Fee.Percent = p.PaymentMethod.Fee.Percent

		payment.EmployeeShiftInvoice.EmployeeShiftUUID = p.EmployeeShift.EmployeeShiftUUID
		payment.EmployeeShiftInvoice.UserUUID = p.EmployeeShift.UserUUID
		payment.EmployeeShiftInvoice.Name = p.EmployeeShift.Name

		data.Payments = append(data.Payments, payment)
	}

	// order groups
	data.OrderGroups = []domain.OrderGroup{}
	for _, og := range req.OrderGroups {
		var dataOg domain.OrderGroup
		dataOg.BranchUUID = branchId
		DTOtoDomain_UpsertActiveOrder(&dataOg, &og)

		data.OrderGroups = append(data.OrderGroups, dataOg)
	}

	// credit contracts
	data.CreditContracts = []domain.InvoiceCreditContract{}
	for _, c := range req.CreditContracts {
		var cc domain.InvoiceCreditContract
		dueBy, _ := time.Parse(time.RFC3339Nano, c.DueBy)
		createdAt, _ := time.Parse(time.RFC3339Nano, c.CreatedAt)

		cc.UUID = c.UUID
		cc.Number = c.Number
		cc.Note = c.Note
		cc.DueBy = dueBy.UnixMicro()
		cc.CreatedAt = createdAt.UnixMicro()

		data.CreditContracts = append(data.CreditContracts, cc)
	}
}

func DomainToDTO_InsertInvoice(resp *dto.InvoiceResponse, data *domain.Invoice) {
	resp.UUID = data.UUID
	resp.BranchUUID = data.BranchUUID
	resp.TotalAmount = data.TotalAmount
	resp.TotalTax = data.TotalTax
	resp.TotalDiscount = data.TotalDiscount
	resp.TotalChange = data.TotalChange
	resp.TotalTip = data.TotalTip
	resp.Note = data.Note
	resp.CreatedAt = time.UnixMicro(data.CreatedAt)
	if data.CompletedAt != nil {
		completedAt := time.UnixMicro(*data.CompletedAt)
		resp.CompletedAt = &completedAt
	}
	if data.UpdatedAt != nil {
		updatedAt := time.UnixMicro(*data.UpdatedAt)
		resp.UpdatedAt = &updatedAt
	}

	if data.Customer != nil {
		var customer dto.Customer

		customer.Name = data.Customer.Name
		customer.Phone = data.Customer.Phone
		customer.Address = data.Customer.Address

		resp.Customer = &customer
	}

	// payments
	resp.Payments = []dto.InvoicePaymentResponse{}
	for _, p := range data.Payments {
		var payment dto.InvoicePaymentResponse
		payment.UUID = p.UUID
		payment.OrderGroupUUID = p.OrderGroupUUID
		payment.Total = p.Total
		payment.Fee = p.Fee
		payment.CreatedAt = time.UnixMicro(p.CreatedAt)

		payment.PaymentMethod.PaymentMethodUUID = p.PaymentMethod.PaymentMethodUUID
		payment.PaymentMethod.Group = p.PaymentMethod.Group
		payment.PaymentMethod.Name = p.PaymentMethod.Name
		payment.PaymentMethod.Fee.Fixed = p.PaymentMethod.Fee.Fixed
		payment.PaymentMethod.Fee.Percent = p.PaymentMethod.Fee.Percent

		payment.EmployeeShift.EmployeeShiftUUID = p.EmployeeShiftInvoice.EmployeeShiftUUID
		payment.EmployeeShift.UserUUID = p.EmployeeShiftInvoice.UserUUID
		payment.EmployeeShift.Name = p.EmployeeShiftInvoice.Name

		resp.Payments = append(resp.Payments, payment)
	}

	// credit contracts
	resp.CreditContracts = []dto.InvoiceCreditContractResponse{}
	for _, c := range data.CreditContracts {
		var cc dto.InvoiceCreditContractResponse
		cc.UUID = c.UUID
		cc.Number = c.Number
		cc.Note = c.Note
		cc.DueBy = time.UnixMicro(c.DueBy)
		cc.CreatedAt = time.UnixMicro(c.CreatedAt)

		resp.CreditContracts = append(resp.CreditContracts, cc)
	}

	// order groups
	resp.OrderGroups = []dto.OrderGroupResponse{}
	for _, og := range data.OrderGroups {
		var respOg dto.OrderGroupResponse

		DomainToDTO_UpsertActiveOrder(&respOg, &og)

		resp.OrderGroups = append(resp.OrderGroups, respOg)
	}

	resp.Refunds = []dto.InvoiceRefundResponse{}
}
