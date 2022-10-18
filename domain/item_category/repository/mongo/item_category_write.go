package mongo

import (
	"context"
	"errors"
	"net/http"
	"time"

	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *itemCategoryMongoRepository) UpsertItemCategoryAndModifiers(ctx context.Context, branchId string, data *domain.ItemCategory) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": data.UUID},
		},
	}
	countItemCategory, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countItemCategory > 0 {
		contents := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: data.Name},
				{Key: "modifier_groups", Value: data.ModifierGroups},
				{Key: "updated_at", Value: time.Now().UTC().UnixMicro()},
			}},
		}

		if _, err := repo.Collection.UpdateOne(ctx, filter, contents); err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else {
		contents := bson.D{
			{Key: "uuid", Value: data.UUID},
			{Key: "branch_uuid", Value: branchId},
			{Key: "name", Value: data.Name},
			{Key: "modifier_groups", Value: data.ModifierGroups},
			{Key: "items", Value: bson.A{}},
			{Key: "created_at", Value: time.Now().UTC().UnixMicro()},
			{Key: "updated_at", Value: nil},
		}

		if _, err := repo.Collection.InsertOne(ctx, contents); err != nil {
			return nil, http.StatusInternalServerError, err
		}
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
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if countItemCategory > 0 { // update item and its variants
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
				{"branch_uuid": branchId},
				{"uuid": itemCategoryId},
			},
		}
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, contents)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&itemcategory); err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{}).Warn(err)
			return nil, http.StatusNotFound, errors.New("item category not found")
		}

		logger.Log(logrus.Fields{}).Error(err)
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
