package dto

import "time"

// response
type OrderGroupResponse struct {
	UUID         string            `json:"uuid"`
	BranchUUID   string            `json:"branch_uuid"`
	SpaceUUID    *string           `json:"space_uuid"`
	Delivery     *DeliveryResponse `json:"delivery"`
	Queue        *Queue            `json:"queue"`
	CreatedBy    string            `json:"created_by"`
	Orders       []OrderResponse   `json:"orders"`
	Taxes        []TaxOrderGroup   `json:"taxes"`
	CancelReason *string           `json:"cancel_reason,omitempty"`
	Pending      *bool             `json:"pending"`
	Completed    bool              `json:"completed"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    *time.Time        `json:"updated_at"`
	DeletedAt    *time.Time        `json:"deleted_at"`
}

type DeliveryResponse struct {
	UUID        string     `json:"uuid"`
	Number      uint       `json:"number"`
	Partner     string     `json:"partner"`
	Driver      string     `json:"driver"`
	Customer    Customer   `json:"customer"`
	ScheduledAt *time.Time `json:"scheduled_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type OrderResponse struct {
	UUID        string          `json:"uuid"`
	Item        ItemOrder       `json:"item"`
	Modifiers   []ModifierOrder `json:"modifiers"`
	Discounts   []DiscountOrder `json:"discounts"`
	Waiter      *WaiterOrder    `json:"waiter"`
	RefundedQty int32           `json:"refunded_qty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at"`
}

type TaxOrderGroup struct {
	UUID  string  `json:"uuid"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// requests
type OrderGroupUpsertRequest struct {
	UUID      string          `json:"uuid"`
	SpaceUUID *string         `json:"space_uuid"`
	Delivery  *Delivery       `json:"delivery"`
	Queue     *Queue          `json:"queue"`
	Orders    []Order         `json:"orders"`
	Taxes     []TaxOrderGroup `json:"taxes"`
	CreatedBy string          `json:"created_by"`
	Pending   *bool           `json:"pending"`
	Completed bool            `json:"completed"`
	CreatedAt string          `json:"created_at"`
}

type OrderGroupDeleteRequest struct {
	CancelReason string `json:"cancel_reason"`
}

// -- requests

type Order struct {
	UUID        string          `json:"uuid"`
	Item        ItemOrder       `json:"item"`
	Modifiers   []ModifierOrder `json:"modifiers"`
	Discounts   []DiscountOrder `json:"discounts"`
	Waiter      *WaiterOrder    `json:"waiter"`
	RefundedQty int32           `json:"refunded_qty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   *string         `json:"updated_at"`
	DeletedAt   *string         `json:"deleted_at"`
}

type ItemOrder struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Label    string  `json:"label"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type ModifierOrder struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

type DiscountOrder struct {
	UUID    string  `json:"uuid"`
	Name    string  `json:"name"`
	Fixed   float64 `json:"fixed"`
	Percent float64 `json:"percent"`
}

type WaiterOrder struct {
	UUID       string `json:"uuid"`
	BranchUUID string `json:"branch_uuid"`
	Name       string `json:"name"`
}

type Delivery struct {
	UUID        string   `json:"uuid"`
	Number      uint     `json:"number"`
	Partner     string   `json:"partner"`
	Driver      string   `json:"driver"`
	Customer    Customer `json:"customer"`
	ScheduledAt *string  `json:"scheduled_at"`
	CreatedAt   string   `json:"created_at"`
}

type Queue struct {
	UUID        string   `json:"uuid"`
	Number      uint     `json:"number"`
	Customer    Customer `json:"customer"`
	ScheduledAt *string  `json:"scheduled_at"`
	CreatedAt   string   `json:"created_at"`
}

type Customer struct {
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}
