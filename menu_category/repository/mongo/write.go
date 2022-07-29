package mongo

import (
	"context"
	"strconv"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
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

func (repo *menuCategoryMongoRepository) UpdateMenuCategory(ctx context.Context, id string, data *domain.MenuCategory) (*domain.MenuCategory, error) {
	var menucategory domain.MenuCategory

	_, err := repo.FindMenuCategory(ctx, id)
	if err != nil {
		return &menucategory, err
	}

	intID, _ := strconv.Atoi(id)

	update := bson.M{
		"name": data.Name,
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