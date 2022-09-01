package domain

import "time"

// Response and Request
type ItemCategoryResponse struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ItemCategoryUpsertRequest struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}

type ItemCategoryFindAllResponse struct {
	UUID      string                `json:"uuid"`
	Name      string                `json:"name"`
	Items     []ItemFindAllResponse `json:"menus"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at,omitempty"`
	DeletedAt *time.Time            `json:"deleted_at,omitempty"`
}

type ItemFindAllResponse struct {
	UUID        string     `json:"uuid"`
	MainUUID    *string    `json:"main_uuid,omitempty"`
	Name        string     `json:"name" bson:"name"`
	Price       float32    `json:"price" bson:"price"`
	Description *string    `json:"description"`
	Label       string     `json:"label" bson:"label"`
	Public      bool       `json:"public" bson:"public"`
	ImageUrl    *string    `json:"image_url" bson:"image_url"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// Item request

type ItemCreateRequest struct {
	UUID        string  `json:"uuid" bson:"uuid"`
	MainUUID    *string `json:"main_uuid,omitempty" bson:"main_uuid"`
	Name        string  `json:"name" bson:"name"`
	Price       float32 `json:"price" bson:"price"`
	Description *string `json:"description" bson:"description"`
	Label       string  `json:"label" bson:"label"`
	Public      bool    `json:"public" bson:"public"`
	ImageUrl    *string `json:"image_url" bson:"image_url"`
	CreatedAt   string  `json:"created_at" bson:"created_at"`
}

type ItemUpdateRequest struct {
	MainUUID    string  `json:"main_uuid,omitempty" bson:"main_uuid"`
	Name        string  `json:"name" bson:"name"`
	Price       float32 `json:"price" bson:"price"`
	Description *string `json:"description" bson:"description"`
	Label       string  `json:"label" bson:"label"`
	Public      bool    `json:"public" bson:"public"`
}

type ItemResponse struct {
	UUID        string     `json:"uuid"`
	Name        string     `json:"name"`
	Price       float32    `json:"price"`
	Description *string    `json:"description,omitempty"`
	Label       string     `json:"Label"`
	Public      bool       `json:"public"`
	ImageUrl    *string    `json:"image_url,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
