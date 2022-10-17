package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
}

func (repo *employeeShiftMongoRepository) History(ctx context.Context, branchId string, limit, page int) ([]domain.EmployeeShift, int, error) {
	var shifts []domain.EmployeeShift

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"end_time": bson.M{"$ne": nil}},
			// {"end_cash": bson.M{"$ne": nil}},
			{"deleted_at": nil},
		},
	}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(limit * page))

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = cursor.All(ctx, &shifts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return shifts, http.StatusOK, nil
}

// not yet fully implemented
// this function need search to order group that hasn't been closed yet
func (repo *employeeShiftMongoRepository) Active(ctx context.Context, branchId string) ([]domain.EmployeeShift, int, error) {
	var shifts []domain.EmployeeShift

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"end_time": nil},
			{"deleted_at": nil},
		},
	}

	opts := options.Find().SetSort(bson.M{"start_time": -1})

	cursor, err := repo.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	err = cursor.All(ctx, &shifts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return shifts, http.StatusOK, nil
}
