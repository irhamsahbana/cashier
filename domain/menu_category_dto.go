package domain

import "time"

// Response and Request

type MenuCategoryUpsertRequest struct {
	UUID		string			`json:"uuid" bson:"uuid"`
	Name		string			`json:"name" bson:"name"`
	CreatedAt	string			`json:"created_at" bson:"created_at"`
}

type MenuCategoryFindResponse struct {
	UUID		string			`json:"uuid"`
	Name		string			`json:"name"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty"`
}

type MenuCategoryFindAllResponse struct {
	UUID		string					`json:"uuid"`
	Name		string					`json:"name"`
	Menus		[]MenuFindAllResponse	`json:"menus"`
	CreatedAt	time.Time				`json:"created_at"`
	UpdatedAt	*time.Time				`json:"updated_at,omitempty"`
	DeletedAt	*time.Time				`json:"deleted_at,omitempty"`
}

type MenuFindAllResponse struct {
	UUID			string		`json:"uuid"`
	MainUUID		string		`json:"main_uuid,omitempty"`
	Name			string		`json:"name" bson:"name"`
	Price			float32		`json:"price" bson:"price"`
	Description		*string		`json:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	ImageUrl		*string		`json:"image_url" bson:"image_url"`
	CreatedAt		time.Time	`json:"created_at" bson:"created_at"`
	UpdatedAt		*time.Time	`json:"updated_at,omitempty"`
	DeletedAt		*time.Time	`json:"deleted_at,omitempty"`
}

type MenuCategoryUpsertResponse struct {
	UUID		string			`json:"uuid"`
	Name		string			`json:"name"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty"`
}

type MenuCategoryDeleteResponse struct {
	UUID	string			`json:"uuid"`
	Name	string			`json:"name"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
	DeletedAt	time.Time	`json:"deleted_at"`
}

// Menu request

type MenuCreateRequest struct {
	UUID			string		`json:"uuid"`
	Name			string		`json:"name"`
	Price			float32		`json:"price"`
	Description		*string		`json:"description"`
	Label			string		`json:"label"`
	Public			bool		`json:"public"`
	CreatedAt		string		`json:"created_at"`
}

type MenuCreateResponse struct {
	UUID		string		`json:"uuid"`
	Name		string		`json:"name"`
	Price		float32		`json:"price"`
	Description	*string		`json:"description,omitempty"`
	Label		string		`json:"Label"`
	Public		bool		`json:"public"`
	CreatedAt	time.Time	`json:"created_at"`
}

type MenuFindResponse struct {
	UUID		string		`json:"uuid"`
	Name		string		`json:"name"`
	Price		float32		`json:"price"`
	Description	*string		`json:"description,omitempty"`
	Label		string		`json:"Label"`
	Public		bool		`json:"public"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
	DeletedAt	*time.Time	`json:"deleted_at,omitempty"`
}

type MenuDeleteResponse struct {
	UUID		string		`json:"uuid"`
	Name		string		`json:"name"`
	Price		float32		`json:"price"`
	Description	*string		`json:"description,omitempty"`
	Label		string		`json:"Label"`
	Public		bool		`json:"public"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
	DeletedAt	time.Time	`json:"deleted_at"`
}