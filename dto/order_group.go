package dto

import "time"

// response
type OrderGroupResponse struct {
	UUID       string            `json:"uuid"`
	BranchUUID string            `json:"branch_uuid"`
	SpaceUUID  *string           `json:"space_uuid"`
	Delivery   *DeliveryResponse `json:"delivery"`
	Queue      *Queue            `json:"queue"`
	Orders     []OrderResponse   `json:"orders"`
	CreatedBy  string            `json:"created_by"`
	Tax        float32           `json:"tax"`
	Tip        float32           `json:"tip"`
	Pending    *bool             `json:"pending"`
	Completed  bool              `json:"completed"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  *time.Time        `json:"updated_at"`
	DeletedAt  *time.Time        `json:"deleted_at"`
}

type DeliveryResponse struct {
	UUID      string     `json:"uuid"`
	Number    string     `json:"number"`
	Partner   string     `json:"partner"`
	Driver    string     `json:"driver"`
	Customer  Customer   `json:"customer"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type OrderResponse struct {
	UUID        string          `json:"uuid"`
	Item        ItemOrder       `json:"item"`
	Modifiers   []ModifierOrder `json:"modifiers"`
	Waiter      *WaiterOrder    `json:"waiter"`
	RefundedQty int32           `json:"refunded_qty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at"`
}

// requests

type OrderGroupUpsertRequest struct {
	UUID      string    `json:"uuid"`
	SpaceUUID *string   `json:"space_uuid"`
	Delivery  *Delivery `json:"delivery"`
	Queue     *Queue    `json:"queue"`
	Orders    []Order   `json:"orders"`
	CreatedBy string    `json:"created_by"`
	Tax       float32   `json:"tax"`
	Tip       float32   `json:"tip"`
	Pending   *bool     `json:"pending"`
	Completed bool      `json:"completed"`
	CreatedAt string    `json:"created_at"`
}

type Order struct {
	UUID        string          `json:"uuid"`
	Item        ItemOrder       `json:"item"`
	Modifiers   []ModifierOrder `json:"modifiers"`
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
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ModifierOrder struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type WaiterOrder struct {
	UUID       string `json:"uuid"`
	BranchUUID string `json:"branch_uuid"`
	Name       string `json:"name"`
}

type Delivery struct {
	UUID      string   `json:"uuid"`
	Number    string   `json:"number"`
	Partner   string   `json:"partner"`
	Driver    string   `json:"driver"`
	Customer  Customer `json:"customer"`
	CreatedAt string   `json:"created_at"`
}

type Queue struct {
	UUID       string   `json:"uuid"`
	Number     string   `json:"number"`
	Customer   Customer `json:"customer"`
	PromisedAt *string  `json:"promised_at"`
	CreatedAt  string   `json:"created_at"`
}

type Customer struct {
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}
