package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
