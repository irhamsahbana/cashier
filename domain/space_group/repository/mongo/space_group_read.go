package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type spaceGroupMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewSpaceGroupMongoRepository(DB mongo.Database) domain.SpaceGroupRepositoryContract {
	return &spaceGroupMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("space_groups"),
	}
}

func (repo *spaceGroupMongoRepository) FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*domain.SpaceGroup, int, error) {
	var spaceGroup domain.SpaceGroup
	var filter bson.M

	if withTrashed {
		filter = bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"uuid": id},
			},
		}
	} else {
		filter = bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"uuid": id},
				{"deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&spaceGroup); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("space group not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &spaceGroup, http.StatusOK, nil
}
