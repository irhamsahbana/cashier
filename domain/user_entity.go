package domain

import (
	"context"
)

type User struct {
	UUID				string		`bson:"uuid"`
	Name				string		`bson:"name"`
	Role				string		`bson:"role"`
	Email				string		`bson:"email"`
	Password			string		`bson:"password"`
	Phone				*string		`bson:"phone,omitempty"`
	WA					*string		`bson:"wa,omitempty"`
	ProfileUrl			*string		`bson:"profile_url,omitempty"`
	Tokens				[]Token		`bson:"tokens,omitempty"`
	CreatedAt			int64	 	`bson:"created_at"`
	UpdatedAt			*int64		`bson:"updated_at,omitempty"`
	DeletedAt			*int64		`bson:"deleted_at,omitempty"`
}

type Token struct {
	AccessToken		string	`bson:"access_token"`
	RefreshToken	string	`bson:"refresh_token"`
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
