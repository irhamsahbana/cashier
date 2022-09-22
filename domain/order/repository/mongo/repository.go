package mongo

import (
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewOrderRepository(DB mongo.Database) domain.OrderRepositoryContract {
	return &orderRepository{
		DB:         DB,
		Collection: *DB.Collection("orders"),
	}
}

type orderRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}
