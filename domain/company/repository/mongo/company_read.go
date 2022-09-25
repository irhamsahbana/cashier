package mongo

import (
	"context"
	"errors"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewCompanyMongoRepository(DB mongo.Database) domain.CompanyRepositoryContract {
	return &companyMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("companies"),
	}
}

func (r *companyMongoRepository) FindCompanyByUUID(ctx context.Context, id string, withTrashed bool) (*domain.Company, int, error) {
	var filter bson.M
	var company domain.Company

	if withTrashed {
		filter = bson.M{"uuid": id}
	} else {
		filter = bson.M{"uuid": id, "deleted_at": nil}
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{
				"uuid": id,
			}).Error("company not found")
			return nil, http.StatusNotFound, errors.New("company not found")
		}

		logger.Log(logrus.Fields{
			"uuid": id,
		}).Error("error when find company")
		return nil, http.StatusInternalServerError, err
	}

	return &company, http.StatusOK, nil
}

func (r *companyMongoRepository) FindCompanyByBranchUUID(ctx context.Context, id string, withTrashed bool) (*domain.Company, int, error) {
	var filter bson.M
	var company domain.Company

	if withTrashed {
		filter = bson.M{"branches.uuid": id}
	} else {
		filter = bson.M{"branches.uuid": id, "deleted_at": nil}
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Log(logrus.Fields{
				"uuid": id,
			}).Error("company not found")
			return nil, http.StatusNotFound, errors.New("company not found")
		}

		logger.Log(logrus.Fields{
			"uuid": id,
		}).Error("error when find company")
		return nil, http.StatusInternalServerError, err
	}

	return &company, http.StatusOK, nil
}
