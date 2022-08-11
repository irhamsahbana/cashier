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

type menuCategoryMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewMenuCategoryMongoRepository(DB mongo.Database) domain.MenuCategoryRepositoryContract {
	return &menuCategoryMongoRepository{
			DB: DB,
			Collection: *DB.Collection("menu_categories"),
		}
}

func (repo *menuCategoryMongoRepository) FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, int, error) {
	var menucategories []domain.MenuCategory
	var menucategory domain.MenuCategory

	var filterMenuCategory bson.M
	var filterMenu bson.M

	if withTrashed == true {
		filterMenuCategory = bson.M{"uuid": id}
		filterMenu = bson.M{"uuid": id}
	} else {
		filterMenuCategory = bson.M{
			"$and": bson.A{
				bson.M{"uuid": id},
				bson.M{"deleted_at": bson.M{"$exists": false}},
			},
		}

		filterMenu = bson.M{
			"$or": bson.A{
				bson.M{"uuid": id},
				bson.M{"menus.deleted_at": bson.M{"$exists": false}},
			},
		}
	}

	matchStage := bson.D{{"$match", filterMenuCategory}}
	unwindMenusStage := bson.D{{
							"$unwind", bson.M{
								"path": "$menus",
								"preserveNullAndEmptyArrays": true,
							},
						}}
	filterMenusWithExistanceOfDeletedAtStage := bson.D{{"$match", filterMenu}}
	groupByUuidStage := bson.D{{
							"$group", bson.M{
								"_id": "$uuid",
								"uuid": bson.M{"$first": "$uuid"},
								"branch_uuid": bson.M{"$first": "$branch_uuid"},
								"name": bson.M{"$first": "$name"},
								"menus": bson.M{"$push": "$menus"},
								"created_at": bson.M{"$first": "$created_at"},
								"updated_at": bson.M{"$first": "$updated_at"},
								"deleted_at": bson.M{"$first": "$deleted_at"},
							},
						}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindMenusStage, filterMenusWithExistanceOfDeletedAtStage, groupByUuidStage})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &menucategories)
	err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(menucategories) > 0 {
		menucategory = menucategories[0]
	} else {
		return nil, http.StatusNotFound ,errors.New("Menu category not found")
	}

	return &menucategory, http.StatusOK, nil
}

func (repo *menuCategoryMongoRepository) FindMenuCategories(ctx context.Context, withTrashed bool) ([]domain.MenuCategory, int, error) {
	var menucategories []domain.MenuCategory

	var filterMenuCategory bson.M
	var filterMenu bson.M

	filterWithTrashed := bson.M{}

	filterWithoutTrashed := bson.M{"deleted_at": bson.M{"$exists": false}}
	filterMenuWithoutTrashed := bson.M{"menus.deleted_at": bson.M{"$exists": false}}

	if withTrashed == true {
		filterMenuCategory = filterWithTrashed
		filterMenu = filterWithTrashed
	} else {
		filterMenuCategory = filterWithoutTrashed
		filterMenu = filterMenuWithoutTrashed
	}

	matchStage := bson.D{{"$match", filterMenuCategory}}
	unwindMenusStage := bson.D{{
							"$unwind", bson.M{
								"path": "$menus",
								"preserveNullAndEmptyArrays": true,
							},
						}}
	filterMenusBaseOnExistanceOfDeletedAtStage := bson.D{{"$match", filterMenu}}
	groupByUuidStage := bson.D{{
						"$group", bson.M{
							"_id": "$uuid",
							"uuid": bson.M{"$first": "$uuid"},
							"branch_uuid": bson.M{"$first": "$branch_uuid"},
							"name": bson.M{"$first": "$name"},
							"menus": bson.M{"$push": "$menus"},
							"created_at": bson.M{"$first": "$created_at"},
							"updated_at": bson.M{"$first": "$updated_at"},
							"deleted_at": bson.M{"$first": "$deleted_at"},
						},
					}}

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindMenusStage, filterMenusBaseOnExistanceOfDeletedAtStage, groupByUuidStage})
	if err != nil {
		return menucategories, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &menucategories); err != nil {
		return menucategories, http.StatusInternalServerError, err
	}

	return menucategories, http.StatusOK, nil
}

// Menu

func (repo *menuCategoryMongoRepository) FindMenu(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	var filter bson.M
	var elemMatch bson.M

	println("pake trashed:", withTrashed)

	if withTrashed == true {
		filter = bson.M{"menus.uuid": id}

		elemMatch = bson.M{"uuid": id}
	} else {
		filter = bson.M{
			"$and": bson.A{
				bson.M{"menus.uuid": id},
				bson.M{"menus.deleted_at": bson.M{"$exists": false}},
			},
		}

		elemMatch = bson.M{
						"$and": bson.A{
							bson.M{"uuid": id},
							bson.M{"deleted_at": bson.M{"$exists": false}},
						},
					}
	}

	err := repo.Collection.FindOne(
									ctx,
									filter,
									options.FindOne().
											SetProjection(
														bson.M{
															"uuid": 1,
															"branch_uuid": 1,
															"name": 1,
															"created_at": 1,
															"updated_at": 1,
															"menus": bson.M{
																"$elemMatch": elemMatch,
															},
														},
													),
								).Decode(&menucategory)

		if err != nil {
			return &menucategory, err
		}

	return &menucategory, nil
}