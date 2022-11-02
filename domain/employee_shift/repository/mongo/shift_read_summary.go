package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/helper"
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

	itemCategories := []domain.ItemCategory{}

	type totalRefund struct {
		TotalRefunds int64 `bson:"total_refunds"`
	}
	totalRefunds := []totalRefund{}

	// orders summary
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

	// search for item categories
	itemOrderIds := []string{}
	for _, order := range orders {
		if exists := helper.Contain(itemOrderIds, order.UUID); !exists {
			itemOrderIds = append(itemOrderIds, order.UUID)
		}
	}

	if len(itemOrderIds) != 0 {
		filter := bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"$or": []bson.M{
					{"items.uuid": bson.M{
						"$in": itemOrderIds,
					}},
					{"items.variants.uuid": bson.M{
						"$in": itemOrderIds,
					}},
				}},
			},
		}

		cursor, err = r.CollItemCategory.Find(ctx, filter)
		if err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return &summary, http.StatusInternalServerError, err
		}

		if err = cursor.All(ctx, &itemCategories); err != nil {
			logger.Log(logrus.Fields{}).Error(err)
			return &summary, http.StatusInternalServerError, err
		}
	}

	// give name and category name to orders summary
	finalOrders := []dto.EmployeeShiftSummaryOrder{}
	for _, order := range orders {
		for _, itemCategory := range itemCategories {
			// check for main item
			for _, item := range itemCategory.Items {
				if order.UUID == item.UUID {
					var x dto.EmployeeShiftSummaryOrder
					x = order
					x.Name = item.Name
					x.Category = itemCategory.Name
					finalOrders = append(finalOrders, x)
					break
				}

				// check for variant
				for _, variant := range item.Variants {
					if order.UUID == variant.UUID {
						var x dto.EmployeeShiftSummaryOrder
						x = order
						x.Name = item.Name
						x.Category = itemCategory.Name
						finalOrders = append(finalOrders, x)
						break
					}
				}
			}
		}
	}
	summary.Orders = finalOrders

	// payments summary
	pipeline = mongo.Pipeline{
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
			"payments": bson.M{
				"$push": "$payments",
			},
		}}},
		{{Key: "$unwind", Value: "$payments"}},
		{{Key: "$addFields", Value: bson.M{
			"payment": "$payments",
		}}},
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"payment_method_uuid": "$payment.payment_method.payment_method_uuid",
			},
			"quantity": bson.M{
				"$sum": 1,
			},
			"total": bson.M{
				"$sum": bson.M{
					"$convert": bson.M{
						"input": "$payment.total",
						"to":    "decimal",
					},
				},
			},
		}}},
		{{Key: "$addFields", Value: bson.M{
			"uuid": "$_id.payment_method_uuid",
			"total": bson.M{
				"$convert": bson.M{
					"input": "$total",
					"to":    "double",
				},
			},
		}}},
	}

	cursor, err = r.CollInvoice.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &payments); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}
	summary.Payments = payments

	// total refunds
	pipeline = mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
			},
		}}},
		{{Key: "$unwind", Value: "$payments"}},
		{{Key: "$match", Value: bson.M{
			"$and": []bson.M{
				{"refunds.employee_shift.employee_shift_uuid": bson.M{
					"$in": shiftIds,
				}},
			},
		}}},
		{{Key: "$group", Value: bson.M{
			"_id":          bson.M{"uuid": "$uuid"},
			"order_groups": bson.M{"$first": "$order_groups"},
			"refunds":      bson.M{"$first": "$refunds"},
		}}},
		// count total refunds amount
		{{Key: "$unwind", Value: "$refunds"}},
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"uuid": "$uuid",
			},
			"total_refunds": bson.M{
				"$sum": bson.M{
					"$convert": bson.M{
						"input": "$refunds.total",
						"to":    "decimal",
					},
				},
			},
		}}},
	}
	// convert to mongoDB query
	// db.invoices.aggregate([
	// 	{
	// 		$match: {
	// 			$and: [
	// 				{ branch_uuid: "5f9f1b9b-1b1f-4b1f-9f9b-1b1f4b1f9f9b" },
	// 			],
	// 		},
	// 	},
	// 	{
	// 		$unwind: "$payments",
	// 	},
	// 	{
	// 		$match: {
	// 			$and: [
	// 				{
	// 					"refunds.employee_shift.employee_shift_uuid": {
	// 						$in: [
	// 							"5f9f1b9b-1b1f-4b1f-9f9b-1b1f4b1f9f9b",
	// 							"5f9f1b9b-1b1f-4b1f-9f9b-1b1f4b1f9f9b",
	// 						],
	// 					},
	// 				},
	// 			],
	// 		},
	// 	},
	// 	{
	// 		$group: {
	// 			_id: { uuid: "$uuid" },
	// 			order_groups: { $first: "$order_groups" },
	// 			refunds: { $first: "$refunds" },
	// 		},
	// 	},
	// 	{
	// 		$unwind: "$refunds",
	// 	},
	// 	{
	// 		$group: {
	// 			_id: { uuid: "$uuid" },
	// 			total_refunds: { $sum: { $convert: { input: "$refunds.total", to: "decimal" } } },
	// 		},
	// 	},
	// ])

	cursor, err = r.CollInvoice.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &totalRefunds); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return &summary, http.StatusInternalServerError, err
	}

	var finalTotalRefunds int64
	for _, totalRefund := range totalRefunds {
		finalTotalRefunds += totalRefund.TotalRefunds
	}

	summary.TotalRefunds = finalTotalRefunds

	return &summary, http.StatusOK, nil
}
