package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *userRoleReposiotry) DeleteUserRole(ctx context.Context, id string) (*domain.UserRole, int, error) {
	var userRole domain.UserRole
	var filter bson.M
	var result *mongo.SingleResult
	var err error

	filter = bson.M{"uuid": id}
	result = repo.Collection.FindOneAndDelete(ctx, filter)
	err = result.Decode(&userRole)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &userRole, http.StatusOK, nil
}

func (repo *userRoleReposiotry) UpsertUserRole(ctx context.Context, data *domain.UserRole) (*domain.UserRole, int, error) {
	var userRole domain.UserRole

	filter := bson.M{"uuid": data.UUID}
	var contents bson.D
	opts := options.Update().SetUpsert(true)

	countUserRole, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countUserRole > 0 {
		updatedAt := time.Now().UTC().UnixMicro()

		update := bson.D{{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "power", Value: data.Power},
				{Key: "updated_at", Value: updatedAt},
			},
		}}

		contents = update
	} else {
		insert := bson.D{{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "power", Value: data.Power},
				{Key: "created_at", Value: data.CreatedAt},
			},
		}}

		contents = insert
	}

	if _, err = repo.Collection.UpdateOne(ctx, filter, contents, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&userRole); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return data, http.StatusOK, nil
}
