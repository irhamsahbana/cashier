package domain

import (
	"context"
	"time"
)

type User struct {
	UUID                   string   `bson:"uuid"`
	BranchUUID             string   `bson:"branch_uuid"`
	RoleUUID               string   `bson:"role_uuid"`
	Name                   string   `bson:"name"`
	Role                   string   `bson:"role"`
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

type UserModel struct {
	UUID                   string                  `bson:"uuid"`
	BranchUUID             string                  `bson:"branch_uuid"`
	RoleUUID               string                  `bson:"role_uuid"`
	Name                   string                  `bson:"name"`
	Role                   string                  `bson:"role"`
	Email                  string                  `bson:"email"`
	EmailVerifiedAt        *int64                  `bson:"email_verified_at,omitempty"`
	EmailVerificationCodes []EmailVerificationCode `bson:"email_verification_code,omitempty"`
	Password               string                  `bson:"password"`
	Phone                  *string                 `bson:"phone,omitempty"`
	Whatsapp               *string                 `bson:"whatsapp,omitempty"`
	ProfileUrl             *string                 `bson:"profile_url,omitempty"`
	Tokens                 []Token                 `bson:"tokens,omitempty"`
	CreatedAt              time.Time               `bson:"created_at"`
	UpdatedAt              *time.Time              `bson:"updated_at,omitempty"`
	DeletedAt              *time.Time              `bson:"deleted_at,omitempty"`
}

type UserUsecaseContract interface {
	FindUser(ctx context.Context, id string, withTrashed bool) (*UserResponse, int, error)

	Login(ctx context.Context, request *UserLoginRequest) (*UserResponse, int, error)
	RefreshToken(ctx context.Context, oldAccessToken string, oldRefreshToken string, userId string) (*UserResponse, int, error)
	Logout(ctx context.Context, accessToken string, userId string) (*UserResponse, int, error)
}

type UserRepositoryContract interface {
	FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*User, int, error)

	GenerateTokens(ctx context.Context, userId, accessToken, refreshToken string) (aToken, rToken string, code int, err error)
	RefreshToken(ctx context.Context, oldAT, oldRT, newAT, newRT, userId string) (aToken, rToken string, code int, err error)
	RevokeToken(ctx context.Context, accessToken string) (code int, err error)
}
