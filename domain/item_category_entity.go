package domain

import (
	"context"
)

type ItemCategory struct {
	UUID       string `json:"uuid" bson:"uuid"`
	BranchUUID string `json:"branch_uuid" bson:"branch_uuid"`
	Name       string `json:"name" bson:"name"`
	Items      []Item `json:"menus" bson:"menus"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  *int64 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt  *int64 `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type Item struct {
	UUID        string  `json:"uuid" bson:"uuid"`
	MainUUID    *string `json:"main_uuid,omitempty" bson:"main_uuid"`
	Name        string  `json:"name" bson:"name"`
	Price       float32 `json:"price" bson:"price"`
	Description *string `json:"description" bson:"description"`
	Label       string  `json:"label" bson:"label"`
	Public      bool    `json:"public" bson:"public"`
	ImageUrl    *string `json:"image_url" bson:"image_url"`
	CreatedAt   int64   `json:"created_at" bson:"created_at"`
	UpdatedAt   *int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   *int64  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type ItemCategoryUsecaseContract interface {
	UpsertItemCategory(ctx context.Context, data *ItemCategoryUpsertRequest) (*ItemCategoryResponse, int, error)
	FindItemCategories(ctx context.Context, withTrashed bool) ([]ItemCategoryFindAllResponse, int, error)
	FindItemCategory(ctx context.Context, id string, withTrashed bool) (*ItemCategoryResponse, int, error)
	DeleteItemCategory(ctx context.Context, id string) (*ItemCategoryResponse, int, error)

	CreateItem(ctx context.Context, menuCategoryId string, data *ItemCreateRequest) (*ItemResponse, int, error)
	UpdateItem(ctx context.Context, id string, payload *ItemUpdateRequest) (*ItemResponse, int, error)
	FindItem(ctx context.Context, id string, withTrashed bool) (*ItemResponse, int, error)
	DeleteItem(ctx context.Context, id string) (*ItemResponse, int, error)
}

type ItemCategoryRepositoryContract interface {
	UpsertItemCategory(ctx context.Context, data *ItemCategory) (*ItemCategory, int, error)
	FindItemCategories(ctx context.Context, withTrashed bool) ([]ItemCategory, int, error)
	FindItemCategory(ctx context.Context, id string, withTrashed bool) (*ItemCategory, int, error)
	DeleteItemCategory(ctx context.Context, id string) (*ItemCategory, int, error)

	InsertItem(ctx context.Context, menuCategoryId string, data *Item) (*ItemCategory, int, error)
	UpdateItem(ctx context.Context, id string, data *Item) (*ItemCategory, int, error)
	FindItem(ctx context.Context, id string, withTrashed bool) (*ItemCategory, int, error)
	DeleteItem(ctx context.Context, id string) (*ItemCategory, int, error)
}
