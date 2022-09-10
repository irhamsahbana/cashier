package domain

type EmployeeShift struct {
	UUID       string                   `bson:"uuid"`
	BranchUUID string                   `bson:"branch_uuid"`
	UserUUID   string                   `bson:"user_uuid"`
	startTime  int64                    `bson:"start_time"`
	StartCash  float32                  `bson:"start_cash"`
	EndTime    *int64                   `bson:"end_time"`
	EndCash    *float32                 `bson:"end_cash"`
	Suporters  []EmployeeShiftSupporter `bson:"supporters"`
	createdAt  int64                    `bson:"created_at"`
	UpdatedAt  *int64                   `bson:"updated_at"`
	DeletedAt  *int64                   `bson:"deleted_at"`
}

type EmployeeShiftSupporter struct {
	UUID      string `bson:"uuid"`
	StartTime int64  `bson:"start_time"`
	EndTime   *int64 `bson:"end_time"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt *int64 `bson:"updated_at"`
	DeletedAt *int64 `bson:"deleted_at"`
}

type EmployeeShiftUsecaseContract interface {
}

type EmployeeShiftRepositoryContract interface {
}
