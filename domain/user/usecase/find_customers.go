package usecase

import (
	"context"
	"fmt"
	"lucy/cashier/dto"
)

func (u *userUsecase) FindCustomers(c context.Context, branchId string, limit, page int, withTrashed bool) ([]dto.CustomerResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	userRole, code, err := u.userRoleRepo.FindUserRoleByName(ctx, "Customer", true)
	if err != nil {
		return nil, code, err
	}

	fmt.Println(withTrashed)

	result, code, err := u.userRepo.FindUsers(ctx, branchId, []string{userRole.UUID}, limit, page, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp []dto.CustomerResponse
	for _, r := range result {
		var e dto.CustomerResponse

		userDomainToDTO_UpsertCustomer(&e, &r)
		resp = append(resp, e)
	}

	if len(resp) == 0 {
		resp = []dto.CustomerResponse{}
	}

	return resp, code, nil
}
