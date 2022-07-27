package domain

import (
	"context"
	"time"
)

type MenuCategory struct {
	ID			uint64			`json:"id" bson:"id"`
	UUID		string			`json:"uuid" bson:"uuid" validate:"required"`
	BranchID	uint64			`json:"branch_id" bson:"branch_id" validate:"required"`
	Name		string			`json:"name" bson:"name" validate:"required"`
	// Menus		[]*Menu			`json:"menus" bson:"menus"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuCategoryRepositoryContract interface {
	InsertOne(ctx context.Context, entity *MenuCategory) (*MenuCategory, error)
	FindOne(ctx context.Context, id string) (*MenuCategory, error)
}

type MenuCategoryUsecaseContract interface {
	CreateMenuCategory(ctx context.Context, entity *MenuCategory) (*MenuCategory, error, int)
	FindMenuCategory(ctx context.Context, id string) (*MenuCategory, error, int)
}

