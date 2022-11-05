package mongo

import (
	"context"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *orderRepository) InsertRefund(ctx context.Context, branchId, invoiceId string, data *domain.InvoiceRefund, refunds []domain.OrderRefundData) (*domain.Invoice, int, error) {
	var invoice domain.Invoice

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": invoiceId},
		},
	}

	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "refunds", Value: data},
		}},
	}

	res, err := r.CollInvoice.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, 404, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, 500, err
	}

	if res.MatchedCount == 0 {
		return nil, 404, err
	}

	// bulk updated orders, (field refunded_qty) based on refunds.OrderUUID
	for _, refund := range refunds {
		filter := bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"uuid": invoiceId},
				{"order_groups.orders.uuid": refund.OrderUUID},
			},
		}

		// update with array filter
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "order_groups.$[og].orders.$[o].refunded_qty", Value: refund.Qty},
			}},
		}

		arrayFilters := []interface{}{
			bson.M{"og.uuid": refund.OrderGroupUUID},
			bson.M{"o.uuid": refund.OrderUUID},
		}

		res, err := r.CollInvoice.UpdateOne(ctx, filter, update, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: arrayFilters,
		}))

		fmt.Println("matchCount", res.MatchedCount)
		fmt.Println("modifiedCount", res.ModifiedCount)
		fmt.Println("upsertedId", res.UpsertedID)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, 404, err
			}

			logger.Log(logrus.Fields{}).Error(err)
			return nil, 500, err
		}

		if res.MatchedCount == 0 {
			return nil, 404, err
		}

	}

	err = r.CollInvoice.FindOne(ctx, filter).Decode(&invoice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, 404, err
		}

		logger.Log(logrus.Fields{}).Error(err)
		return nil, 500, err
	}

	fmt.Println("isi invoice", invoice)

	return &invoice, 200, nil
}
