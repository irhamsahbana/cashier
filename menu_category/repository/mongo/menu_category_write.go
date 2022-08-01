package mongo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *menuCategoryMongoRepository) InsertMenuCategory(ctx context.Context, data *domain.MenuCategory) (*domain.MenuCategory, error) {
	_, err := repo.Collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repo *menuCategoryMongoRepository) DeleteMenuCategory(ctx context.Context, id string) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	intID, _ := strconv.Atoi(id)

	err := repo.Collection.FindOneAndDelete(
								ctx,
								bson.M{
									"$or":
										bson.A{
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
											"$or":
												bson.A{
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
				"$or":
					bson.A{
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

func (repo *menuCategoryMongoRepository) InsertMenu(ctx context.Context, menuCategoryId string, data *domain.Menu) (*domain.Menu, error) {
	var menucategory domain.MenuCategory
	var menu domain.Menu

	// check if menu category is exists
	_, err := repo.FindMenuCategory(ctx, menuCategoryId)
	if err != nil {
		return data, err
	}

	intMenuCategoryId, _ := strconv.Atoi(menuCategoryId)

	// add created_at
	data.CreatedAt = time.Now()

	// create a menu inside a collection (in 'menus' field)
	result, err := repo.Collection.UpdateOne(
											ctx,
											bson.M{
												"$or":
													bson.A{
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
		return data, err
	}

	// if there is document effected by update query then ..
	if result.MatchedCount == 1 {
		// search effected document
		fmt.Println("this is an result: ", result)
		err := repo.Collection.FindOne(
										ctx,
										bson.M{
											"$or":
												bson.A{
													bson.M{"uuid": menuCategoryId},
													bson.M{"id": intMenuCategoryId},
												},
										},
										options.FindOne().SetProjection(bson.M{"menus.$": 1}),
									).Decode(&menu)

		fmt.Printf("this is an what?: %v and %v", menucategory, err)

		// if not marshaled, then return error
		if err != nil {
			return data, err
		}
	}

	// return menu
	resp := domain.Menu{

	}


	return &resp, err
}