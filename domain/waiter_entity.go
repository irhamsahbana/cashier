package domain

import (
	"context"
)

type Waiter struct {
	UUID       string `json:"uuid" bson:"uuid"`
	BranchUUID string `json:"branch_uuid" bson:"branch_uuid"`
	Name       string `json:"name" bson:"name"`
	LastActive *int64 `json:"last_active,omitempty" bson:"last_active,omitempty"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	UpdatedAt  *int64 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt  *int64 `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

type WaiterUsecaseContract interface {
	UpsertWaiter(ctx context.Context, branchId string, data *WaiterUpsertrequest) (*WaiterResponse, int, error)
	FindWaiter(ctx context.Context, id string, withTrashed bool) (*WaiterResponse, int, error)
	DeleteWaiter(ctx context.Context, id string) (*WaiterResponse, int, error)
}

type WaiterRepositoryContract interface {
	UpsertWaiter(ctx context.Context, data *Waiter) (*Waiter, int, error)
	FindWaiter(ctx context.Context, id string, withTrashed bool) (*Waiter, int, error)
	DeleteWaiter(ctx context.Context, id string) (*Waiter, int, error)
}
