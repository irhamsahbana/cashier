package mongo

import (
	"context"
	"errors"
	"strconv"

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

func (repo *menuCategoryMongoRepository) FindMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	if id == "0" {
		return &menucategory, errors.New("menu category identifier not found!")
	}

	countMenuCategoryByUUID, err := repo.Collection.CountDocuments(ctx, bson.M{"uuid": id})
	if countMenuCategoryByUUID > 0 {
		err := repo.Collection.FindOne(ctx, bson.M{"uuid": id}).Decode(&menucategory)
		if err != nil {
			return &menucategory, err
		}

		return &menucategory, nil
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return &menucategory, errors.New("menu category id not found!")
	}

	err = repo.Collection.FindOne(ctx, bson.M{"id": intID}).Decode(&menucategory)
	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}

// Menu

func (repo *menuCategoryMongoRepository) FindMenu(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	if  id == "0" {
		return &menucategory, errors.New("menu identifier not found!")
	}

	countMenuByUUID, err := repo.Collection.CountDocuments(ctx, bson.M{"menus.uuid": id})
	if countMenuByUUID > 0 {
		err := repo.Collection.FindOne(
										ctx,
										bson.M{"menus.uuid": id},
										options.FindOne().
												SetProjection(
															bson.M{
																"id": 1,
																"uuid": 1,
																"branch_id": 1,
																"name": 1,
																"created_at": 1,
																"updated_at": 1,
																"menus": bson.M{
																	"$elemMatch": bson.M{
																		"uuid": id,
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

	intID, err := strconv.Atoi(id)
	if err != nil {
		return &menucategory, errors.New("menu id not found!")
	}

	err = repo.Collection.FindOne(
									ctx,
									bson.M{"menus.id": intID},
									options.FindOne().
											SetProjection(
												bson.M{
													"id": 1,
													"uuid": 1,
													"branch_id": 1,
													"name": 1,
													"created_at": 1,
													"updated_at": 1,
													"menus": bson.M{
														"$elemMatch": bson.M{
															"menus.id": id,
														},
													},
												},
											),
								).
								Decode(&menucategory)

	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}