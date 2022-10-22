package dto

import "time"

// response
type InvoiceResponse struct {
	UUID            string                  `json:"uuid"`
	Customer        *Customer               `json:"customer"`
	OrderGroups     []OrderGroupResponse    `json:"order_groups"`
	Payments        []InvoicePayment        `json:"payments"`
	Refunds         []InvoiceRefund         `json:"refunds"`
	CreditContracts []InvoiceCreditContract `json:"credit_contracts"`
	TotalAmount     float64                 `json:"total_amount"`
	TotalTax        float64                 `json:"total_tax"`
	TotalDiscount   float64                 `json:"total_discount"`
	TotalChange     float64                 `json:"total_change"`
	TotalTip        float64                 `json:"total_tip"`
	Note            *string                 `json:"note"`
	CompletedAt     *time.Time              `json:"completed_at"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       *time.Time              `json:"updated_at"`
}

type OrderGroupResponse struct {
	UUID         string            `json:"uuid"`
	BranchUUID   string            `json:"branch_uuid"`
	CreatedBy    string            `json:"created_by"`
	SpaceUUID    *string           `json:"space_uuid"`
	Delivery     *DeliveryResponse `json:"delivery"`
	Queue        *Queue            `json:"queue"`
	Orders       []OrderResponse   `json:"orders"`
	Taxes        []TaxOrderGroup   `json:"taxes"`
	CancelReason *string           `json:"cancel_reason,omitempty"`
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
	Note        *string         `json:"note"`
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
type InvoiceInsertRequest struct {
	UUID        string                    `json:"uuid"`
	Customer    *Customer                 `json:"customer"`
	OrderGroups []OrderGroupUpsertRequest `json:"order_groups"`
	Payments    []InvoicePayment          `json:"payments"`
	// Refunds         []InvoiceRefund           `json:"refunds"`
	CreditContracts []InvoiceCreditContract `json:"credit_contracts"`
	TotalAmount     float64                 `json:"total_amount"`
	TotalTax        float64                 `json:"total_tax"`
	TotalDiscount   float64                 `json:"total_discount"`
	TotalChange     float64                 `json:"total_change"`
	TotalTip        float64                 `json:"total_tip"`
	Note            *string                 `json:"note"`
	CompletedAt     *string                 `json:"completed_at"`
	CreatedAt       string                  `json:"created_at"`
	UpdatedAt       *string                 `json:"updated_at"`
}

type InvoicePayment struct {
	UUID           string               `json:"uuid"`
	OrderGroupUUID string               `json:"order_group_uuid"`
	PaymentMethod  InvoicePaymentMethod `json:"payment_method"`
	EmployeeShift  EmployeeShift        `json:"employee_shift"`
	Total          float64              `json:"total"`
	Fee            float64              `json:"fee"`
	CreatedAt      string               `json:"created_at"`
}

type InvoicePaymentMethod struct {
	PaymentMethodUUID string                  `json:"payment_method_uuid"`
	Group             string                  `json:"group"`
	Name              string                  `json:"name"`
	Fee               InvoicePaymentMethodFee `json:"fee"`
}

type InvoicePaymentMethodFee struct {
	Percent float64 `json:"percent"`
	Fixed   float64 `json:"fixed"`
}

type InvoiceRefund struct {
	UUID           string        `json:"uuid"`
	OrderGroupUUID string        `json:"order_group_uuid"`
	EmployeeShift  EmployeeShift `json:"employee_shift"`
	Total          float64       `json:"total"`
	CreatedAt      string        `json:"created_at"`
}

type InvoiceCreditContract struct {
	UUID      string `json:"uuid"`
	Number    string `json:"number"`
	Note      string `json:"note"`
	DueBy     string `json:"due_by"`
	CreatedAt string `json:"created_at"`
}

type EmployeeShift struct {
	EmployeeShiftUUID string `json:"employee_shift_uuid"`
	UserUUID          string `json:"user_uuid"`
	Name              string `json:"name"`
}

type OrderGroupUpsertRequest struct {
	UUID      string          `json:"uuid"`
	SpaceUUID *string         `json:"space_uuid"`
	Delivery  *Delivery       `json:"delivery"`
	Queue     *Queue          `json:"queue"`
	Orders    []Order         `json:"orders"`
	Taxes     []TaxOrderGroup `json:"taxes"`
	CreatedBy string          `json:"created_by"`
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
	Note        *string         `json:"note"`
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
