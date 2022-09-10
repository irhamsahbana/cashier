package mongo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *itemCategoryMongoRepository) UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, data *domain.ItemCategory) (*domain.ItemCategory, int, error) {
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
		return nil, http.StatusInternalServerError, err
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

	if err := repo.Collection.FindOne(ctx, filter).Decode(&itemcategory); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("item category not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	_, err := repo.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}

// Item

func (repo *itemCategoryMongoRepository) UpsertItemAndVariants(ctx context.Context, branchId, itemCategoryId string, data *domain.Item) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory
	var contents bson.D

	filter := bson.M{
		"$and": []bson.M{
			{"uuid": itemCategoryId},
			{"branch_uuid": branchId},
			{"items.uuid": data.UUID},
		},
	}

	countItemCategory, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Println("error 1")
		return nil, http.StatusInternalServerError, err
	}

	if countItemCategory > 0 {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "items.$.name", Value: data.Name},
				{Key: "items.$.price", Value: data.Price},
				{Key: "items.$.label", Value: data.Label},
				{Key: "items.$.description", Value: data.Description},
				{Key: "items.$.public", Value: data.Public},
				{Key: "items.$.image_path", Value: data.ImagePath},
				{Key: "items.$.variants", Value: data.Variants},
				{Key: "items.$.updated_at", Value: time.Now().UnixMicro()},
			}},
		}

		contents = update
	} else {
		insert := bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "items", Value: bson.D{
					{Key: "uuid", Value: data.UUID},
					{Key: "name", Value: data.Name},
					{Key: "price", Value: data.Price},
					{Key: "label", Value: data.Label},
					{Key: "description", Value: data.Description},
					{Key: "public", Value: data.Public},
					{Key: "image_path", Value: data.ImagePath},
					{Key: "variants", Value: data.Variants},
					{Key: "created_at", Value: data.CreatedAt},
				}},
			}},
		}

		contents = insert

		filter = bson.M{
			"$and": []bson.M{
				{"uuid": itemCategoryId},
				{"branch_uuid": branchId},
			},
		}
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, contents)
	if err != nil {
		fmt.Println("error 2")
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&itemcategory); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, err
		}

		fmt.Println("error 3")
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) DeleteItemAndVariants(ctx context.Context, branchId, id string) (*domain.ItemCategory, int, error) {
	filter := bson.M{
		"branch_uuid": branchId,
		"items.uuid":  id,
	}

	update := bson.M{
		"$pull": bson.M{
			"items": bson.M{"uuid": id},
		},
	}

	item, code, err := repo.FindItemAndVariants(ctx, branchId, id)
	if err != nil {
		return nil, code, err
	}

	result, err := repo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if result.MatchedCount == 0 {
		return nil, http.StatusNotFound, nil
	}

	return item, http.StatusOK, nil
}
