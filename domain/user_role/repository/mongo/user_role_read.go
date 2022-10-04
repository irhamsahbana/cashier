package mongo

import (
	"context"
	"errors"
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

	if withTrashed {
		filter = bson.M{"uuid": id}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": nil},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&userRole)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("user role not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}

func (repo *userRoleReposiotry) FindUserRoleByName(ctx context.Context, name string, withTrashed bool) (*domain.UserRole, int, error) {
	var userRole domain.UserRole
	var filter bson.M

	if withTrashed {
		filter = bson.M{"name": name}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"name": name},
				bson.M{"deleted_at": nil},
			},
		}
	}

	err := repo.Collection.FindOne(ctx, filter).Decode(&userRole)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("user role not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}

func (repo *userRoleReposiotry) FindUserRoleBy(ctx context.Context, key string, val interface{}) (*domain.UserRole, int, error) {
	var userRole domain.UserRole

	filter := bson.M{key: val}
	result := repo.Collection.FindOne(ctx, filter)
	err := result.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}
