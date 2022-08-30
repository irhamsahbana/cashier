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

func (repo *itemCategoryMongoRepository) FindItemCategory(ctx context.Context, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	var itemcategories []domain.ItemCategory
	var itemcategory domain.ItemCategory

	var filterItemCategory bson.M
	var filterItem bson.M

	if withTrashed == true {
		filterItemCategory = bson.M{"uuid": id}
		filterItem = bson.M{"uuid": id}
	} else {
		filterItemCategory = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}

		filterItem = bson.M{
			"$or": bson.A{
				bson.M{"uuid": id},
				bson.M{"items.deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	matchStage := bson.D{{Key: "$match", Value: filterItemCategory}}
	unwindItemsStage := bson.D{{
		Key: "$unwind", Value: bson.M{
			"path":                       "$items",
			"preserveNullAndEmptyArrays": true,
		},
	}}
	filterItemsWithExistanceOfDeletedAtStage := bson.D{{Key: "$match", Value: filterItem}}
	groupByUuidStage := bson.D{{
		Key: "$group", Value: bson.M{
			"_id":         "$uuid",
			"uuid":        bson.M{"$first": "$uuid"},
			"branch_uuid": bson.M{"$first": "$branch_uuid"},
			"name":        bson.M{"$first": "$name"},
			"items":       bson.M{"$push": "$items"},
			"created_at":  bson.M{"$first": "$created_at"},
			"updated_at":  bson.M{"$first": "$updated_at"},
			"deleted_at":  bson.M{"$first": "$deleted_at"},
		},
	}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindItemsStage, filterItemsWithExistanceOfDeletedAtStage, groupByUuidStage})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &itemcategories); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(itemcategories) > 0 {
		itemcategory = itemcategories[0]
	} else {
		return nil, http.StatusNotFound, errors.New("Item category not found")
	}

	return &itemcategory, http.StatusOK, nil
}

func (repo *itemCategoryMongoRepository) FindItemCategories(ctx context.Context, withTrashed bool) ([]domain.ItemCategory, int, error) {
	var itemcategories []domain.ItemCategory

	var filterItemCategory bson.M
	var filterItem bson.M

	filterWithTrashed := bson.M{}

	filterWithoutTrashed := bson.M{"deleted_at": bson.M{"$exists": false}}
	filterItemWithoutTrashed := bson.M{"items.deleted_at": bson.M{"$exists": false}}

	if withTrashed == true {
		filterItemCategory = filterWithTrashed
		filterItem = filterWithTrashed
	} else {
		filterItemCategory = filterWithoutTrashed
		filterItem = filterItemWithoutTrashed
	}

	matchStage := bson.D{{Key: "$match", Value: filterItemCategory}}
	unwindItemsStage := bson.D{{
		Key: "$unwind", Value: bson.M{
			"path":                       "$items",
			"preserveNullAndEmptyArrays": true,
		},
	}}
	filterItemsBaseOnExistanceOfDeletedAtStage := bson.D{{Key: "$match", Value: filterItem}}
	groupByUuidStage := bson.D{{
		Key: "$group", Value: bson.M{
			"_id":         "$uuid",
			"uuid":        bson.M{"$first": "$uuid"},
			"branch_uuid": bson.M{"$first": "$branch_uuid"},
			"name":        bson.M{"$first": "$name"},
			"items":       bson.M{"$push": "$items"},
			"created_at":  bson.M{"$first": "$created_at"},
			"updated_at":  bson.M{"$first": "$updated_at"},
			"deleted_at":  bson.M{"$first": "$deleted_at"},
		},
	}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindItemsStage, filterItemsBaseOnExistanceOfDeletedAtStage, groupByUuidStage})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &itemcategories); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return itemcategories, http.StatusOK, nil
}

// Item

func (repo *itemCategoryMongoRepository) FindItem(ctx context.Context, id string, withTrashed bool) (*domain.ItemCategory, int, error) {
	var itemcategory domain.ItemCategory

	var filter bson.M
	var elemMatch bson.M

	if withTrashed == true {
		filter = bson.M{"items.uuid": id}

		elemMatch = bson.M{"uuid": id}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"items.uuid": id},
				bson.M{"items.deleted_at": bson.M{"$exists": false}},
			},
		}

		elemMatch = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	countItem, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if countItem < 1 {
		return nil, http.StatusNotFound, errors.New("Item not found")
	}

	err = repo.Collection.FindOne(
		ctx,
		filter,
		options.FindOne().
			SetProjection(
				bson.M{
					"uuid":        1,
					"branch_uuid": 1,
					"name":        1,
					"created_at":  1,
					"updated_at":  1,
					"items": bson.M{
						"$elemMatch": elemMatch,
					},
				},
			),
	).Decode(&itemcategory)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &itemcategory, http.StatusOK, nil
}
