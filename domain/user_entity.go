package domain

import (
	"context"
)

type User struct {
	UUID                   string   `bson:"uuid"`
	BranchUUID             string   `bson:"branch_uuid"`
	RoleUUID               string   `bson:"role_uuid"`
	Name                   string   `bson:"name"`
	Email                  string   `bson:"email"`
	EmailVerifiedAt        *int64   `bson:"email_verified_at,omitempty"`
	EmailVerificationCodes []string `bson:"email_verification_code,omitempty"`
	Password               string   `bson:"password"`
	Phone                  *string  `bson:"phone,omitempty"`
	Whatsapp               *string  `bson:"whatsapp,omitempty"`
	ProfileUrl             *string  `bson:"profile_url,omitempty"`
	Tokens                 []string `bson:"tokens,omitempty"`
	CreatedAt              int64    `bson:"created_at"`
	UpdatedAt              *int64   `bson:"updated_at,omitempty"`
	DeletedAt              *int64   `bson:"deleted_at,omitempty"`
}

type UserUsecaseContract interface {
	FindUser(ctx context.Context, id string, withTrashed bool) (*UserResponse, int, error)

	Login(ctx context.Context, request *UserLoginRequest) (*UserResponse, int, error)
	RefreshToken(ctx context.Context, oldAccessToken string, oldRefreshToken string, userId string) (*UserResponse, int, error)
	Logout(ctx context.Context, accessToken string, userId string) (*UserResponse, int, error)
}

type UserRepositoryContract interface {
	FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*User, int, error)

	InsertToken(ctx context.Context, userId, tokenId string) (*User, int, error)
	RemoveToken(ctx context.Context, userId, tokenId string) (*User, int, error)
}
