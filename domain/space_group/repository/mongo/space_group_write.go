package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *spaceGroupMongoRepository) UpsertSpaceGroup(ctx context.Context, data *domain.SpaceGroup) (*domain.SpaceGroup, int, error) {
	var spacegroup domain.SpaceGroup
	var contents bson.D

	filter := bson.M{
		"$and": []bson.M{
			{"uuid": data.UUID},
			{"branch_uuid": data.BranchUUID},
		},
	}
	opts := options.Update().SetUpsert(true)

	countSpaceGroup, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countSpaceGroup > 0 {
		updatedAt := time.Now().UTC().UnixMicro()

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "code", Value: data.Code},
				{Key: "shape", Value: data.Shape},
				{Key: "pax", Value: data.Pax},
				{Key: "floor", Value: data.Floor},
				{Key: "smooking", Value: data.Smooking},
				{Key: "updated_at", Value: updatedAt},
			}},
		}

		contents = update
	} else {
		insert := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "branch_uuid", Value: data.BranchUUID},
				{Key: "code", Value: data.Code},
				{Key: "shape", Value: data.Shape},
				{Key: "pax", Value: data.Pax},
				{Key: "floor", Value: data.Floor},
				{Key: "smooking", Value: data.Smooking},
				{Key: "created_at", Value: data.CreatedAt},
			}},
		}

		contents = insert
	}

	if _, err := repo.Collection.UpdateOne(ctx, filter, contents, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&spacegroup); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &spacegroup, http.StatusOK, nil
}

func (repo *spaceGroupMongoRepository) DeleteSpaceGroup(ctx context.Context, branchId, id string) (*domain.SpaceGroup, int, error) {
	var spacegroup domain.SpaceGroup

	filter := bson.M{
		"$and": []bson.M{
			{"uuid": id},
			{"branch_uuid": branchId},
		},
	}

	update := bson.A{
		bson.M{
			"$set": bson.M{
				"deleted_at": bson.M{
					"$ifNull": bson.A{
						"$deleted_at",
						time.Now().UTC().UnixMicro(),
					},
				},
			},
		},
	}

	result, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, nil
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&spacegroup); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &spacegroup, http.StatusOK, nil
}

// Space

func (repo *spaceGroupMongoRepository) InsertSpace(ctx context.Context, branchId, SpaceGroupId string, data *domain.Space) (*domain.SpaceGroup, int, error) {
	var spaceGroup domain.SpaceGroup

	_, code, err := repo.FindSpaceGroup(ctx, branchId, SpaceGroupId, false)
	if err != nil {
		return nil, code, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": SpaceGroupId},
		},
	}

	insert := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "spaces", Value: bson.D{
				{Key: "uuid", Value: data.UUID},
				{Key: "number", Value: data.Number},
				{Key: "occupied", Value: data.Occupied},
				{Key: "description", Value: data.Description},
				{Key: "created_at", Value: data.CreatedAt},
			}},
		}},
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, insert)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// projection
	filter = bson.M{"spaces.uuid": data.UUID}
	projection := bson.M{"spaces.$": 1}
	opts := options.FindOne().SetProjection(projection)

	if err := repo.Collection.FindOne(ctx, filter, opts).Decode(&spaceGroup); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &spaceGroup, http.StatusOK, nil
}

func (repo *spaceGroupMongoRepository) DeleteSpace(ctx context.Context, branchId, id string) (*domain.SpaceGroup, int, error) {
	filter := bson.M{
		"$and": bson.A{
			bson.M{"branch_uuid": branchId},
			bson.M{"spaces.uuid": id},
		},
	}

	arrayFilters := options.ArrayFilters{
		Filters: bson.A{
			bson.M{
				"elem.uuid":       id,
				"elem.deleted_at": bson.M{"$exists": false},
			},
		},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "spaces.$[elem].deleted_at", Value: time.Now().UTC().UnixMicro()},
		}},
	}

	var updateOptions options.UpdateOptions
	updateOptions.ArrayFilters = &arrayFilters

	result, err := repo.Collection.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if result.ModifiedCount == 0 {
		return nil, http.StatusNotFound, errors.New("space not found")
	}

	space, code, err := repo.FindSpace(ctx, branchId, id, false)
	if err != nil {
		return nil, code, err
	}

	return space, http.StatusOK, nil
}

func (repo *spaceGroupMongoRepository) UpdateSpace(ctx context.Context, branchId, id string, data *domain.Space) (*domain.SpaceGroup, int, error) {
	// filter := bson.M{"branch_uuid": branchId, "spaces.uuid": id}
	filter := bson.M{"$and": bson.A{
		bson.M{"branch_uuid": branchId},
		bson.M{"spaces.uuid": id},
		bson.M{"spaces.deleted_at": bson.M{"$exists": false}},
	}}

	arrayFilters := options.ArrayFilters{
		Filters: bson.A{
			bson.M{"elem.uuid": id},
		},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "spaces.$[elem].number", Value: data.Number},
			{Key: "spaces.$[elem].occupied", Value: data.Occupied},
			{Key: "spaces.$[elem].description", Value: data.Description},
			{Key: "spaces.$[elem].updated_at", Value: time.Now().UTC().UnixMicro()},
		}},
	}

	var updateOptions options.UpdateOptions
	updateOptions.ArrayFilters = &arrayFilters

	result, err := repo.Collection.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, nil
	}

	space, code, err := repo.FindSpace(ctx, branchId, id, false)
	if err != nil {
		return nil, code, err
	}

	return space, http.StatusOK, nil
}
