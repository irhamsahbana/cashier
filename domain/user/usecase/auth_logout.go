package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
)

func (u *userUsecase) Logout(c context.Context, accessToken string, userId string) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, code, err := u.userRepo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	code, err = u.userRepo.RevokeToken(ctx, accessToken)
	if err != nil {
		return nil, code, err
	}

	if code != http.StatusOK {
		return nil, code, nil
	}

	return nil, code, nil
}
