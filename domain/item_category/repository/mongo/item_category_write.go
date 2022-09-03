package mongo

import (
	"context"
	"net/http"
	"time"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *itemCategoryMongoRepository) UpsertItemCategory(ctx context.Context, branchId string, data *domain.ItemCategory) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory
	var contents bson.M

	filter := bson.M{"uuid": data.UUID}

	countItemCategory, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return &itemcategory, http.StatusInternalServerError, err
	}

	if countItemCategory > 0 {
		updatedAt := time.Now().UTC().UnixMicro()

		update := bson.M{
			"$set": bson.M{
				"name":       data.Name,
				"updated_at": updatedAt,
			},
		}

		contents = update
	} else {
		insert := bson.M{
			"$set": bson.M{
				"branch_uuid": data.BranchUUID,
				"name":        data.Name,
				"created_at":  data.CreatedAt,
			},
		}

		contents = insert
	}

	opts := options.Update().SetUpsert(true)

	if _, err := repo.Collection.UpdateOne(ctx, filter, contents, opts); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&itemcategory); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) DeleteItemCategory(ctx context.Context, branchId, id string) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory

	filter := bson.M{"uuid": id}
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

	findOne := repo.Collection.FindOne(ctx, bson.M{"uuid": id})

	if err = findOne.Decode(&itemcategory); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}

// Item

func (repo *itemCategoryMongoRepository) InsertItem(ctx context.Context, branchId, itemCategoryId string, data *domain.Item) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory

	_, code, err := repo.FindItemCategory(ctx, branchId, itemCategoryId, false)
	if err != nil {
		return nil, code, err
	}

	// create a item inside a collection (in 'items' field)
	result, err := repo.Collection.UpdateOne(
		ctx,
		bson.M{"uuid": itemCategoryId},
		bson.A{
			bson.M{
				"$set": bson.M{
					"items": bson.M{
						"$ifNull": bson.A{
							bson.M{"$concatArrays": bson.A{"$items", bson.A{data}}},
							bson.A{data},
						},
					},
				},
			},
		},
	)

	// check if when update error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// if there is document effected by update query then ..
	if result.MatchedCount == 1 {
		// search effected document
		err = repo.Collection.FindOne(
			ctx,
			bson.M{"items.uuid": data.UUID},
			options.FindOne().
				SetProjection(
					bson.M{
						"uuid":        1,
						"branch_uuid": 1,
						"name":        1,
						"created_at":  1,
						"items": bson.M{
							"$elemMatch": bson.M{
								"uuid": data.UUID,
							},
						},
					},
				),
		).Decode(&itemcategory)

		// if can't marshaled, then return error
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &itemcategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) UpdateItem(ctx context.Context, branchId, id string, data *domain.Item) (*domain.ItemCategory, int, error) {
	filter := bson.M{"items.uuid": id}

	arrayFilters := options.ArrayFilters{
		Filters: bson.A{
			bson.M{"elem.uuid": id},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"items.$[elem].name":        data.Name,
			"items.$[elem].price":       data.Price,
			"items.$[elem].description": data.Description,
			"items.$[elem].label":       data.Label,
			"items.$[elem].public":      data.Public,
			"items.$[elem].updated_at":  time.Now().UTC().UnixMicro(),
		},
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

	item, _, err := repo.FindItem(ctx, branchId, id, true)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return item, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) DeleteItem(ctx context.Context, branchId, id string) (*domain.ItemCategory, int, error) {
	filter := bson.M{"items.uuid": id}

	arrayFilters := options.ArrayFilters{
		Filters: bson.A{
			bson.M{
				"elem.uuid":       id,
				"elem.deleted_at": bson.M{"$exists": false},
			},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"items.$[elem].deleted_at": time.Now().UTC().UnixMicro(),
		},
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

	item, _, err := repo.FindItem(ctx, branchId, id, true)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return item, http.StatusOK, nil
}
