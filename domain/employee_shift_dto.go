package domain

import "time"

type EmployeeShiftResponse struct {
	UUID       string                   `json:"uuid"`
	BranchUUID string                   `json:"branch_uuid"`
	UserUUID   string                   `json:"user_uuid"`
	StartTime  time.Time                `json:"start_time"`
	StartCash  float32                  `json:"start_cash"`
	EndTime    *time.Time               `json:"end_time"`
	EndCash    *float32                 `json:"end_cash"`
	Suporters  []EmployeeShiftSupporter `json:"supporters"`
	CreatedAt  int64                    `json:"created_at"`
	UpdatedAt  *int64                   `json:"updated_at"`
	DeletedAt  *int64                   `json:"deleted_at"`
}

type EmployeeShiftSupporterResponse struct {
	UUID      string     `json:"uuid"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type EmployeeShiftClockInRequest struct {
	UUID           string   `json:"uuid"`
	SupportingUUID *string  `json:"supporting_uuid"`
	StartTime      string   `json:"start_time"`
	StartCash      *float32 `json:"start_cash"`
}
