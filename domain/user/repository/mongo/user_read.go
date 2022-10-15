package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
				bson.M{"deleted_at": nil},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Error(err)
			return nil, http.StatusNotFound, errors.New("user not found")
		}

		logger.Log(logrus.Fields{
			"error": err,
		}).Error("error while fetching user")
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func (repo *userRepository) FindUsers(ctx context.Context, branchId string, roles []string, limit, page int, withTrashed bool) ([]domain.User, int, error) {
	var users []domain.User
	var filter bson.M
	opts := *options.Find().SetLimit(int64(limit)).SetSkip(int64(page * limit))

	if withTrashed {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"branch_uuid": branchId},
				bson.M{"role_uuid": bson.M{"$in": roles}},
			},
		}

		if len(roles) == 0 {
			filter = bson.M{"branch_uuid": branchId}
		}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"branch_uuid": branchId},
				bson.M{"role_uuid": bson.M{"$in": roles}},
				bson.M{"deleted_at": nil},
			},
		}

		if len(roles) == 0 {
			filter = bson.M{
				"$and": bson.A{
					bson.M{"branch_uuid": branchId},
					bson.M{"deleted_at": nil},
				},
			}
		}
	}

	cursor, err := repo.Collection.Find(ctx, filter, &opts)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	return users, http.StatusOK, nil
}
