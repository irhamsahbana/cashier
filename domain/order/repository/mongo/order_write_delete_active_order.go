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
)

func (repo *orderRepository) DeleteActiveOrder(ctx context.Context, branchId, orderId, reason string) (*domain.OrderGroup, int, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": orderId},
		},
	}

	var data domain.OrderGroup

	err := repo.CollActive.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Warn(err)
			return nil, http.StatusNotFound, errors.New("order group not found")
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	data.CancelReason = &reason
	deletedAt := time.Now().UTC().UnixMicro()
	data.DeletedAt = &deletedAt

	// save to deleted_order_groups
	_, err = repo.CollDeleted.InsertOne(ctx, data)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	// delete from active_order_groups
	_, err = repo.CollActive.DeleteOne(ctx, filter)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return &data, http.StatusOK, nil
}
