package domain

import (
	"context"
	"time"
)

type MenuCategory struct {
	UUID		string			`json:"uuid" bson:"uuid"`
	BranchUUID	string			`json:"branch_uuid" bson:"branch_uuid"`
	Name		string			`json:"name" bson:"name"`
	Menus		[]Menu			`json:"menus" bson:"menus"`
	CreatedAt	int64			`json:"created_at" bson:"created_at"`
	UpdatedAt	*int64			`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*int64		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type Menu struct {
	UUID			string		`json:"uuid" bson:"uuid"`
	MainUUID		string		`json:"main_uuid,omitempty" bson:"main_uuid"`
	Name			string		`json:"name" bson:"name" validate:"required"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	ImageUrl		*string		`json:"image_url" bson:"image_url"`
	CreatedAt		int64		`json:"created_at" bson:"created_at"`
	UpdatedAt		*int64		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt		*int64		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuCategoryUpsertRequest struct {
	UUID		string			`json:"uuid" bson:"uuid"`
	Name		string			`json:"name" bson:"name"`
	CreatedAt	string			`json:"created_at" bson:"created_at"`
}

type MenuCategoryUpsertResponse struct {
	UUID		string			`json:"uuid"`
	Name		string			`json:"name"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty"`
}

type MenuCreateRequestResponse struct {
	UUID			string		`json:"uuid" bson:"uuid"`
	Name			string		`json:"name" bson:"name"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	CreatedAt		int64		`json:"created_at" bson:"created_at"`
}

type MenuCategoryUsecaseContract interface {
	UpsertMenuCategory(ctx context.Context, payload *MenuCategoryUpsertRequest) (*MenuCategoryUpsertResponse, int, error)
	FindMenuCategories(ctx context.Context, withTrashed bool) ([]MenuCategory, int, error)
	FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*MenuCategory, int, error)
	DeleteMenuCategory(ctx context.Context, id string, forceDelete bool) (*MenuCategory, int, error)

	CreateMenu(ctx context.Context, menuCategoryId string, payload *MenuCreateRequestResponse) (*MenuCategory, int, error)
	FindMenu(ctx context.Context, id string, withTrashed bool) (*MenuCategory, int, error)
	DeleteMenu(ctx context.Context, id string, forceDelete bool) (*MenuCategory, int, error)
}

type MenuCategoryRepositoryContract interface {
	UpsertMenuCategory(ctx context.Context, payload *MenuCategory) (*MenuCategory, error)
	FindMenuCategories(ctx context.Context, withTrashed bool) ([]MenuCategory, error)
	FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*MenuCategory, error)
	DeleteMenuCategory(ctx context.Context, id string, forceDelete bool) (*MenuCategory, error)

	InsertMenu(ctx context.Context, menuCategoryId string, data *Menu) (*MenuCategory, error)
	FindMenu(ctx context.Context, id string, withTrashed bool) (*MenuCategory, error)
	DeleteMenu(ctx context.Context, id string, forceDelete bool) (*MenuCategory, error)
}