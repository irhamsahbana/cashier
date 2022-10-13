package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *orderRepository) UpsertActiveOrder(ctx context.Context, branchId string, data *domain.OrderGroup) (*domain.OrderGroup, int, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": data.UUID},
		},
	}
	data.BranchUUID = branchId

	countActiveOrderGroup, err := repo.CollActive.CountDocuments(ctx, filter)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if countActiveOrderGroup > 0 {
		return repo.updateActiveOrderGroup(ctx, data)
	} else {
		return repo.createNewOrderGroup(ctx, data)
	}
}

func (repo *orderRepository) updateActiveOrderGroup(ctx context.Context, data *domain.OrderGroup) (*domain.OrderGroup, int, error) {
	var db domain.OrderGroup

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": data.BranchUUID},
			{"uuid": data.UUID},
		},
	}

	// meta data ==================================================
	err := repo.CollActive.FindOne(ctx, filter).Decode(&db)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	// delivery, update delivery data
	if data.Delivery != nil && db.Delivery != nil {
		db.Delivery = data.Delivery
		deliveryUpdatedAt := time.Now().UTC().UnixMicro()
		db.Delivery.UpdatedAt = &deliveryUpdatedAt
	}

	if data.Delivery == nil && db.Delivery != nil {
		deliveryDeletedAt := time.Now().UTC().UnixMicro()
		db.Delivery.DeletedAt = &deliveryDeletedAt
	}

	// queue, update queue data
	if data.Queue != nil && db.Queue != nil {
		db.Queue = data.Queue
		queueUpdatedAt := time.Now().UTC().UnixMicro()
		db.Queue.UpdatedAt = &queueUpdatedAt
	}

	if data.Queue == nil && db.Queue != nil {
		queueDeletedAt := time.Now().UTC().UnixMicro()
		db.Queue.DeletedAt = &queueDeletedAt
	}

	// space, update space data
	if data.SpaceUUID != nil && db.SpaceUUID != nil {
		db.SpaceUUID = data.SpaceUUID
	}

	// -- meta data ============================================================

	// update order
	dbOrderUUID := make([]string, 0)
	for _, order := range db.Orders {
		dbOrderUUID = append(dbOrderUUID, order.UUID)
	}

	dataOrderUUID := make([]string, 0)
	for _, order := range data.Orders {
		dataOrderUUID = append(dataOrderUUID, order.UUID)
	}

	// updated order
	updatedOrders := make([]domain.Order, 0)
	for _, dbOrder := range db.Orders {
		if !helper.Contain(dataOrderUUID, dbOrder.UUID) { // if order not exist in data, delete it
			dbOrderDeletedAt := time.Now().UTC().UnixMicro()
			dbOrder.DeletedAt = &dbOrderDeletedAt
			updatedOrders = append(updatedOrders, dbOrder)
		} else { // if order exist in data, update it,
			var updatedOrder domain.Order
			for _, dataOrder := range data.Orders {
				if dataOrder.UUID == dbOrder.UUID {
					updatedOrder = dataOrder
					updatedOrder.CreatedAt = dbOrder.CreatedAt
					updatedAt := time.Now().UTC().UnixMicro()
					updatedOrder.UpdatedAt = &updatedAt
					break
				}
			}
			updatedOrders = append(updatedOrders, updatedOrder)
		}

		for _, dataOrder := range data.Orders {
			if dataOrder.UUID == dbOrder.UUID {
				dbOrder.Discounts = dataOrder.Discounts
				break
			}
		}
	}

	// new order
	for _, dataOrder := range data.Orders {
		if !helper.Contain(dbOrderUUID, dataOrder.UUID) {
			updatedOrders = append(updatedOrders, dataOrder)
		}
	}

	db.Orders = updatedOrders
	dbUpdatedAt := time.Now().UTC().UnixMicro()
	db.UpdatedAt = &dbUpdatedAt

	// updating order
	err = repo.CollActive.FindOneAndReplace(ctx, filter, db).Decode(&db)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusNotFound, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	err = repo.CollActive.FindOne(ctx, filter).Decode(&db)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusNotFound, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return &db, http.StatusOK, nil
}

func (repo *orderRepository) createNewOrderGroup(ctx context.Context, data *domain.OrderGroup) (*domain.OrderGroup, int, error) {
	var db domain.OrderGroup

	data.CreatedAt = time.Now().UnixMicro()
	data.Taxes = make([]domain.TaxOrderGroup, 0)

	_, err := repo.CollActive.InsertOne(ctx, data)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": data.BranchUUID},
			{"uuid": data.UUID},
		},
	}

	err = repo.CollActive.FindOne(ctx, filter).Decode(&db)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusNotFound, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return &db, http.StatusOK, nil
}
