package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type waiterMongoRepository struct {
	DB			mongo.Database
	Collection	mongo.Collection
}

func NewWaiterMongoRepository(DB mongo.Database) domain.WaiterRepositoryContract{
	return &waiterMongoRepository{
		DB: DB,
		Collection: *DB.Collection("waiters"),
	}
}

func (repo *waiterMongoRepository) FindWaiter(ctx context.Context, id string, withTrashed bool) (*domain.Waiter, int, error) {
	var waiter domain.Waiter
	var filter bson.M

	if withTrashed {
		filter = bson.M{"uuid": id}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	result := repo.Collection.FindOne(ctx, filter)

	if result == nil {
		return nil, http.StatusNotFound, errors.New("waiter not found")
	}

	err := result.Decode(waiter)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &waiter, http.StatusOK, nil
}