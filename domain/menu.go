package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID				uint64
	UUID			uuid.UUID
	MenuID			uint64
	Name			string
	Description		string
	Label			string
	Public			bool
	ImageUrl		*string
	CreatedAt		time.Time
	UpdatedAt		*time.Time
	DeletedAt		*time.Time
}

type MenuRepositoryContract interface {
	InsertOne(ctx context.Context, m *Menu) (*Menu, error)
	UpdateOne(ctx context.Context, id string) (*Menu, error)
	DeleteOne(ctx context.Context, id string) (*Menu, error)
}

type MenuUsecaseContract interface {
	InsertOne(ctx context.Context, m *Menu) (*Menu, error)
	UpdateOne(ctx context.Context, id string) (*Menu, error)
	DeleteOne(ctx context.Context, id string) (*Menu, error)
}