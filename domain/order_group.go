package domain

import (
	"context"
	"lucy/cashier/dto"
)

type OrderGroup struct {
	UUID       string    `bson:"uuid"`
	BranchUUID string    `bson:"branch_uuid"`
	SpaceUUID  *string   `bson:"space_uuid"`
	Delivery   *Delivery `bson:"delivery"`
	Queue      *Queue    `bson:"queue"`
	Orders     []Order   `bson:"orders"`
	CreatedBy  string    `bson:"created_by"`
	Tax        float32   `bson:"tax"`
	Tip        float32   `bson:"tip"`
	Pending    *bool     `bson:"pending"`
	Completed  bool      `bson:"completed"`
	CreatedAt  int64     `bson:"created_at"`
	UpdatedAt  *int64    `bson:"updated_at"`
	DeletedAt  *int64    `bson:"deleted_at"`
}

type Order struct {
	UUID        string          `bson:"uuid"`
	Item        ItemOrder       `bson:"item"`
	Modifiers   []ModifierOrder `bson:"modifiers"`
	Waiter      *WaiterOrder    `bson:"waiter"`
	RefundedQty int32           `bson:"refunded_qty"`
	CreatedAt   int64           `bson:"created_at"`
	UpdatedAt   *int64          `bson:"updated_at"`
	DeletedAt   *int64          `bson:"deleted_at"`
}

type ItemOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Label    string  `bson:"label"`
	Price    float32 `bson:"price"`
	Quantity int     `bson:"quantity"`
}

type ModifierOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Quantity int     `bson:"quantity"`
	Price    float32 `bson:"price"`
}

type WaiterOrder struct {
	UUID       string `bson:"uuid"`
	BranchUUID string `bson:"branch_uuid"`
	Name       string `bson:"name"`
}

type Delivery struct {
	UUID      string   `bson:"uuid"`
	Number    string   `bson:"number"`
	Partner   string   `bson:"partner"`
	Driver    string   `bson:"driver"`
	Customer  Customer `bson:"customer"`
	CreatedAt int64    `bson:"created_at"`
	UpdatedAt *int64   `bson:"updated_at"`
	DeletedAt *int64   `bson:"deleted_at"`
}

type Queue struct {
	UUID       string   `bson:"uuid"`
	Number     string   `bson:"number"`
	Customer   Customer `bson:"customer"`
	PromisedAt *int64   `bson:"promised_at"`
	CreatedAt  int64    `bson:"created_at"`
	UpdatedAt  *int64   `bson:"updated_at"`
	DeletedAt  *int64   `bson:"deleted_at"`
}

type Customer struct {
	Name    string  `bson:"name"`
	Phone   *string `bson:"phone"`
	Address *string `bson:"address"`
}

type OrderRepositoryContract interface {
	UpsertActiveOrder(ctx context.Context, branchId string, OrderGroup *OrderGroup) (*OrderGroup, int, error)
	FindActiveOrders(ctx context.Context, branchId string) ([]OrderGroup, int, error)
}

type OrderUsecaseContract interface {
	UpsertActiveOrder(ctx context.Context, branchId string, req *dto.OrderGroupUpsertRequest) (*dto.OrderGroupResponse, int, error)
	FindActiveOrders(ctx context.Context, branchId string) ([]dto.OrderGroupResponse, int, error)
}
