package usecase

import (
	"context"
	"errors"
	"lucy/cashier/dto"
	jwthandler "lucy/cashier/lib/jwt_handler"
	passwordhandler "lucy/cashier/lib/password_handler"
	"net/http"
)

func (u *userUsecase) Login(c context.Context, req *dto.UserLoginRequest) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	userResult, code, err := u.userRepo.FindUserBy(ctx, "email", req.Email, false)
	if err != nil {
		return nil, code, err
	}

	userRoleResult, code, err := u.userRoleRepo.FindUserRole(ctx, userResult.RoleUUID, false)
	if err != nil {
		return nil, code, err
	}

	if ok := passwordhandler.VerifyPassword(userResult.Password, req.Password); !ok {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(userResult.UUID, userRoleResult.Name, userResult.BranchUUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenUUID, code, err := u.tokenRepo.GenerateTokens(ctx, userResult.UUID, accesstoken, refreshtoken)
	if err != nil {
		return nil, code, err
	}

	userResult, code, err = u.userRepo.InsertToken(ctx, userResult.UUID, tokenUUID)
	if err != nil {
		return nil, code, err
	}

	var resp dto.UserResponse
	resp.UUID = userResult.UUID
	resp.BranchUUID = userResult.BranchUUID
	resp.RoleUUID = &userResult.RoleUUID
	resp.Name = userResult.Name
	resp.Role = userRoleResult.Name
	resp.Token = &accesstoken
	resp.RefreshToken = &refreshtoken

	return &resp, http.StatusOK, nil
}
