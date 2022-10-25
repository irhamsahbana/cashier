package mongo

import (
	"context"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *orderRepository) FindInvoiceHistories(ctx context.Context) ([]domain.Invoice, *int64, *int64, int, error) {
	var invoices []domain.Invoice
	var filter bson.M
	var sortType int
	var cursor int64
	var operator string

	var nextCursor *int64
	var prevCursor *int64

	limit := ctx.Value("limit").(int)
	limit = limit + 1
	cursorString := ctx.Value("cursor").(string)
	directionString := ctx.Value("direction").(string)
	sortTypeString := ctx.Value("sort_type").(string)

	// validate cursor
	if cursorTime, err := time.Parse(time.RFC3339, cursorString); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, nil, nil, http.StatusBadRequest, err
	} else {
		cursor = cursorTime.UnixMicro()
	}

	if sortTypeString == "desc" {
		sortType = -1
	} else {
		sortType = 1
	}

	if directionString == "next" {
		operator = "$gt"
	} else {
		operator = "$lt"
	}

	branchId := ctx.Value("branch_uuid").(string)
	fromString := string(ctx.Value("from").(string))
	toString := ctx.Value("to").(string)

	if fromString == "" || toString == "" {
		filter = bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
			},
		}
	} else {
		fromTime, err := time.Parse(time.RFC3339, fromString)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, nil, nil, http.StatusInternalServerError, err
		}

		toTime, err := time.Parse(time.RFC3339, toString)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, nil, nil, http.StatusInternalServerError, err
		}

		filter = bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"created_at": bson.M{"$gte": fromTime.UnixMicro()}},
				{"created_at": bson.M{"$lte": toTime.UnixMicro()}},
			},
		}
	}

	pipeline := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"created_at": sortType}},
		{"$match": bson.M{"created_at": bson.M{operator: cursor}}},
		{"$limit": limit},
	}

	cursorMongo, err := r.CollInvoice.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, nil, nil, http.StatusInternalServerError, err
	}

	if err = cursorMongo.All(ctx, &invoices); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, nil, nil, http.StatusInternalServerError, err
	}
	fmt.Println("====================================")
	fmt.Println("len(Invoices): ", len(invoices))
	fmt.Println("limit: ", limit)
	fmt.Println("cursor: ", cursor)
	fmt.Println("operator: ", operator)
	fmt.Println("sortType: ", sortType)
	fmt.Println("directionString: ", directionString)

	if len(invoices) == limit {
		invoices = invoices[:len(invoices)-1]
		if directionString == "next" {
			nextCursor = &invoices[len(invoices)-1].CreatedAt
			prevCursor = &cursor
		} else {
			prevCursor = &invoices[0].CreatedAt
			nextCursor = &cursor
		}
	}

	return invoices, nextCursor, prevCursor, http.StatusOK, nil
}
