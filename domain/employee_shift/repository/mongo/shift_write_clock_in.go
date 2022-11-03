package mongo

import (
	"context"
	"errors"
	"lucy/cashier/bootstrap"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	// "github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *employeeShiftMongoRepository) ClockIn(ctx context.Context, branchId string, data *domain.EmployeeShiftClockInData) (*domain.EmployeeShift, int, error) {
	// startTime := time.UnixMicro(data.StartTime).UTC()

	// beginingOfDay := carbon.Time2Carbon(startTime).StartOfDay().Carbon2Time().UnixMicro()
	// endOfDay := carbon.Time2Carbon(startTime).EndOfDay().Carbon2Time().UnixMicro()

	var shift domain.EmployeeShift

	// if as main cashier
	if data.StartCash != nil {
		// var mainShift domain.EmployeeShift

		// // check if there is shift in the same day, if there is, return error
		// filter := bson.M{
		// 	"$and": []bson.M{
		// 		{"branch_uuid": branchId},
		// 		{"user_uuid": data.UserUUID},
		// 		{"start_time": bson.M{"$gte": beginingOfDay}},
		// 		{"start_time": bson.M{"$lte": endOfDay}},
		// 		{"end_time": nil},
		// 		{"deleted_at": nil},
		// 	},
		// }

		// // check existing shift
		// err := repo.Collection.FindOne(ctx, filter).Decode(&mainShift)
		// if err != nil && err != mongo.ErrNoDocuments {
		// 	bootstrap.App.Log.Warn(err)
		// 	return nil, http.StatusInternalServerError, err
		// }

		// // if there is shift, return error
		// if mainShift.UUID != "" {
		// 	bootstrap.App.Log.Warn("there is shift in the same day")
		// 	return nil, http.StatusConflict, errors.New("there is shift in the same day")
		// }

		// create new shift
		doc := bson.D{
			{Key: "uuid", Value: data.UUID},
			{Key: "branch_uuid", Value: branchId},
			{Key: "user_uuid", Value: data.UserUUID},
			{Key: "start_time", Value: data.StartTime},
			{Key: "start_cash", Value: data.StartCash},
			{Key: "end_time", Value: nil},
			{Key: "end_cash", Value: nil},
			{Key: "supporters", Value: []domain.EmployeeShiftSupporter{}},
			{Key: "cash_entries", Value: []domain.CashEntry{}},
			{Key: "created_at", Value: time.Now().UnixMicro()},
			{Key: "updated_at", Value: nil},
			{Key: "deleted_at", Value: nil},
		}

		_, err := repo.Collection.InsertOne(ctx, doc)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}

		filter := bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"user_uuid": data.UserUUID},
				{"start_time": data.StartTime},
			},
		}

		err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, http.StatusNotFound, errors.New("employee shift not found")
			}

			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}

		return &shift, http.StatusOK, nil
	} else { // if as cashier's supporter
		filter := bson.D{
			{Key: "$and", Value: []bson.M{
				{"branch_uuid": branchId},
				{"uuid": data.SupportingUUID},
				{"deleted_at": nil},
			}},
		}

		err := repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				logger.Log(logrus.Fields{}).Warn(err)
				return nil, http.StatusNotFound, errors.New("main shift not found")
			}

			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}

		// deny if main shift id same as supporting shift id
		if shift.UUID == data.UUID {
			return nil, http.StatusConflict, errors.New("main shift id same as supporting shift id")
		}

		// check if there is supporter shift in the same day and same main shift, if there is, return error
		filter = bson.D{
			{Key: "$and", Value: []bson.M{
				{"branch_uuid": branchId},
				{"uuid": data.SupportingUUID},
				{"supporters.user_uuid": data.UserUUID},
				{"supporters.end_time": nil},
			}},
		}

		err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil && err != mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Warn(err)
			return nil, http.StatusInternalServerError, err
		}

		if shift.EndTime != nil {
			return nil, http.StatusForbidden, errors.New("main shift already ended")
		}

		for _, s := range shift.Supporters {
			if s.UserUUID == data.UserUUID && s.EndTime == nil {
				bootstrap.App.Log.WithFields(map[string]interface{}{
					"branch_uuid":     branchId,
					"main_shift_uuid": shift.UUID,
					"user_uuid":       data.UserUUID,
				}).Warn("there is supporter shift in the same day and same main shift")

				return nil, http.StatusConflict, errors.New("there is supporter shift in the same day")
			}
		}

		// create new supporter shift
		filter = bson.D{
			{Key: "$and", Value: []bson.M{
				{"branch_uuid": branchId},
				{"uuid": data.SupportingUUID},
			}},
		}

		doc := bson.D{
			{Key: "uuid", Value: data.UUID},
			{Key: "user_uuid", Value: data.UserUUID},
			{Key: "start_time", Value: data.StartTime},
			{Key: "end_time", Value: nil},
			{Key: "created_at", Value: time.Now().UnixMicro()},
			{Key: "updated_at", Value: nil},
			{Key: "deleted_at", Value: nil},
		}

		update := bson.M{
			"$push": bson.M{
				"supporters": doc,
			},
		}

		_, err = repo.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		filter = bson.D{
			{Key: "$and", Value: []bson.M{
				{"branch_uuid": branchId},
				{"uuid": data.SupportingUUID},
			}},
		}

		err = repo.Collection.FindOne(ctx, filter).Decode(&shift)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusInternalServerError, err
		}

		return &shift, http.StatusOK, nil
	}
}
