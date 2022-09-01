package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *spaceGroupMongoRepository) UpsertSpaceGroup(ctx context.Context, data *domain.SpaceGroup) (*domain.SpaceGroup, int, error) {
	var spacegroup domain.SpaceGroup
	var contents bson.M

	filter := bson.M{"uuid": data.UUID}
	opts := options.Update().SetUpsert(true)

	countSpaceGroup, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countSpaceGroup > 0 {
		updatedAt := time.Now().UTC().UnixMicro()

		update := bson.M{
			"$set": bson.M{
				"code":       data.Code,
				"shape":      data.Shape,
				"pax":        data.Pax,
				"floor":      data.Floor,
				"smooking":   data.Smooking,
				"updated_at": updatedAt,
			},
		}

		contents = update
	} else {
		insert := bson.M{
			"$set": bson.M{
				"branch_uuid": data.BranchUUID,
				"code":        data.Code,
				"shape":       data.Shape,
				"pax":         data.Pax,
				"floor":       data.Floor,
				"smooking":    data.Smooking,
				"created_at":  data.CreatedAt,
			},
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
		"$and": bson.A{
			bson.M{"branch_uuid": branchId},
			bson.M{"uuid": id},
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

	if result.ModifiedCount == 0 {
		return nil, http.StatusNotFound, nil
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&spacegroup); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &spacegroup, http.StatusOK, nil
}
