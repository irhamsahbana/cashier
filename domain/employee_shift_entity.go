package domain

import "context"

type EmployeeShift struct {
	UUID       string                   `bson:"uuid"`
	BranchUUID string                   `bson:"branch_uuid"`
	UserUUID   string                   `bson:"user_uuid"`
	StartTime  int64                    `bson:"start_time"`
	StartCash  float32                  `bson:"start_cash"`
	EndTime    *int64                   `bson:"end_time"`
	EndCash    *float32                 `bson:"end_cash"`
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

type EmployeeShiftClockInData struct {
	UUID           string   `bson:"uuid"`
	UserUUID       string   `bson:"user_uuid"`
	SupportingUUID *string  `bson:"supporting_uuid"`
	StartTime      int64    `bson:"start_time"`
	StartCash      *float32 `bson:"start_cash"`
}

type EmployeeShiftUsecaseContract interface {
	ClockIn(ctx context.Context, branchId string, req *EmployeeShiftClockInRequest) (*EmployeeShiftResponse, int, error)
}

type EmployeeShiftRepositoryContract interface {
	ClockIn(ctx context.Context, branchId string, data *EmployeeShiftClockInData) (*EmployeeShift, int, error)

	FindShiftByStartTime(ctx context.Context, branchId, userId string, startTime int64) (*EmployeeShift, int, error)
}
