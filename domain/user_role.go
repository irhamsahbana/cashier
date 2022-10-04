package domain

import (
	"context"
	"time"
)

type UserRole struct {
	UUID      string `bson:"uuid"`
	Name      string `bson:"name"`
	Power     int    `bson:"power"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt *int64 `bson:"updated_at,omitempty"`
	DeletedAt *int64 `bson:"deleted_at,omitempty"`
}

type UserRoleModel struct {
	UUID      string     `bson:"uuid"`
	Name      string     `bson:"name"`
	Power     int        `bson:"power"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

type UserRoleUsecaseContract interface {
	UpsertUserRole(ctx context.Context, data *UserRoleUpsertrequest) (*UserRoleModel, int, error)
	// FindUserRole(ctx context.Context, id string, withTrashed bool) (*UserRoleResponse, int, error)
	// DeleteUserRole(ctx context.Context, id string) (*UserRoleResponse, int, error)
}

type UserRoleRepositoryContract interface {
	FindUserRole(ctx context.Context, id string, withTrashed bool) (*UserRole, int, error)
	FindUserRoleByName(ctx context.Context, name string, withTrashed bool) (*UserRole, int, error)
	UpsertUserRole(ctx context.Context, userRole *UserRole) (*UserRole, int, error)
	DeleteUserRole(ctx context.Context, id string) (*UserRole, int, error)
}
