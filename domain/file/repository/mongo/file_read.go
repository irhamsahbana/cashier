package mongo

import (
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type fileMongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewFileMongoRepository(DB mongo.Database) domain.FileRepositoryContract {
	return &fileMongoRepository{
		DB:         DB,
		Collection: *DB.Collection("files"),
	}
}
