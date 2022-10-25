package usecase

import (
	"context"
	"lucy/cashier/dto"
	"net/http"
)

func (u *orderUsecase) FindInvoiceHistories(c context.Context) ([]dto.InvoiceResponse, *int64, *int64, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	resp := []dto.InvoiceResponse{}
	result, nextCur, prevCur, code, err := u.orderRepo.FindInvoiceHistories(ctx)
	if err != nil {
		return resp, nil, nil, code, err
	}
	for _, v := range result {
		var invoice dto.InvoiceResponse
		DomainToDTO_InsertInvoice(&invoice, &v)
		resp = append(resp, invoice)
	}

	return resp, nextCur, prevCur, http.StatusOK, nil
}
