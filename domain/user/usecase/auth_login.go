package usecase

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	jwthandler "lucy/cashier/lib/jwt_handler"
	passwordhandler "lucy/cashier/lib/password_handler"
	"net/http"
)

func (u *userUsecase) Login(c context.Context, req *domain.UserLoginRequest) (*domain.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "email", req.Email, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if ok := passwordhandler.VerifyPassword(user.Password, req.Password); !ok {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(user.UUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	aToken, rToken, code, err  := u.userRepo.GenerateTokens(ctx, user.UUID, accesstoken, refreshtoken)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UserResponse
	resp.UUID = user.UUID
	resp.Name = user.Name
	resp.Role = user.Role
	resp.Token = &aToken
	resp.RefreshToken = &rToken

	return &resp, http.StatusOK, nil
}
