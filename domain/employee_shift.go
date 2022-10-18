package domain

import (
	"context"
	"lucy/cashier/dto"
)

type EmployeeShift struct {
	UUID        string                   `bson:"uuid"`
	BranchUUID  string                   `bson:"branch_uuid"`
	UserUUID    string                   `bson:"user_uuid"`
	StartTime   int64                    `bson:"start_time"`
	StartCash   float64                  `bson:"start_cash"`
	EndTime     *int64                   `bson:"end_time"`
	EndCash     *float64                 `bson:"end_cash"`
	Supporters  []EmployeeShiftSupporter `bson:"supporters"`
	CashEntries []CashEntry              `bson:"cash_entries"`
	CreatedAt   int64                    `bson:"created_at"`
	UpdatedAt   *int64                   `bson:"updated_at"`
	DeletedAt   *int64                   `bson:"deleted_at"`
}

type EmployeeShiftSupporter struct {
	UUID      string `bson:"uuid"`
	UserUUID  string `bson:"user_uuid"`
	StartTime int64  `bson:"start_time"`
	EndTime   *int64 `bson:"end_time"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt *int64 `bson:"updated_at"`
	DeletedAt *int64 `bson:"deleted_at"`
}

type CashEntry struct {
	Username    string  `bson:"username"`
	Description string  `bson:"description"`
	Expense     bool    `bson:"expense"`
	Value       float64 `bson:"value"`
	CreatedAt   int64   `bson:"created_at"`
}

// data that are used to pass in to the repository
type EmployeeShiftClockInData struct {
	UUID           string   `bson:"uuid"`
	UserUUID       string   `bson:"user_uuid"`
	SupportingUUID *string  `bson:"supporting_uuid"`
	StartTime      int64    `bson:"start_time"`
	StartCash      *float64 `bson:"start_cash"`
}

type EmployeeShiftClockOutData struct {
	UUID    string   `bson:"uuid"`
	EndTime int64    `bson:"end_time"`
	EndCash *float64 `bson:"end_cash"`
}

// -- data that are used to pass in to the repository

type EmployeeShiftUsecaseContract interface {
	ClockIn(c context.Context, branchId string, req *dto.EmployeeShiftClockInRequest) (*dto.EmployeeShiftResponse, int, error)
	ClockOut(c context.Context, branchId string, req *dto.EmployeeShiftClockOutRequest) (*dto.EmployeeShiftResponse, int, error)
	History(c context.Context, branchId string, limit, page int) ([]dto.EmployeeShiftResponse, int, error)
	Active(c context.Context, branchId string) ([]dto.EmployeeShiftResponse, int, error)

	InsertEntryCash(c context.Context, branchId, shiftId string, req *dto.CashEntryInsertRequest) ([]dto.CashEntryResponse, int, error)
}

type EmployeeShiftRepositoryContract interface {
	ClockIn(c context.Context, branchId string, data *EmployeeShiftClockInData) (*EmployeeShift, int, error)
	ClockOut(c context.Context, branchId string, data *EmployeeShiftClockOutData) (*EmployeeShift, int, error)

	History(c context.Context, branchId string, limit, page int) ([]EmployeeShift, int, error)
	Active(c context.Context, branchId string) ([]EmployeeShift, int, error)

	InsertEntryCash(c context.Context, branchId, shiftId string, data *CashEntry) ([]CashEntry, int, error)
}
