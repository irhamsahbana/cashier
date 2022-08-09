package mongo

import (
	"context"
	"errors"
	"net/http"
	"time"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *menuCategoryMongoRepository) UpsertMenuCategory(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, int, error) {
	var menucategory domain.MenuCategory
	var contents bson.M

	filter := bson.M{"uuid": data.UUID}

	countMenuCategory, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return &menucategory, http.StatusInternalServerError, err
	}

	if countMenuCategory > 0 {
		updatedAt := time.Now().UnixMicro()

		update := bson.M{
			"$set": bson.M{
				"name": data.Name,
				"updated_at": updatedAt,
			},
		}

		contents = update
	} else {
		insert := bson.M{
			"$set": bson.M{
				"branch_uuid": data.BranchUUID,
				"name": data.Name,
				"created_at": data.CreatedAt,
			},
		}

		contents = insert
	}


	opts := options.Update().SetUpsert(true)

	if _, err := repo.Collection.UpdateOne(ctx, filter, contents, opts);  err != nil {
		return &menucategory, http.StatusInternalServerError, err
	}

	if err := repo.Collection.FindOne(ctx, filter).Decode(&menucategory); err != nil {
		return &menucategory, http.StatusInternalServerError, err
	}

	return &menucategory, http.StatusOK, nil
}

func (repo *menuCategoryMongoRepository) DeleteMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	var menucategory domain.MenuCategory

	result, err := repo.Collection.UpdateOne(
										ctx,
										bson.M{"uuid": id},
										bson.A{
											bson.M{
												"$set": bson.M{
													"deleted_at": bson.M{
														"$ifNull": bson.A{
															"$deleted_at",
															time.Now().UnixMicro(),
														},
													},
												},
											},
										},
									)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if result.ModifiedCount > 0 {
		findOne := repo.Collection.FindOne(ctx, bson.M{"uuid": id})

		if err = findOne.Decode(&menucategory);  err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &menucategory, http.StatusOK, nil
}

// Menu

func (repo *menuCategoryMongoRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.MenuCategory, int, error) {
	var menucategory domain.MenuCategory

	_, code, err := repo.FindMenuCategory(ctx, menuCategoryId, true)
	if err != nil {
		return nil, code, err
	}

	// validate if uuid for menu exists
	 countMenu, err := repo.Collection.CountDocuments(ctx, bson.M{"menus.uuid": data.UUID})
	if countMenu > 0 {
		return nil, http.StatusConflict, errors.New("uuid of menu is exists in menu category collection")
	}

	// create a menu inside a collection (in 'menus' field)
	result, err := repo.Collection.UpdateOne(
											ctx,
											bson.M{"uuid": menuCategoryId},
											bson.A{
												bson.M{
													"$set": bson.M{
														"menus": bson.M{
															"$ifNull": bson.A{
																bson.M{"$concatArrays": bson.A{"$menus", bson.A{data}}},
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
										bson.M{"menus.uuid": data.UUID},
										options.FindOne().
												SetProjection(
													bson.M{
														"uuid": 1,
														"branch_uuid": 1,
														"name": 1,
														"created_at": 1,
														"menus": bson.M{
															"$elemMatch": bson.M{
																"uuid": data.UUID,
															},
														},
													},
												),
									).Decode(&menucategory)

		// if can't marshaled, then return error
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &menucategory, http.StatusOK, nil
}

func (repo *menuCategoryMongoRepository) DeleteMenu(ctx context.Context, id string) (*domain.MenuCategory, int, error) {
	var menucategory domain.MenuCategory

	filter := bson.M{"menus.uuid": id}
	update := bson.M{
				"$set": bson.M{
					"menus.$[elem].deleted_at": bson.M{
						"$ifNull": bson.A{
							"$menus.$[elem].deleted_at",
							time.Now().UnixMicro(),
						},
					},
				},
			}

	arrayFilter :=  options.ArrayFilters{
						Filters: bson.A{
							bson.M{
								"elem.uuid": id,
								"elem.deleted_at": bson.M{"$exists": false},
							},
						},
					}

	var updateOptions options.UpdateOptions
	updateOptions.ArrayFilters = &arrayFilter

	result, err := repo.Collection.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if result.ModifiedCount > 0 {
		menu, err :=  repo.FindMenu(ctx, id, true)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return menu, http.StatusOK, nil
	}

	return &menucategory, http.StatusOK, nil
}