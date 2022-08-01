package domain

import (
	"context"
	"time"
)

type MenuCategory struct {
	ID			uint64			`json:"id" bson:"id"`
	UUID		string			`json:"uuid" bson:"uuid"`
	BranchID	uint64			`json:"branch_id" bson:"branch_id"`
	Name		string			`json:"name" bson:"name"`
	Menus		[]Menu			`json:"menus" bson:"menus"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type Menu struct {
	ID				uint64		`json:"id" bson:"id"`
	UUID			string		`json:"uuid" bson:"uuid"`
	Name			string		`json:"name" bson:"name" validate:"required"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	ImageUrl		*string		`json:"image_url" bson:"image_url"`
	CreatedAt		time.Time	`json:"created_at" bson:"created_at"`
	UpdatedAt		*time.Time	`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt		*time.Time	`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuCategoryUpdateRequest struct {
	Name		string		`json:"name" bson:"name"`
}

type MenuCreateRequestResponse struct {
	UUID			string		`json:"uuid" bson:"uuid"`
	Name			string		`json:"name" bson:"name"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	CreatedAt		time.Time	`json:"created_at" bson:"created_at"`
}

type MenuCategoryRepositoryContract interface {
	InsertMenuCategory(ctx context.Context, payload *MenuCategory) (*MenuCategory, error)
	FindMenuCategory(ctx context.Context, id string) (*MenuCategory, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategory, error)
	UpdateMenuCategory(ctx context.Context, id string, payload *MenuCategoryUpdateRequest) (*MenuCategory, error)

	InsertMenu(ctx context.Context, menuCategoryId string, menu *Menu) (*Menu, error)
}

type MenuCategoryUsecaseContract interface {
	CreateMenuCategory(ctx context.Context, payload *MenuCategory) (*MenuCategory, int, error)
	FindMenuCategory(ctx context.Context, id string) (*MenuCategory, int, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategory, int, error)
	UpdateMenuCategory(ctx context.Context, id string, payload *MenuCategoryUpdateRequest) (*MenuCategory, int, error)

	CreateMenu(ctx context.Context, menuCategoryId string, payload *MenuCreateRequestResponse) (*MenuCreateRequestResponse, int, error)
}

