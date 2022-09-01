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

	userRole, code, err := u.userRoleRepo.FindUserRole(ctx, user.RoleUUID, false)
	if err != nil {
		return nil, code, err
	}
	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	token, code, err := u.tokenRepo.FindTokenWithATandRT(ctx, oldAT, oldRT)
	if err != nil {
		return nil, code, err
	}
	if code == http.StatusNotFound {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	if token.AccessToken != oldAT || token.RefreshToken != oldRT {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	accesstoken, refreshtoken, err := jwthandler.GenerateAllTokens(user.UUID, userRole.Name, user.BranchUUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	tokenUUID, code, err := u.tokenRepo.RefreshTokens(ctx, user.UUID, oldAT, oldRT, accesstoken, refreshtoken)
	if err != nil {
		return nil, code, err
	}

	user, code, err = u.userRepo.RemoveToken(ctx, user.UUID, token.UUID)
	if err != nil {
		return nil, code, err
	}

	user, code, err = u.userRepo.InsertToken(ctx, user.UUID, tokenUUID)
	if err != nil {
		return nil, code, err
	}

	var resp domain.UserResponse
	resp.UUID = user.UUID
	resp.Name = user.Name
	resp.Role = user.Role
	resp.Token = &accesstoken
	resp.RefreshToken = &refreshtoken

	return &resp, http.StatusOK, nil
}
