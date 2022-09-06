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
	var contents bson.D

	filter := bson.M{
		"$and": []bson.M{
			{"uuid": data.UUID},
			{"branch_uuid": branchId},
		},
	}
	opts := options.Update().SetUpsert(true)

	countItemCategory, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return &itemcategory, http.StatusInternalServerError, err
	}

	if countItemCategory > 0 {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "modifier_groups", Value: data.ModifierGroups},
				{Key: "updated_at", Value: time.Now().UTC().UnixMicro()},
			}},
		}

		contents = update
	} else {
		insert := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "modifier_groups", Value: data.ModifierGroups},
				{Key: "created_at", Value: time.Now().UTC().UnixMicro()},
			}},
		}

		contents = insert
	}

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

	if err := repo.Collection.FindOne(ctx, filter).Decode(&itemcategory); err != nil {
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

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": itemCategoryId},
		},
	}

	insert := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "items", Value: bson.D{
				{Key: "uuid", Value: data.UUID},
				{Key: "main_uuid", Value: data.MainUUID},
				{Key: "name", Value: data.Name},
				{Key: "price", Value: data.Price},
				{Key: "public", Value: data.Public},
				{Key: "label", Value: data.Label},
				{Key: "image_url", Value: data.ImageUrl},
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
	filter = bson.M{"items.uuid": data.UUID}
	projection := bson.M{"items.$": 1}
	opts := options.FindOne().SetProjection(projection)

	if err = repo.Collection.FindOne(ctx, filter, opts).Decode(&itemcategory); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) UpdateItem(ctx context.Context, branchId, id string, data *domain.Item) (*domain.ItemCategory, int, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"items.uuid": id},
			{"spaces.deleted_at": bson.M{"$exists": false}},
		},
	}

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

	item, code, err := repo.FindItem(ctx, branchId, id, false)
	if err != nil {
		return nil, code, err
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
