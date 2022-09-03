package mongo

import (
	"context"
	"errors"
	"net/http"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type itemCategoryMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewItemCategoryMongoRepository(DB mongo.Database) domain.ItemCategoryRepositoryContract {
	return &itemCategoryMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("item_categories"),
	}
}

func (repo *itemCategoryMongoRepository) FindItemCategory(ctx context.Context, branchId, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	var itemCategory domain.ItemCategory
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

	err := repo.Collection.FindOne(ctx, filter).Decode(&itemCategory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("item category not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &itemCategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) FindItemCategories(ctx context.Context, branchId string, withTrashed bool) ([]domain.ItemCategory, int, error) {
	var itemcategories []domain.ItemCategory
	var pipeline mongo.Pipeline

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}

	if withTrashed {
		delete(filter["$and"].([]bson.M)[1], "deleted_at")
	}

	// projection
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})

	if !withTrashed {
		filterItem := bson.D{{Key: "$project", Value: bson.M{
			"uuid":        1,
			"branch_uuid": 1,
			"name":        1,
			"items": bson.M{
				"$filter": bson.M{
					"input": "$items",
					"as":    "item",
					"cond":  bson.M{"$lt": bson.A{"$$item.deleted_at", 0}},
				},
			},
			"created_at": 1,
			"updated_at": 1,
			"deleted_at": 1,
		}}}

		pipeline = append(pipeline, filterItem)
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := cursor.All(ctx, &itemcategories); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return itemcategories, http.StatusOK, nil
}

// Item

func (repo *itemCategoryMongoRepository) FindItem(ctx context.Context, branchId, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory

	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"items.uuid": id},
			{"items.deleted_at": bson.M{"$exists": false}},
		},
	}

	if withTrashed {
		delete(filter["$and"].([]bson.M)[2], "items.deleted_at")
	}

	// projection
	projection := bson.M{"spaces.$": 1}
	opts := options.FindOne().SetProjection(projection)

	if err := repo.Collection.FindOne(ctx, filter, opts).Decode(&itemcategory); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errors.New("item not found")
		}

		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}
