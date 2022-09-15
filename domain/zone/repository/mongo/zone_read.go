package mongo

import (
	"lucy/cashier/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type zoneMongoRepository struct {
	DB                   mongo.Database
	Collection           mongo.Collection
	CollectionSpaceGroup mongo.Collection
}

func NewZoneMongoRepository(DB mongo.Database) domain.ZoneRepositoryContract {
	return &zoneMongoRepository{
		DB:                   DB,
		Collection:           *DB.Collection("zones"),
		CollectionSpaceGroup: *DB.Collection("space_groups"),
	}
}
