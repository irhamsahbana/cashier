package mongo

import (
	"context"
	"errors"
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

	_, err = repo.Collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repo *menuCategoryMongoRepository) DeleteMenuCategory(ctx context.Context, id string, forceDelete bool) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory
	var err error

	if forceDelete {
		err = repo.Collection.FindOneAndDelete(ctx, bson.M{"uuid": id}).Decode(&menucategory)
	} else {
		filter := bson.M{"uuid": id, "deleted_at": bson.M{"$exists": false}}

		err = repo.Collection.FindOneAndUpdate(
												ctx,
												filter,
												bson.M{"$set":
													bson.M{"deleted_at": time.Now()},
												},
											).Decode(&menucategory)
	}

	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}

func (repo *menuCategoryMongoRepository) UpdateMenuCategory(ctx context.Context, id string, data *domain.MenuCategoryUpdateRequest) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	_, err := repo.FindMenuCategory(ctx, id, false)
	if err != nil {
		return &menucategory, err
	}

	update := bson.M{
		"name": data.Name,
		"updated_at": time.Now(),
	}

	result, err := repo.Collection.UpdateOne(ctx, bson.M{"uuid": id}, bson.M{"$set": update})

	if err != nil {
		return &menucategory, err
	}

	if result.MatchedCount == 1 {
		err := repo.Collection.FindOne(ctx,bson.M{"uuid": id}).Decode(&menucategory)

		if err != nil {
			return &menucategory, err
		}
	}

	return &menucategory, nil
}

// Menu

func (repo *menuCategoryMongoRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory
	var err error

	// check if menu category is exists
	_, err = repo.FindMenuCategory(ctx, menuCategoryId, false)
	if err != nil {
		return &menucategory, err
	}

	_, err = repo.FindMenu(ctx, data.UUID, false)

	// validate if uuid or id for menu exists
	 countMenu, err := repo.Collection.CountDocuments(ctx, bson.M{"menus.uuid": data.UUID})
	if countMenu > 0 {
		return &menucategory, errors.New("uuid of menu is exists in menu category collection")
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
		return &menucategory, err
	}

	// if there is document effected by update query then ..
	if result.MatchedCount == 1 {
		// search effected document
		err = repo.Collection.FindOne(
										ctx,
										bson.M{
											"menus.uuid": data.UUID,
										},
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
			return &menucategory, err
		}
	}

	return &menucategory, err
}

func (repo *menuCategoryMongoRepository) DeleteMenu(ctx context.Context, id string, forceDelete bool) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory
	var err error

	if forceDelete {
		err = repo.Collection.FindOneAndUpdate(
											ctx,
											bson.M{"menus.uuid": id},
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
													"uuid": 1,
													"branch_uuid": 1,
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
	} else {
		filter := bson.M{"menus.uuid": id}

		_, err = repo.FindMenu(ctx, id, false)
		if err != nil {
			return &menucategory, err
		}

		err = repo.Collection.FindOneAndUpdate(
												ctx,
												filter,
												bson.M{
													"$set": bson.M{"menus.$[elem].deleted_at": time.Now()},
												},
												options.FindOneAndUpdate().
												SetArrayFilters(options.ArrayFilters{
													Filters: bson.A{
														bson.M{
															"elem.uuid": id,
															"elem.deleted_at": bson.M{"$exists": false},
														},
													},
												}),
												options.FindOneAndUpdate().
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
															},
														},
													},
												),
											).Decode(&menucategory)

	}

	if err != nil {
		return &menucategory, err
	}

	return &menucategory, nil
}