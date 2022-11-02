package dto

import "time"

// Responses
type InvoiceResponse struct {
	UUID            string                          `json:"uuid"`
	BranchUUID      string                          `json:"branch_uuid"`
	Customer        *Customer                       `json:"customer"`
	OrderGroups     []OrderGroupResponse            `json:"order_groups"`
	Payments        []InvoicePaymentResponse        `json:"payments"`
	Refunds         []InvoiceRefundResponse         `json:"refunds"`
	CreditContracts []InvoiceCreditContractResponse `json:"credit_contracts"`
	TotalAmount     float64                         `json:"total_amount"`
	TotalTax        float64                         `json:"total_tax"`
	TotalDiscount   float64                         `json:"total_discount"`
	TotalChange     float64                         `json:"total_change"`
	TotalTip        float64                         `json:"total_tip"`
	Note            *string                         `json:"note"`
	CompletedAt     *time.Time                      `json:"completed_at"`
	CreatedAt       time.Time                       `json:"created_at"`
	UpdatedAt       *time.Time                      `json:"updated_at"`
}

type InvoicePaymentResponse struct {
	UUID           string               `json:"uuid"`
	OrderGroupUUID string               `json:"order_group_uuid"`
	PaymentMethod  InvoicePaymentMethod `json:"payment_method"`
	EmployeeShift  EmployeeShift        `json:"employee_shift"`
	Total          float64              `json:"total"`
	Fee            float64              `json:"fee"`
	CreatedAt      time.Time            `json:"created_at"`
}

type InvoiceRefundResponse struct {
	UUID           string        `json:"uuid"`
	OrderGroupUUID string        `json:"order_group_uuid"`
	EmployeeShift  EmployeeShift `json:"employee_shift"`
	Total          float64       `json:"total"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      *time.Time    `json:"updated_at"`
	DeletedAt      *time.Time    `json:"deleted_at"`
}

type InvoiceCreditContractResponse struct {
	UUID      string    `json:"uuid"`
	Number    string    `json:"number"`
	Note      string    `json:"note"`
	DueBy     time.Time `json:"due_by"`
	CreatedAt time.Time `json:"created_at"`
}

// Requests
type InvoiceInsertRequest struct {
	UUID            string                    `json:"uuid"`
	Customer        *Customer                 `json:"customer"`
	OrderGroups     []OrderGroupUpsertRequest `json:"order_groups"`
	Payments        []InvoicePayment          `json:"payments"`
	CreditContracts []InvoiceCreditContract   `json:"credit_contracts"`
	TotalAmount     float64                   `json:"total_amount"`
	TotalTax        float64                   `json:"total_tax"`
	TotalDiscount   float64                   `json:"total_discount"`
	TotalChange     float64                   `json:"total_change"`
	TotalTip        float64                   `json:"total_tip"`
	Note            *string                   `json:"note"`
	CompletedAt     *string                   `json:"completed_at"`
	CreatedAt       string                    `json:"created_at"`
	UpdatedAt       *string                   `json:"updated_at"`
}
