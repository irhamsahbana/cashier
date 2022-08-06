package mongo

import (
	"context"
	"errors"

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

func (repo *menuCategoryMongoRepository) FindMenuCategory(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, error) {
	var menucategories []domain.MenuCategory
	var menucategory domain.MenuCategory

	var filterMenuCategory bson.M
	var filterMenu bson.M

	filterWithTrashed := bson.M{"uuid": id}

	filterWithoutTrashed := bson.M{
		"$and": bson.A{
			bson.M{"uuid": id},
			bson.M{"deleted_at": bson.M{"$exists": false}},
		},
	}

	filterMenuWithoutTrashed := bson.M{
		"$and": bson.A{
			bson.M{"uuid": id},
			bson.M{"menus.deleted_at": bson.M{"$exists": false}},
		},
	}

	if withTrashed == true {
		filterMenuCategory = filterWithTrashed
		filterMenu = filterWithTrashed
	} else {
		filterMenuCategory = filterWithoutTrashed
		filterMenu = filterMenuWithoutTrashed
	}

	matchStage := bson.D{{"$match", filterMenuCategory}}
	unwindMenusStage := bson.D{{"$unwind", "$menus"}}
	filterMenusWithExistanceOfDeletedAt := bson.D{{"$match", filterMenu}}
	groupByUuid := bson.D{{
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

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindMenusStage, filterMenusWithExistanceOfDeletedAt, groupByUuid})
	if err != nil {
		return &menucategory, err
	}

	if err = cursor.All(ctx, &menucategories); err != nil {
		return &menucategory, err
	}

	if len(menucategories) == 0 {
		return &menucategory, errors.New("Menu category not found")
	}

	menucategory = menucategories[0]
	return &menucategory, nil
}

func (repo *menuCategoryMongoRepository) FindMenuCategories(ctx context.Context, withTrashed bool) ([]domain.MenuCategory, error) {
	var menucategories []domain.MenuCategory

	var filterMenuCategory bson.M
	var filterMenu bson.M

	filterWithTrashed := bson.M{
						// "uuid": id,
					}

	filterWithoutTrashed := bson.M{
		"$and": bson.A{
			// bson.M{"uuid": id},
			bson.M{"deleted_at": bson.M{"$exists": false}},
		},
	}

	filterMenuWithoutTrashed := bson.M{
		"$and": bson.A{
			// bson.M{"uuid": id},
			bson.M{"menus.deleted_at": bson.M{"$exists": false}},
		},
	}

	if withTrashed == true {
		filterMenuCategory = filterWithTrashed
		filterMenu = filterWithTrashed
	} else {
		filterMenuCategory = filterWithoutTrashed
		filterMenu = filterMenuWithoutTrashed
	}

	matchStage := bson.D{{"$match", filterMenuCategory}}
	unwindMenusStage := bson.D{{"$unwind", "$menus"}}
	filterMenusWithExistanceOfDeletedAt := bson.D{{"$match", filterMenu}}
	groupByUuid := bson.D{{
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

	cursor, err := repo.Collection.Aggregate(ctx, mongo.Pipeline{matchStage, unwindMenusStage, filterMenusWithExistanceOfDeletedAt, groupByUuid})
	if err != nil {
		return menucategories, err
	}

	if err = cursor.All(ctx, &menucategories); err != nil {
		return menucategories, err
	}

	return menucategories, nil
}

// Menu

func (repo *menuCategoryMongoRepository) FindMenu(ctx context.Context, id string, withTrashed bool) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

		err := repo.Collection.FindOne(
										ctx,
										bson.M{
											"menus.uuid": id,
											"menus.deleted_at": bson.M{"$exists": withTrashed},
										},
										options.FindOne().
												SetProjection(
															bson.M{
																"uuid": 1,
																"branch_uuid": 1,
																"name": 1,
																"created_at": 1,
																"updated_at": 1,
																"menus": bson.M{
																	"$elemMatch": bson.M{
																		"uuid": id,
																		"deleted_at": bson.M{"$exists": withTrashed},
																	},
																},
															},
														),
									).Decode(&menucategory)

		if err != nil {
			return &menucategory, err
		}

	return &menucategory, nil
}