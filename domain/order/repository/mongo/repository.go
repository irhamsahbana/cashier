package mongo

import (
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewOrderMongoRepository(DB mongo.Database) domain.OrderRepositoryContract {
	return &orderRepository{
		DB:          DB,
		CollActive:  *DB.Collection("active_order_groups"),
		CollInvoice: *DB.Collection("invoices"),
		CollDeleted: *DB.Collection("deleted_order_groups"),
	}
}

type orderRepository struct {
	DB          mongo.Database
	CollActive  mongo.Collection
	CollInvoice mongo.Collection
	CollDeleted mongo.Collection
}
