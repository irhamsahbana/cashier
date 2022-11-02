package mongo

import (
	"context"
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type employeeShiftMongoRepository struct {
	DB               mongo.Database
	Collection       mongo.Collection
	CollInvoice      mongo.Collection
	CollItemCategory mongo.Collection
}

func NewEmployeeShiftMongoRepository(DB mongo.Database) domain.EmployeeShiftRepositoryContract {
	return &employeeShiftMongoRepository{
		DB:               DB,
		Collection:       *DB.Collection("employee_shifts"),
		CollInvoice:      *DB.Collection("invoices"),
		CollItemCategory: *DB.Collection("item_categories"),
	}
}

func (repo *employeeShiftMongoRepository) FindShiftByStartTime(ctx context.Context, branchId, userId string, startTime int64) (*domain.EmployeeShift, int, error) {
	panic("implement me")
}
