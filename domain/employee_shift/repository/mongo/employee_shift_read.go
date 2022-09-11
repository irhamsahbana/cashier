package mongo

import (
	"context"
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type employeeShiftMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewEmployeeShiftMongoRepository(DB mongo.Database) domain.EmployeeShiftRepositoryContract {
	return &employeeShiftMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("employee_shifts"),
	}
}

func (repo *employeeShiftMongoRepository) FindShiftByStartTime(ctx context.Context, branchId, userId string, startTime int64) (*domain.EmployeeShift, int, error) {
	panic("implement me")

	// t := time.UnixMicro(startTime).UTC()

	// beginingOfDay := carbon.Time2Carbon(t).StartOfDay().Carbon2Time().UnixMicro()
	// endOfDay := carbon.Time2Carbon(t).EndOfDay().Carbon2Time().UnixMicro()

	// var shift domain.EmployeeShift
}
