package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewUserMongoRepository(DB mongo.Database) domain.UserRepositoryContract {
	return &userRepository{
		DB:         DB,
		Collection: *DB.Collection("users"),
	}
}

func (repo *userRepository) FindUserBy(ctx context.Context, key string, val interface{}, withTrashed bool) (*domain.User, int, error) {
	var user domain.User
	var filter bson.M

	if withTrashed {
		filter = bson.M{key: val}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{key: val},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("user not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}
