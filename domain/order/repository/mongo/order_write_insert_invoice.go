package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *orderRepository) InsertInvoice(c context.Context, branchId string, data *domain.Invoice) (*domain.Invoice, int, error) {
	var result domain.Invoice

	if _, err := r.CollInvoice.InsertOne(c, data); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	filter := bson.M{"uuid": data.UUID}
	if err := r.CollInvoice.FindOne(c, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusNotFound, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return &result, http.StatusOK, nil
}
