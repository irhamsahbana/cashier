package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
