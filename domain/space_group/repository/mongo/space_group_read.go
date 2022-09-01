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

func (repo *spaceGroupMongoRepository) FindSpaceGroups(ctx context.Context, branchId string, withTrashed bool) ([]domain.SpaceGroup, int, error) {
	var spaceGroups []domain.SpaceGroup

	var filterSpaceGroup bson.M
	var filterSpace bson.M

	filterWithTrashed := bson.M{"branch_uuid": branchId}

	filterWithoutTrashed := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}
	filterSpaceWithoutTrashed := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"spaces.deleted_at": bson.M{"$exists": false}},
		},
	}

	if withTrashed {
		filterSpaceGroup = filterWithTrashed
		filterSpace = filterWithTrashed
	} else {
		filterSpaceGroup = filterWithoutTrashed
		filterSpace = filterSpaceWithoutTrashed
	}

	matchStage := bson.D{{Key: "$match", Value: filterSpaceGroup}}
	unwindSpacesStage := bson.D{{
		Key: "$unwind", Value: bson.M{
			"path":                       "$spaces",
			"preserveNullAndEmptyArrays": true,
		},
	}}
	filterDeletedSpacesStage := bson.D{{Key: "$match", Value: filterSpace}}
	groupByUuidStage := bson.D{{
		Key: "$group", Value: bson.M{
			"_id":         "$uuid",
			"uuid":        bson.M{"$first": "$uuid"},
			"branch_uuid": bson.M{"$first": "$branch_uuid"},
			"spaces":      bson.M{"$push": "$spaces"},
			"code":        bson.M{"$first": "$code"},
			"shape":       bson.M{"$first": "$shape"},
			"pax":         bson.M{"$first": "$pax"},
			"floor":       bson.M{"$first": "$floor"},
			"created_at":  bson.M{"$first": "$created_at"},
			"updated_at":  bson.M{"$first": "$updated_at"},
			"deleted_at":  bson.M{"$first": "$deleted_at"},
		},
	}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindSpacesStage, filterDeletedSpacesStage, groupByUuidStage})
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
			{"deleted_at": bson.M{"$exists": false}},
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
