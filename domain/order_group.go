package domain

import (
	"context"
	"lucy/cashier/dto"
)

type OrderGroup struct {
	UUID         string          `bson:"uuid"`
	BranchUUID   string          `bson:"branch_uuid"`
	CreatedBy    string          `bson:"created_by"`
	SpaceUUID    *string         `bson:"space_uuid"`
	Delivery     *Delivery       `bson:"delivery"`
	Queue        *Queue          `bson:"queue"`
	Discounts    []DiscountOrder `bson:"discounts"`
	Orders       []Order         `bson:"orders"`
	Taxes        []TaxOrderGroup `bson:"taxes"`
	CancelReason *string         `bson:"cancel_reason"`
	CreatedAt    int64           `bson:"created_at"`
	UpdatedAt    *int64          `bson:"updated_at"`
	DeletedAt    *int64          `bson:"deleted_at"`
}

type Order struct {
	UUID        string          `bson:"uuid"`
	Item        ItemOrder       `bson:"item"`
	Modifiers   []ModifierOrder `bson:"modifiers"`
	Discounts   []DiscountOrder `bson:"discounts"`
	Waiter      *WaiterOrder    `bson:"waiter"`
	RefundedQty int32           `bson:"refunded_qty"`
	Note        *string         `bson:"note"`
	CreatedAt   int64           `bson:"created_at"`
	UpdatedAt   *int64          `bson:"updated_at"`
	DeletedAt   *int64          `bson:"deleted_at"`
}

type ItemOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Label    string  `bson:"label"`
	Price    float64 `bson:"price"`
	Quantity uint    `bson:"quantity"`
}

type ModifierOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Quantity uint    `bson:"quantity"`
	Price    float64 `bson:"price"`
}

type DiscountOrder struct {
	UUID    string  `bson:"uuid"`
	Name    string  `bson:"name"`
	Fixed   float64 `bson:"fixed"`
	Percent float64 `bson:"percent"`
}

type TaxOrderGroup struct {
	UUID  string  `bson:"uuid"`
	Name  string  `bson:"name"`
	Value float64 `bson:"value"`
}

type WaiterOrder struct {
	UUID       string `bson:"uuid"`
	BranchUUID string `bson:"branch_uuid"`
	Name       string `bson:"name"`
}

type Delivery struct {
	UUID        string   `bson:"uuid"`
	Number      uint     `bson:"number"`
	Partner     string   `bson:"partner"`
	Driver      string   `bson:"driver"`
	Customer    Customer `bson:"customer"`
	ScheduledAt *int64   `bson:"scheduled_at"`
	CreatedAt   int64    `bson:"created_at"`
	UpdatedAt   *int64   `bson:"updated_at"`
	DeletedAt   *int64   `bson:"deleted_at"`
}

type Queue struct {
	UUID        string   `bson:"uuid"`
	Number      uint     `bson:"number"`
	Customer    Customer `bson:"customer"`
	ScheduledAt *int64   `bson:"scheduled_at"`
	CreatedAt   int64    `bson:"created_at"`
	UpdatedAt   *int64   `bson:"updated_at"`
	DeletedAt   *int64   `bson:"deleted_at"`
}

type Customer struct {
	Name    string  `bson:"name"`
	Phone   *string `bson:"phone"`
	Address *string `bson:"address"`
}

type OrderRepositoryContract interface {
	UpsertActiveOrder(ctx context.Context, branchId string, OrderGroup *OrderGroup) (*OrderGroup, int, error)
	FindActiveOrders(ctx context.Context, branchId string) ([]OrderGroup, int, error)
	DeleteActiveOrder(ctx context.Context, branchId, OrderId, reason string) (*OrderGroup, int, error)
}

type OrderUsecaseContract interface {
	UpsertActiveOrder(ctx context.Context, branchId string, req *dto.OrderGroupUpsertRequest) (*dto.OrderGroupResponse, int, error)
	FindActiveOrders(ctx context.Context, branchId string) ([]dto.OrderGroupResponse, int, error)
	DeleteActiveOrder(ctx context.Context, branchId, OrderId string, req *dto.OrderGroupDeleteRequest) (*dto.OrderGroupResponse, int, error)
}
