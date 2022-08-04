package mongo

import (
	"context"
	"errors"
	"strconv"
	"time"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *menuCategoryMongoRepository) InsertMenuCategory(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, error) {
	// validate if uuid or id for menu exists
	countMenuCategory, err := repo.Collection.CountDocuments(ctx, bson.M{"uuid": data.UUID})
	if err != nil {
		return data, err
	}
	if countMenuCategory > 0 {
		return data,  errors.New("uuid of menu category is exists in menu category collection")
	}

	countMenuCategory, err = repo.Collection.CountDocuments(ctx, bson.M{"id": data.ID})
	if err != nil {
		return data, err
	}
	if countMenuCategory > 0 && data.ID != 0 {
		return data,  errors.New("id of menu category is exists in menu category collection")
	}

	_, err = repo.Collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repo *menuCategoryMongoRepository) DeleteMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	_, err := repo.FindMenuCategory(ctx, id)
	if err != nil {
		return &menucategory, err
	}

	intID, _ := strconv.Atoi(id)

	err = repo.Collection.FindOneAndDelete(
								ctx,
								bson.M{
									"$or": bson.A{
										bson.M{"uuid": id},
										bson.M{"id": intID},
									},
								},
							).
							Decode(&menucategory)

	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}

func (repo *menuCategoryMongoRepository) UpdateMenuCategory(ctx context.Context, id string, data *domain.MenuCategoryUpdateRequest) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	_, err := repo.FindMenuCategory(ctx, id)
	if err != nil {
		return &menucategory, err
	}

	intID, _ := strconv.Atoi(id)

	update := bson.M{
		"name": data.Name,
		"updated_at": time.Now().Format(time.RFC3339),
	}

	result, err := repo.Collection.UpdateOne(
										ctx,
										bson.M{
											"$or": bson.A{
												bson.M{"uuid": id},
												bson.M{"id": intID},
											},
										},
										bson.M{"$set": update},
									)

	if err != nil {
		return &menucategory, err
	}

	if result.MatchedCount == 1 {
		err := repo.Collection.FindOne(
			ctx,
			bson.M{
				"$or": bson.A{
					bson.M{"uuid": id},
					bson.M{"id": intID},
				},
			},
		).
		Decode(&menucategory)

		if err != nil {
			return &menucategory, err
		}
	}

	return &menucategory, nil
}

// Menu

func (repo *menuCategoryMongoRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	// check if menu category is exists
	_, err := repo.FindMenuCategory(ctx, menuCategoryId)
	if err != nil {
		return &menucategory, err
	}

	intMenuCategoryId, _ := strconv.Atoi(menuCategoryId)

	// validate if uuid or id for menu exists
	 countMenu, err := repo.Collection.CountDocuments(ctx, bson.M{"menus.uuid": data.UUID})
	if countMenu > 0 {
		return &menucategory, errors.New("uuid of menu is exists in menu category collection")
	}

	countMenu, err = repo.Collection.CountDocuments(ctx, bson.M{"menus.id": data.ID})
	if countMenu > 0 && data.ID != 0 {
		return &menucategory, errors.New("id of menu is exists in menu category collection")
	}

	// create a menu inside a collection (in 'menus' field)
	result, err := repo.Collection.UpdateOne(
											ctx,
											bson.M{
												"$or": bson.A{
													bson.M{"uuid": menuCategoryId},
													bson.M{"id": intMenuCategoryId},
												},
											},
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
		return &menucategory, err
	}

	// if there is document effected by update query then ..
	if result.MatchedCount == 1 {
		// search effected document
		err = repo.Collection.FindOne(
										ctx,
										bson.M{
											"$or":
												bson.A{
													bson.M{"uuid": menuCategoryId},
													bson.M{"id": intMenuCategoryId},
												},
											"menus.uuid": data.UUID,
											"menus.id": data.ID,
										},
										options.FindOne().
												SetProjection(
													bson.M{
														"id": 1,
														"uuid": 1,
														"branch_id": 1,
														"name": 1,
														"created_at": 1,
														"menus": bson.M{
															"$elemMatch": bson.M{
																"$and": bson.A{
																	bson.M{"uuid": data.UUID},
																	bson.M{"id": data.ID},
																},
															},
														},
													},
												),
									).Decode(&menucategory)

		// if can't marshaled, then return error
		if err != nil {
			return &menucategory, err
		}
	}

	return &menucategory, err
}

func (repo *menuCategoryMongoRepository) DeleteMenu(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	_, err := repo.FindMenu(ctx, id)
	if err != nil {
		return &menucategory, err
	}

	countMenuByUUID, err := repo.Collection.CountDocuments(ctx, bson.M{"menus.uuid": id})
	if countMenuByUUID > 0 {
		err = repo.Collection.FindOneAndUpdate(
			ctx,
			bson.M{
				"$or": bson.A{
					bson.M{"menus.uuid": id},
				},
			},
			bson.M{
				"$pull": bson.M{
					"menus": bson.M{
						"uuid": id,
					},
				},
			},
			options.FindOneAndUpdate().
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
		).
		Decode(&menucategory)

		if err != nil {
		return &menucategory, err
		}

		return &menucategory, nil
	}

	intId, _ := strconv.Atoi(id)

	err = repo.Collection.FindOneAndUpdate(
		ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"menus.id": intId},
			},
		},
		bson.M{
			"$pull": bson.M{
				"menus": bson.M{
					"id": intId,
				},
			},
		},
		options.FindOneAndUpdate().
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
						"id": intId,
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