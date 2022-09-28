package mongo

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/lib/logger"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

func (repo *zoneMongoRepository) Zones(ctx context.Context, branchId string) ([]domain.ZoneWithSpaceGroups, int, error) {
	var zonesWithGroups []domain.ZoneWithSpaceGroups

	filter := bson.M{"branch_uuid": branchId}

	pipeline := []bson.M{
		{"$match": filter},
		{"$lookup": bson.M{
			"from":         "space_groups",
			"localField":   "space_groups",
			"foreignField": "uuid",
			"as":           "space_groups",
		}},
		{"$project": bson.M{
			"uuid":         1,
			"branch_uuid":  1,
			"name":         1,
			"description":  1,
			"space_groups": 1,
			"created_at":   1,
			"updated_at":   1,
		}},
	}

	cursor, err := repo.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	if err = cursor.All(ctx, &zonesWithGroups); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	// find all space groups that not related to any zone, then make special zone for it
	spaceGroupInZone := []string{}

	for _, zone := range zonesWithGroups {
		for _, spaceGroup := range zone.SpaceGroups {
			spaceGroupInZone = append(spaceGroupInZone, spaceGroup.UUID)
		}
	}

	filter = bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": bson.M{"$nin": spaceGroupInZone}},
		},
	}

	var notInZone domain.ZoneWithSpaceGroups
	notInZone.UUID = "not-in-zone"
	notInZone.BranchUUID = branchId
	notInZone.Name = "Not In Zone"
	desc := "Space groups that not in any zone"
	notInZone.Description = &desc
	notInZone.CreatedAt = time.Now().UnixMicro()

	pipeline = []bson.M{
		{"$match": filter},
	}

	cursor, err = repo.CollectionSpaceGroup.Aggregate(ctx, pipeline)
	if err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	var spaceGroups []domain.SpaceGroup
	if err = cursor.All(ctx, &spaceGroups); err != nil {
		logger.Log(logrus.Fields{}).Error(err)
		return nil, http.StatusInternalServerError, err
	}

	notInZone.SpaceGroups = spaceGroups
	zonesWithGroups = append(zonesWithGroups, notInZone)

	return zonesWithGroups, http.StatusOK, nil
}
