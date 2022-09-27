package domain

import (
	"context"
	"lucy/cashier/dto"
)

type ItemCategory struct {
	UUID           string          `bson:"uuid"`
	BranchUUID     string          `bson:"branch_uuid"`
	Name           string          `bson:"name"`
	ModifierGroups []ModifierGroup `bson:"modifier_groups"`
	Items          []Item          `bson:"items"`
	CreatedAt      int64           `bson:"created_at"`
	UpdatedAt      *int64          `bson:"updated_at"`
}

type Item struct {
	UUID        string    `bson:"uuid"`
	Name        string    `bson:"name"`
	Price       float32   `bson:"price"`
	Label       string    `bson:"label"`
	Public      bool      `bson:"public"`
	ImagePath   *string   `bson:"image_path"`
	Description string    `bson:"description"`
	Variants    []Variant `bson:"variants"`
	CreatedAt   int64     `bson:"created_at"`
	UpdatedAt   *int64    `bson:"updated_at"`
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
	UUID      string     `bson:"uuid"`
	Name      string     `bson:"name"`
	Quantity  *int       `bson:"quantity"`
	Single    bool       `bson:"single"`
	Required  bool       `bson:"required"`
	Modifiers []Modifier `bson:"modifiers"`
}

type Modifier struct {
	UUID      string  `bson:"uuid"`
	Name      string  `bson:"name"`
	Price     float32 `bson:"price"`
	CreatedAt int64   `bson:"created_at"`
	UpdatedAt *int64  `bson:"updated_at"`
}

type ItemCategoryUsecaseContract interface {
	UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, req *dto.ItemCategoryUpsertRequest) (*dto.ItemCategoryResponse, int, error)
	FindItemCategories(ctx context.Context, branchId string) ([]dto.ItemCategoryResponse, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*dto.ItemCategoryResponse, int, error)

	UpsertItemAndVariants(ctx context.Context, branchId, itemCategoryId string, req *dto.ItemAndVariantsUpsertRequest) (*dto.ItemResponse, int, error)
	FindItemAndVariants(ctx context.Context, branchId, id string) (*dto.ItemResponse, int, error)
	DeleteItemAndVariants(ctx context.Context, branchId, id string) (*dto.ItemResponse, int, error)
}

type ItemCategoryRepositoryContract interface {
	UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, data *ItemCategory) (*ItemCategory, int, error)
	FindItemCategories(ctx context.Context, branchId string) ([]ItemCategory, int, error)
	DeleteItemCategory(ctx context.Context, branchId, id string) (*ItemCategory, int, error)

	UpsertItemAndVariants(ctx context.Context, branchId, itemCategoryId string, data *Item) (*ItemCategory, int, error)
	FindItemAndVariants(ctx context.Context, branchId, id string) (*ItemCategory, int, error)
	DeleteItemAndVariants(ctx context.Context, branchId, id string) (*ItemCategory, int, error)
}
