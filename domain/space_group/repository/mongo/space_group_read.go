package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *spaceGroupMongoRepository) FindSpaceGroups(ctx context.Context, branchId string, withTrashed bool) ([]domain.SpaceGroup, int, error) {
	var spaceGroups []domain.SpaceGroup
	var pipeline mongo.Pipeline

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"deleted_at": nil},
		},
	}

	if withTrashed {
		delete(filter["$and"].([]bson.M)[1], "deleted_at")
	}

	// projection
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})

	if !withTrashed {
		filterSpace := bson.D{
			{Key: "$project", Value: bson.M{
				"uuid":        1,
				"branch_uuid": 1,
				"name":        1,
				"code":        1,
				"shape":       1,
				"pax":         1,
				"length":      1,
				"reservable":  1,
				"disabled":    1,
				"spaces": bson.M{"$filter": bson.M{
					"input": "$spaces",
					"as":    "space",
					"cond":  bson.M{"$lt": bson.A{"$$space.deleted_at", 0}},
				}},
				"created_at": 1,
				"updated_at": 1,
				"deleted_at": 1,
			}},
		}

		pipeline = append(pipeline, filterSpace)
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := cursor.All(ctx, &spaceGroups); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return spaceGroups, http.StatusOK, nil
}

func (repo *spaceGroupMongoRepository) FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*domain.SpaceGroup, int, error) {
	var spaceGroup domain.SpaceGroup

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": id},
			{"deleted_at": nil},
		},
	}

	if withTrashed {
		delete(filter["$and"].([]bson.M)[2], "deleted_at")
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&spaceGroup); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("space group not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &spaceGroup, http.StatusOK, nil
}

// space

func (repo *spaceGroupMongoRepository) FindSpace(ctx context.Context, branchId, id string, withTrashed bool) (*domain.SpaceGroup, int, error) {
	var spaceGroup domain.SpaceGroup
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"spaces.uuid": id},
			{"spaces.deleted_at": nil},
		},
	}

	if withTrashed {
		delete(filter["$and"].([]bson.M)[2], "spaces.deleted_at")
	}

	// projection
	projection := bson.M{"spaces.$": 1}
	opts := options.FindOne().SetProjection(projection)

	if err := repo.Collection.FindOne(ctx, filter, opts).Decode(&spaceGroup); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("space not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &spaceGroup, http.StatusOK, nil
}
