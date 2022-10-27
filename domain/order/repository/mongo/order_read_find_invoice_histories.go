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
	if cursorTime, err := time.Parse(time.RFC3339, cursorString); err != nil && cursorString != "" {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, nil, nil, http.StatusUnprocessableEntity, errors.New("invalid cursor")
	} else {
		cursor = cursorTime.UnixMicro()
	}

	if sortTypeString == "desc" {
		if directionString == "next" {
			sortType = -1
			operator = "$lt"
		} else {
			sortType = 1
			operator = "$gt"
		}
	}

	if sortTypeString == "asc" {
		if directionString == "next" {
			sortType = 1
			operator = "$gt"
		} else {
			sortType = -1
			operator = "$lt"
		}
	}

	branchId := ctx.Value("branch_uuid").(string)
	fromString := string(ctx.Value("from").(string))
	toString := ctx.Value("to").(string)

	// filter range date
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
			return nil, nil, nil, http.StatusInternalServerError, errors.New("invalid 'from' date")
		}

		toTime, err := time.Parse(time.RFC3339, toString)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, nil, nil, http.StatusInternalServerError, errors.New("invalid 'to' date")
		}

		filter = bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"created_at": bson.M{"$gte": fromTime.UnixMicro()}},
				{"created_at": bson.M{"$lte": toTime.UnixMicro()}},
			},
		}
	}
	// -- filter range date

	pipeline := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"created_at": sortType}},
		{"$match": bson.M{"created_at": bson.M{operator: cursor}}},
		{"$limit": limit},
	}

	// if no cursor, then get the first page
	if cursorString == "" {
		pipeline = []bson.M{
			{"$match": filter},
			{"$sort": bson.M{"created_at": sortType}},
			{"$limit": limit},
		}
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

	if directionString == "next" {
		// check if there is next page
		if len(invoices) != limit {
			nextCursor = nil
		} else {
			nextCursor = &invoices[limit-2].CreatedAt
			invoices = invoices[:limit-1]
		}

		// check if there is prev page
		if cursorString == "" {
			prevCursor = nil
		} else {
			prevCursor = &invoices[0].CreatedAt
		}
	} else if directionString == "prev" {
		// check if there is prev page
		if len(invoices) != limit-1 {
			prevCursor = nil
		} else {
			prevCursor = &invoices[0].CreatedAt
			invoices = invoices[:limit-1]
		}

		// check if there is next page
		if cursorString == "" {
			nextCursor = nil
		} else {
			if len(invoices) == 0 {
				nextCursor = nil
			} else {
				nextCursor = &invoices[len(invoices)-1].CreatedAt
			}
		}
	}

	return invoices, nextCursor, prevCursor, http.StatusOK, nil
}
