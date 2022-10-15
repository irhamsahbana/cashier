package domain

import "context"

type EmployeeShift struct {
	UUID       string                   `bson:"uuid"`
	BranchUUID string                   `bson:"branch_uuid"`
	UserUUID   string                   `bson:"user_uuid"`
	StartTime  int64                    `bson:"start_time"`
	StartCash  float64                  `bson:"start_cash"`
	EndTime    *int64                   `bson:"end_time"`
	EndCash    *float64                 `bson:"end_cash"`
	Supporters []EmployeeShiftSupporter `bson:"supporters"`
	CreatedAt  int64                    `bson:"created_at"`
	UpdatedAt  *int64                   `bson:"updated_at"`
	DeletedAt  *int64                   `bson:"deleted_at"`
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
	ClockIn(ctx context.Context, branchId string, req *EmployeeShiftClockInRequest) (*EmployeeShiftResponse, int, error)
	ClockOut(ctx context.Context, branchId string, req *EmployeeShiftClockOutRequest) (*EmployeeShiftResponse, int, error)

	History(ctx context.Context, branchId string, limit, page int) ([]EmployeeShiftResponse, int, error)
	Active(ctx context.Context, branchId string) ([]EmployeeShiftResponse, int, error)
}

type EmployeeShiftRepositoryContract interface {
	ClockIn(ctx context.Context, branchId string, data *EmployeeShiftClockInData) (*EmployeeShift, int, error)
	ClockOut(ctx context.Context, branchId string, data *EmployeeShiftClockOutData) (*EmployeeShift, int, error)

	History(ctx context.Context, branchId string, limit, page int) ([]EmployeeShift, int, error)
	Active(ctx context.Context, branchId string) ([]EmployeeShift, int, error)
}
