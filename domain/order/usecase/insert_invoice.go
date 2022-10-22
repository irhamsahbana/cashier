package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
)

func (u *orderUsecase) InsertInvoice(ctx context.Context, branchId string, req *dto.InvoiceInsertRequest) (interface{}, int, error) {
	return nil, 200, nil

}

func DTOtoDomain_InsertInvoice(data *domain.Invoice, req *dto.InvoiceInsertRequest) {

}

func DomainToDTO_InsertInvoice(resp dto.InvoiceResponse, data *domain.Invoice) {

}
