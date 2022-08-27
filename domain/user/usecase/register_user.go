package usecase

import (
	"context"
	"lucy/cashier/domain"
)

func (u *userUsecase) RegisterUser(c context.Context, req *domain.UserRegisterRequest) (*domain.UserResponse, int, error) {
	return nil, 409, nil
}

func (u *userUsecase) ProfileUser(c context.Context, authUser *domain.User) (*domain.UserResponse, int, error) {
	return nil, 200, nil
}
