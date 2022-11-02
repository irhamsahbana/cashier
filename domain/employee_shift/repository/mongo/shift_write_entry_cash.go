package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *employeeShiftMongoRepository) InsertEntryCash(ctx context.Context, branchId, shiftId string, data *domain.CashEntry) ([]domain.CashEntry, int, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"$or": bson.A{
				bson.M{
					"$and": []bson.M{
						{"uuid": shiftId},
						{"deleted_at": nil},
					},
				},
				bson.M{
					"$and": []bson.M{
						{"supporters.uuid": shiftId},
						{"deleted_at": nil},
					},
				},
			}},
		},
	}

	result := repo.Collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("employee shift not found")
		}

		logger.Log(logrus.Fields{}).Error(result.Err())
		return nil, http.StatusInternalServerError, result.Err()
	}

	update := bson.M{
		"$set": bson.M{
			"cash_entries": bson.M{
				"$ifNull": bson.A{
					bson.M{
						"$concatArrays": bson.A{
							"$cash_entries",
							bson.A{data},
						},
					},
					bson.A{data},
				},
			},
		},
	}

	_, err := repo.Collection.UpdateOne(ctx, filter, bson.A{update})
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	var shift domain.EmployeeShift

	err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return shift.CashEntries, http.StatusOK, nil
}
