package domain

import (
	"context"
	"lucy/cashier/dto"
)

type User struct {
	UUID                   string   `bson:"uuid"`
	BranchUUID             string   `bson:"branch_uuid"`
	RoleUUID               string   `bson:"role_uuid"`
	Name                   string   `bson:"name"`
	Email                  string   `bson:"email"`
	EmailVerifiedAt        *int64   `bson:"email_verified_at"`
	EmailVerificationCodes []string `bson:"email_verification_code"`
	Password               string   `bson:"password"`
	Address                *string  `bson:"address"`
	Phone                  *string  `bson:"phone"`
	Whatsapp               *string  `bson:"whatsapp"`
	ProfileUrl             *string  `bson:"profile_url"`
	Tokens                 []string `bson:"tokens"`
	Dob                    *int64   `bson:"date_of_birth"`
	CreatedAt              int64    `bson:"created_at"`
	UpdatedAt              *int64   `bson:"updated_at"`
	DeletedAt              *int64   `bson:"deleted_at"`
}

type UserUsecaseContract interface {
	FindUser(ctx context.Context, id string, withTrashed bool) (*dto.UserResponse, int, error)
	UserBranchInfo(ctx context.Context, id string, withTrashed bool) (*dto.BranchResponse, int, error)

	UpsertCustomer(ctx context.Context, branchId string, req *dto.CustomerUpserRequest) (*dto.CustomerResponse, int, error)
	FindCustomers(ctx context.Context, branchId string, limit, page int, withTrashed bool) ([]dto.CustomerResponse, int, error)

	Login(ctx context.Context, request *dto.UserLoginRequest) (*dto.UserResponse, int, error)
	RefreshToken(ctx context.Context, oldAccessToken string, oldRefreshToken string, userId string) (*dto.UserResponse, int, error)
	Logout(ctx context.Context, accessToken string, userId string) (*dto.UserResponse, int, error)
}

type UserRepositoryContract interface {
	FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*User, int, error)

	UpsertUser(ctx context.Context, user *User) (*User, int, error)
	FindUsers(ctx context.Context, branchId string, roles []string, limit, page int, withTrashed bool) ([]User, int, error)

	InsertToken(ctx context.Context, userId, tokenId string) (*User, int, error)
	RemoveToken(ctx context.Context, userId, tokenId string) (*User, int, error)
}
