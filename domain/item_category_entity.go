package domain

import (
	"context"
)

type ModifierGroupCondition string

const (
	ModifierGroupCondition_MIN   ModifierGroupCondition = "MIN"
	ModifierGroupCondition_MAX   ModifierGroupCondition = "MAX"
	ModifierGroupCondition_EQUAL ModifierGroupCondition = "EQUAL"
)

type ItemCategory struct {
	UUID          string          `bson:"uuid"`
	BranchUUID    string          `bson:"branch_uuid"`
	Name          string          `bson:"name"`
	ModifierGroup []ModifierGroup `bson:"modifier_group,omitempty"`
	Items         []Item          `bson:"items,omitempty"`
	CreatedAt     int64           `bson:"created_at"`
	UpdatedAt     *int64          `bson:"updated_at,omitempty"`
	DeletedAt     *int64          `bson:"deleted_at,omitempty"`
}

type Item struct {
	UUID        string  `bson:"uuid"`
	MainUUID    *string `bson:"main_uuid,omitempty"`
	Name        string  `bson:"name"`
	Price       float32 `bson:"price"`
	Label       string  `bson:"label"`
	Public      bool    `bson:"public"`
	ImageUrl    *string `bson:"image_url"`
	Description *string `bson:"description"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at,omitempty"`
	DeletedAt   *int64  `bson:"deleted_at,omitempty"`
}

type ModifierGroup struct {
	UUID      string                  `bson:"uuid"`
	Name      string                  `bson:"name"`
	Condition *ModifierGroupCondition `bson:"condition,omitempty"`
	Quantity  *int                    `bson:"quantity,omitempty"`
	Single    bool                    `bson:"single"`
	Required  bool                    `bson:"required"`
	Modifiers []Modifier              `bson:"modifiers,omitempty"`
}

type Modifier struct {
	UUID      string  `bson:"uuid"`
	Name      string  `bson:"name"`
	Price     float32 `bson:"price"`
	CreatedAt int64   `bson:"created_at"`
	UpdatedAt *int64  `bson:"updated_at,omitempty"`
	DeletedAt *int64  `bson:"deleted_at,omitempty"`
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
