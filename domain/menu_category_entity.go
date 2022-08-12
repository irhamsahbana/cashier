package domain

import (
	"context"
)

type MenuCategory struct {
	UUID		string			`json:"uuid" bson:"uuid"`
	BranchUUID	string			`json:"branch_uuid" bson:"branch_uuid"`
	Name		string			`json:"name" bson:"name"`
	Menus		[]Menu			`json:"menus" bson:"menus"`
	CreatedAt	int64			`json:"created_at" bson:"created_at"`
	UpdatedAt	*int64			`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*int64			`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type Menu struct {
	UUID			string		`json:"uuid" bson:"uuid"`
	MainUUID		string		`json:"main_uuid,omitempty" bson:"main_uuid"`
	Name			string		`json:"name" bson:"name"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	ImageUrl		*string		`json:"image_url" bson:"image_url"`
	CreatedAt		int64		`json:"created_at" bson:"created_at"`
	UpdatedAt		*int64		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt		*int64		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuCategoryUsecaseContract interface {
	UpsertMenuCategory(ctx context.Context, payload *MenuCategoryUpsertRequest) (*MenuCategoryResponse, int, error)
	FindMenuCategories(ctx context.Context, withTrashed bool) ([]MenuCategoryFindAllResponse, int, error)
	FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*MenuCategoryResponse, int, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategoryResponse, int, error)

	CreateMenu(ctx context.Context, menuCategoryId string, payload *MenuCreateRequest) (*MenuResponse, int, error)
	UpdateMenu(ctx context.Context, id string, payload *MenuUpdateRequest) (*MenuResponse, int, error)
	FindMenu(ctx context.Context, id string, withTrashed bool) (*MenuResponse, int, error)
	DeleteMenu(ctx context.Context, id string) (*MenuResponse, int, error)
}

type MenuCategoryRepositoryContract interface {
	UpsertMenuCategory(ctx context.Context, payload *MenuCategory) (*MenuCategory, int, error)
	FindMenuCategories(ctx context.Context, withTrashed bool) ([]MenuCategory, int, error)
	FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*MenuCategory, int, error)
	DeleteMenuCategory(ctx context.Context, id string) (*MenuCategory, int, error)

	InsertMenu(ctx context.Context, menuCategoryId string, data *Menu) (*MenuCategory, int, error)
	UpdateMenu(ctx context.Context, id string, data *Menu) (*MenuCategory, int, error)
	FindMenu(ctx context.Context, id string, withTrashed bool) (*MenuCategory, int, error)
	DeleteMenu(ctx context.Context, id string) (*MenuCategory, int, error)
}