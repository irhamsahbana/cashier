package dto

import "time"

type EmployeeShiftResponse struct {
	UUID        string                           `json:"uuid"`
	BranchUUID  string                           `json:"branch_uuid"`
	UserUUID    string                           `json:"user_uuid"`
	StartTime   time.Time                        `json:"start_time"`
	StartCash   float64                          `json:"start_cash"`
	EndTime     *time.Time                       `json:"end_time"`
	EndCash     *float64                         `json:"end_cash"`
	Supporters  []EmployeeShiftSupporterResponse `json:"supporters"`
	CashEntries []CashEntryResponse              `json:"cash_entries"`
	Summary     EmployeeShiftSummaryResponse     `json:"summary"`
	CreatedAt   time.Time                        `json:"created_at"`
	UpdatedAt   *time.Time                       `json:"updated_at"`
	DeletedAt   *time.Time                       `json:"deleted_at"`
}

type EmployeeShiftSupporterResponse struct {
	UUID      string     `json:"uuid"`
	UserUUID  string     `json:"user_uuid"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type EmployeeShiftSummaryResponse struct {
	TotalRefunds int64                         `json:"total_refunds"`
	Payments     []EmployeeShiftSummaryPayment `json:"payments"`
	Orders       []EmployeeShiftSummaryOrder   `json:"orders"`
}

type EmployeeShiftSummaryPayment struct {
	UUID  string  `json:"uuid"`
	Qty   int64   `json:"quantity"`
	Total float64 `json:"total"`
}

type EmployeeShiftSummaryOrder struct {
	UUID        string  `json:"uuid" bson:"uuid"`
	Category    string  `json:"category" bson:"category"`
	Name        string  `json:"name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
	Qty         int64   `json:"quantity" bson:"quantity"`
	RefundedQty int64   `json:"refunded_quantity" bson:"refunded_quantity"`
}

// requests

type EmployeeShiftClockInRequest struct {
	UUID           string   `json:"uuid"`
	UserUUID       string   `json:"user_uuid"`
	SupportingUUID *string  `json:"supporting_uuid"`
	StartTime      string   `json:"start_time"`
	StartCash      *float64 `json:"start_cash"`
}

type EmployeeShiftClockOutRequest struct {
	UUID    string   `json:"uuid"`
	EndTime string   `json:"end_time"`
	EndCash *float64 `json:"end_cash"`
}

type CashEntryInsertRequest struct {
	Username    string  `json:"username"`
	Description string  `json:"description"`
	Expense     bool    `json:"expense"`
	Value       float64 `json:"value"`
	CreatedAt   string  `json:"created_at"`
}

type CashEntryResponse struct {
	Username    string    `json:"username"`
	Description string    `json:"description"`
	Expense     bool      `json:"expense"`
	Value       float64   `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
}
