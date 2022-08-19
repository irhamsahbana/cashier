package mongo

import (
	"context"
	"errors"
	"fmt"
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

	countWaiter, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countWaiter < 1 {
		return nil, http.StatusNotFound, errors.New("Waiter not found")
	}

	result := repo.Collection.FindOne(ctx, filter)
	fmt.Println(result)

	err = result.Decode(&waiter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &waiter, http.StatusOK, nil
}