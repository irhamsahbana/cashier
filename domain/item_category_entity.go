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
	UUID           string          `bson:"uuid"`
	BranchUUID     string          `bson:"branch_uuid"`
	Name           string          `bson:"name"`
	ModifierGroups []ModifierGroup `bson:"modifier_groups,omitempty"`
	Items          []Item          `bson:"items"`
	CreatedAt      int64           `bson:"created_at"`
	UpdatedAt      *int64          `bson:"updated_at,omitempty"`
}

type Item struct {
	UUID        string    `bson:"uuid"`
	Name        string    `bson:"name"`
	Price       float32   `bson:"price"`
	Label       string    `bson:"label"`
	Public      bool      `bson:"public"`
	ImagePath   *string   `bson:"image_path,omitempty"`
	Description *string   `bson:"description"`
	Variants    []Variant `bson:"variants"`
	CreatedAt   int64     `bson:"created_at"`
	UpdatedAt   *int64    `bson:"updated_at,omitempty"`
}

type Variant struct {
	UUID      string  `bson:"uuid"`
	Price     float32 `bson:"price"`
	Label     string  `bson:"label"`
	Public    bool    `bson:"public"`
	ImagePath *string `bson:"image_path"`
	CreatedAt int64   `bson:"created_at"`
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
}

type ItemCategoryUsecaseContract interface {
	UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, req *ItemCategoryUpsertRequest) (*ItemCategoryResponse, int, error)
	FindItemCategories(ctx context.Context, branchId string) ([]ItemCategoryResponse, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*ItemCategoryResponse, int, error)

	UpsertItemAndVariants(ctx context.Context, branchId, itemCategoryId string, req *ItemAndVariantsUpsertRequest) (*ItemResponse, int, error)
	FindItemAndVariants(ctx context.Context, branchId, id string) (*ItemResponse, int, error)
	DeleteItemAndVariants(ctx context.Context, branchId, id string) (*ItemResponse, int, error)
}

type ItemCategoryRepositoryContract interface {
	UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, data *ItemCategory) (*ItemCategory, int, error)
	FindItemCategories(ctx context.Context, branchId string) ([]ItemCategory, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*ItemCategory, int, error)

	UpsertItemAndVariants(ctx context.Context, branchId, itemCategoryId string, data *Item) (*ItemCategory, int, error)
	FindItemAndVariants(ctx context.Context, branchId, id string) (*ItemCategory, int, error)
	DeleteItemAndVariants(ctx context.Context, branchId, id string) (*ItemCategory, int, error)
}
