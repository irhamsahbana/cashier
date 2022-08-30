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

	userResult, code, err := u.userRepo.FindUserBy(ctx, "email", req.Email, false)
	if err != nil {
		return nil, code, err
	}

	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if ok := passwordhandler.VerifyPassword(userResult.Password, req.Password); !ok {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(userResult.UUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenUUID, code, err := u.tokenRepo.GenerateTokens(ctx, userResult.UUID, accesstoken, refreshtoken)
	if err != nil {
		return nil, code, err
	}

	userResult, code, err = u.userRepo.InsertToken(ctx, userResult.UUID, tokenUUID)

	var resp domain.UserResponse
	resp.UUID = userResult.UUID
	resp.Name = userResult.Name
	resp.Role = userResult.Role
	resp.Token = &accesstoken
	resp.RefreshToken = &refreshtoken

	return &resp, http.StatusOK, nil
}