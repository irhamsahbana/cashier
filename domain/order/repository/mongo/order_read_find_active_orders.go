package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *orderRepository) FindActiveOrders(ctx context.Context, branchId string) ([]domain.OrderGroup, int, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"branch_uuid": branchId}},
	}

	cursor, err := repo.CollActive.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	var result []domain.OrderGroup
	if err = cursor.All(ctx, &result); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
