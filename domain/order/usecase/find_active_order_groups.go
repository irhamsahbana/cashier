package usecase

import (
	"context"
	"lucy/cashier/dto"
)

func (u *orderUsecase) FindActiveOrders(c context.Context, branchId string) ([]dto.OrderGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var resp []dto.OrderGroupResponse
	result, code, err := u.orderRepo.FindActiveOrders(ctx, branchId)
	if err != nil {
		return resp, code, err
	}

	for _, v := range result {
		var orderGroup dto.OrderGroupResponse
		orderDomainToDTO(&orderGroup, &v)
		resp = append(resp, orderGroup)
	}

	return resp, code, nil
}
