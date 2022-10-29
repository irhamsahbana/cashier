package mongo

import (
	"context"
	"lucy/cashier/dto"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *employeeShiftMongoRepository) Summary(ctx context.Context, branchId string, shiftIds []string) (*dto.EmployeeShiftSummaryResponse, int, error) {
	summary := dto.EmployeeShiftSummaryResponse{}
	orders := []dto.EmployeeShiftSummaryOrder{}
	payments := []dto.EmployeeShiftSummaryPayment{}

	pipeline := mongo.Pipeline{
		// filter base on employee shift
		{{Key: "$match", Value: bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
			},
		}}},
		{{Key: "$unwind", Value: "$payments"}},
		{{Key: "$match", Value: bson.M{
			"$and": []bson.M{
				{"payments.employee_shift.employee_shift_uuid": bson.M{
					"$in": shiftIds,
				}},
			},
		}}},
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"uuid": "$uuid",
			},
			"order_groups": bson.M{"$first": "$order_groups"},
			"payments":     bson.M{"$push": "$payments"},
		}}},
		// counting orders
		{{Key: "$unwind", Value: "$order_groups"}},
		{{Key: "$unwind", Value: "$order_groups.orders"}},
		{{Key: "$addFields", Value: bson.M{
			"item":                   "$order_groups.orders.item",
			"item_refunded_quantity": "$order_groups.orders.refunded_quantity",
		}}},
		{{Key: "$addFields", Value: bson.M{
			"item_uuid":     "$item.uuid",
			"item_price":    "$item.price",
			"item_quantity": "$item.quantity",
		}}},
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"item_uuid":  "$item_uuid",
				"item_price": "$item_price",
			},
			"quantity": bson.M{
				"$sum": "$item_quantity",
			},
			"refunded_quantity": bson.M{
				"$sum": "$item_refunded_quantity",
			},
		}}},
		{{Key: "$addFields", Value: bson.M{
			"uuid":  "$_id.item_uuid",
			"price": "$_id.item_price",
		}}},
	}

	cursor, err := r.CollInvoice.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}
	summary.Orders = orders
	summary.Payments = payments

	return &summary, http.StatusOK, nil
}
