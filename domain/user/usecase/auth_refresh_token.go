package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"net/http"
)

func (u *userUsecase) RefreshToken(c context.Context, oldAT, oldRT, userId string) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "uuid", userId, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}


	newAT, newRT, err := jwthandler.GenerateAllTokens(user.UUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	aToken, rToken, code, err := u.userRepo.RefreshToken(ctx, oldAT, oldRT, newAT, newRT, userId)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	var resp domain.UserResponse
	resp.UUID = user.UUID
	resp.Name = user.Name
	resp.Role = user.Role
	resp.Token = &aToken
	resp.RefreshToken = &rToken

	return &resp, http.StatusOK, nil
}
