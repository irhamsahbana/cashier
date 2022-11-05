package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *employeeShiftMongoRepository) ClockOut(ctx context.Context, branchId string, data *domain.EmployeeShiftClockOutData) (*domain.EmployeeShift, int, error) {
	var shift domain.EmployeeShift

	filter := bson.M{
		"$or": bson.A{
			bson.M{
				"$and": []bson.M{
					{"branch_uuid": branchId},
					{"uuid": data.UUID},
					{"deleted_at": nil},
				},
			},
			bson.M{
				"$and": []bson.M{
					{"branch_uuid": branchId},
					{"supporters.uuid": data.UUID},
					{"deleted_at": nil},
				},
			},
		},
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&shift); err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Warn(err)
			return nil, http.StatusNotFound, errors.New("shift not found")
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if shift.EndTime != nil {
		return nil, http.StatusForbidden, errors.New("shift already ended")
	}

	// update main shift
	if shift.UUID == data.UUID { // if as main cashier

		if data.EndCash == nil {
			return nil, http.StatusBadRequest, errors.New("end cash is required")
		}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "end_time", Value: data.EndTime},
				{Key: "end_cash", Value: data.EndCash},
				{Key: "updated_at", Value: time.Now().UnixMicro()},

				{Key: "supporters.$[elem].end_time", Value: data.EndTime},
			}},
		}

		arrayFilters := bson.A{
			bson.M{"elem.end_time": nil},
		}

		_, err := repo.Collection.UpdateOne(ctx, filter, update, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: arrayFilters,
		}))
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else { // if as cashier's supporter
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "supporters.$[elem].end_time", Value: data.EndTime},
				{Key: "supporters.$[elem].updated_at", Value: time.Now().UnixMicro()},
			}},
		}

		arrayFilters := bson.D{
			{Key: "elem.uuid", Value: data.UUID},
		}

		_, err := repo.Collection.UpdateOne(ctx, filter, update, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{arrayFilters},
		}))
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}
	return &shift, http.StatusOK, nil
}
