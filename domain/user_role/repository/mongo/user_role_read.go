package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRoleReposiotry struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewUserRoleMongoRepository(DB mongo.Database) domain.UserRoleRepositoryContract {
	return &userRoleReposiotry{
		DB:         DB,
		Collection: *DB.Collection("user_roles"),
	}
}

func (repo *userRoleReposiotry) FindUserRole(ctx context.Context, id string, withTrashed bool) (*domain.UserRole, int, error) {
	var userRole domain.UserRole
	var filter bson.M
	var result *mongo.SingleResult
	var err error

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
	result = repo.Collection.FindOne(ctx, filter)
	err = result.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}
