package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type branchDiscountMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewBranchDiscountMongoRepository(DB mongo.Database) domain.BranchDiscountRepositoryContract {
	return &branchDiscountMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("branch_discounts"),
	}
}

func (r *branchDiscountMongoRepository) FindBranchDiscounts(ctx context.Context, branchId string) ([]domain.BranchDiscount, int, error) {
	var branchDiscounts []domain.BranchDiscount
	filter := bson.M{"branch_uuid": branchId}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("error while fetching branch discounts")
		return nil, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &branchDiscounts); err != nil {
		logger.Log(logrus.Fields{
			"error": err,
		}).Error("error while fetching branch discounts")
		return nil, http.StatusInternalServerError, err
	}

	return branchDiscounts, http.StatusOK, nil
}
