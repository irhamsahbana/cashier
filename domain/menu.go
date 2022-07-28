package domain

import (
	"context"
	"time"
)

type Menu struct {
	ID				uint64		`json:"id" bson:"id"`
	UUID			string		`json:"uuid" bson:"uuid"`
	Name			string		`json:"name" bson:"name" validate:"required"`
	Description		*string		`json:"description" bson:"description"`
	Label			string		`json:"label" bson:"label"`
	Public			bool		`json:"public" bson:"public"`
	ImageUrl		*string		`json:"image_url" bson:"image_url"`
	CreatedAt	time.Time		`json:"created_at" bson:"created_at"`
	UpdatedAt	*time.Time		`json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt	*time.Time		`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type MenuRepositoryContract interface {
	InsertMenu(ctx context.Context, m *Menu) (*Menu, error)
	FindMenu(ctx context.Context, id string) (*Menu, error)
	DeleteMenu(ctx context.Context, id string) (*Menu, error)
	UpdateMenu(ctx context.Context, id string, entity *Menu) (*Menu, error)
}

type MenuUsecaseContract interface {
	CreateMenu(ctx context.Context, entity *Menu) (*Menu, int, error)
	FindMenu(ctx context.Context, id string) (*Menu, int, error)
	DeleteMenu(ctx context.Context, id string) (*Menu, int, error)
	UpdateMenu(ctx context.Context, id string, entity *Menu) (*Menu, int, error)
}