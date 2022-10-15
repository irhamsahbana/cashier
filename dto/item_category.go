package dto

import "time"

type ItemCategoryResponse struct {
	UUID           string                  `json:"uuid"`
	BranchUUID     string                  `json:"branch_uuid"`
	Name           string                  `json:"name"`
	ModifierGroups []ModifierGroupResponse `json:"modifier_groups"`
	Items          []ItemResponse          `json:"items"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      *time.Time              `json:"updated_at"`
}

type ItemResponse struct {
	UUID        string            `json:"uuid"`
	Name        string            `json:"name"`
	Price       float64           `json:"price"`
	Description string            `json:"description"`
	Label       string            `json:"label"`
	Public      bool              `json:"public"`
	ImagePath   *string           `json:"image_path"`
	Variants    []VariantResponse `json:"variants"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   *time.Time        `json:"updated_at"`
}

type VariantResponse struct {
	UUID      string    `json:"uuid"`
	Label     string    `json:"label"`
	Price     float64   `json:"price"`
	Public    bool      `json:"public"`
	ImagePath *string   `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
}

type ModifierGroupResponse struct {
	UUID      string             `json:"uuid"`
	Name      string             `json:"name"`
	Modifiers []ModifierResponse `json:"modifiers"`
	MaxQty    *int               `json:"max_quantity"`
	MinQty    *int               `json:"min_quantity"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
}

type ModifierResponse struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// requests

type ItemCategoryUpsertRequest struct {
	UUID           string                 `json:"uuid"`
	Name           string                 `json:"name"`
	ModifierGroups []ModifierGroupRequest `json:"modifier_groups"`
}

type ModifierGroupRequest struct {
	UUID      string            `json:"uuid"`
	Name      string            `json:"name"`
	Modifiers []ModifierRequest `json:"modifiers"`
	MaxQty    *int              `json:"max_quantity"`
	MinQty    *int              `json:"min_quantity"`
}

type ModifierRequest struct {
	UUID  string  `json:"uuid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ItemAndVariantsUpsertRequest struct {
	UUID        string           `json:"uuid"`
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	Label       string           `json:"label"`
	Variants    []VariantRequest `json:"variants"`
	Description string           `json:"description"`
	Public      bool             `json:"public"`
	ImagePath   *string          `json:"image_path"`
}

type VariantRequest struct {
	UUID      string  `json:"uuid"`
	Label     string  `json:"label"`
	Price     float64 `json:"price"`
	Public    bool    `json:"public"`
	ImagePath *string `json:"image_path"`
}
