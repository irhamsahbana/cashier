package mongo

import (
	"context"
	"strconv"

	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	intID, _ := strconv.Atoi(id)

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

	return &menucategory, nil
}