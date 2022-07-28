package domain

import (
	"context"
	"time"
)

type MenuCategory struct {
	ID			uint64			`json:"id" bson:"id"`
	UUID		string			`json:"uuid" bson:"uuid"`
	BranchID	uint64			`json:"branch_id" bson:"branch_id"`
	Name		string			`json:"name" bson:"name" validate:"required"`
	// Menus		*[]Menu			`json:"menus" bson:"menus"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuCategoryRepositoryContract interface {
	InsertMenuCategory(ctx context.Context, entity *MenuCategory) (*MenuCategory, error)
	FindMenuCategory(ctx context.Context, id string) (*MenuCategory, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategory, error)
	UpdateMenuCategory(ctx context.Context, id string, entity *MenuCategory) (*MenuCategory, error)
}

type MenuCategoryUsecaseContract interface {
	CreateMenuCategory(ctx context.Context, entity *MenuCategory) (*MenuCategory, int, error)
	FindMenuCategory(ctx context.Context, id string) (*MenuCategory, int, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategory, int, error)
	UpdateMenuCategory(ctx context.Context, id string, entity *MenuCategory) (*MenuCategory, int, error)
}

