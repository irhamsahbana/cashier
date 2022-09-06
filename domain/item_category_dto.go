package domain

import "time"

type ItemCategoryResponse struct {
	UUID           string                  `json:"uuid"`
	BranchUUID     string                  `json:"branch_uuid"`
	Name           string                  `json:"name"`
	Items          []ItemResponse          `json:"items,omitempty"`
	ModifierGroups []ModifierGroupResponse `json:"modifier_groups,omitempty"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      *time.Time              `json:"updated_at,omitempty"`
}

type ItemResponse struct {
	UUID        string     `json:"uuid"`
	MainUUID    *string    `json:"main_uuid,omitempty"`
	Name        string     `json:"name"`
	Price       float32    `json:"price"`
	Description *string    `json:"description,omitempty"`
	Label       string     `json:"Label"`
	Public      bool       `json:"public"`
	ImageUrl    *string    `json:"image_url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ModifierGroupResponse struct {
	UUID      string                  `json:"uuid"`
	Name      string                  `json:"name"`
	Modifiers []ModifierResponse      `json:"modifiers"`
	Condition *ModifierGroupCondition `json:"condition"`
	Quantity  *int                    `json:"quantity"`
	Single    bool                    `json:"single"`
	Required  bool                    `json:"required"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
}

type ModifierResponse struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Request item category and modifier group
type ItemCategoryUpsertRequest struct {
	UUID           string                 `json:"uuid"`
	Name           string                 `json:"name"`
	ModifierGroups []ModifierGroupRequest `json:"modifier_groups"`
}

type ModifierGroupRequest struct {
	UUID      string                  `json:"uuid"`
	Name      string                  `json:"name"`
	Modifiers []ModifierRequest       `json:"modifiers"`
	Condition *ModifierGroupCondition `json:"condition"`
	Quantity  *int                    `json:"quantity"`
	Single    bool                    `json:"single"`
	Required  bool                    `json:"required"`
}

type ModifierRequest struct {
	UUID  string  `json:"uuid"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// -- Request item category and modifier group

// Request item and its variants
type ItemCreateRequest struct {
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Variants    []Variant `json:"variants"`
	Description *string   `json:"description"`
	Public      bool      `json:"public"`
	ImagePath   *string   `json:"image_path"`
	CreatedAt   string    `json:"created_at"`
}

type Variant struct {
	UUID      string  `json:"uuid"`
	Label     string  `json:"label"`
	Price     float32 `json:"price"`
	Public    bool    `json:"public"`
	ImagePath *string `json:"image_path"`
}

type ItemUpdateRequest struct {
	MainUUID    string  `json:"main_uuid,omitempty"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description *string `json:"description"`
	Label       string  `json:"label"`
	Public      bool    `json:"public"`
}

// -- Request item and its variants
