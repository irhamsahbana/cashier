package domain

import (
	"context"
)

type ItemCategory struct {
	UUID       string `bson:"uuid"`
	BranchUUID string `bson:"branch_uuid"`
	Name       string `bson:"name"`
	Items      []Item `bson:"menus"`
	CreatedAt  int64  `bson:"created_at"`
	UpdatedAt  *int64 `bson:"updated_at,omitempty"`
	DeletedAt  *int64 `bson:"deleted_at,omitempty"`
}

type Item struct {
	UUID        string  `bson:"uuid"`
	MainUUID    *string `bson:"main_uuid"`
	Name        string  `bson:"name"`
	Price       float32 `bson:"price"`
	Description *string `bson:"description"`
	Label       string  `bson:"label"`
	Public      bool    `bson:"public"`
	ImageUrl    *string `bson:"image_url"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at,omitempty"`
	DeletedAt   *int64  `bson:"deleted_at,omitempty"`
}

type ItemCategoryUsecaseContract interface {
	UpsertItemCategory(ctx context.Context, branchId string, req *ItemCategoryUpsertRequest) (*ItemCategoryResponse, int, error)
	FindItemCategories(ctx context.Context, branchId string, withTrashed bool) ([]ItemCategoryResponse, int, error)
	FindItemCategory(ctx context.Context, branchId, id string, withTrashed bool) (*ItemCategoryResponse, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*ItemCategoryResponse, int, error)

	CreateItem(ctx context.Context, branchId, itemCategoryId string, data *ItemCreateRequest) (*ItemResponse, int, error)
	UpdateItem(ctx context.Context, branchId, id string, req *ItemUpdateRequest) (*ItemResponse, int, error)
	FindItem(ctx context.Context, branchId, id string, withTrashed bool) (*ItemResponse, int, error)
	DeleteItem(ctx context.Context, branchId, id string) (*ItemResponse, int, error)
}

type ItemCategoryRepositoryContract interface {
	UpsertItemCategory(ctx context.Context, branchId string, data *ItemCategory) (*ItemCategory, int, error)
	FindItemCategories(ctx context.Context, branchId string, withTrashed bool) ([]ItemCategory, int, error)
	FindItemCategory(ctx context.Context, branchId, id string, withTrashed bool) (*ItemCategory, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*ItemCategory, int, error)

	InsertItem(ctx context.Context, branchId, itemCategoryId string, data *Item) (*ItemCategory, int, error)
	UpdateItem(ctx context.Context, branchId, id string, data *Item) (*ItemCategory, int, error)
	FindItem(ctx context.Context, branchId, id string, withTrashed bool) (*ItemCategory, int, error)
	DeleteItem(ctx context.Context, branchId, id string) (*ItemCategory, int, error)
}
