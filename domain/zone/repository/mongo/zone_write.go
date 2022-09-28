package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *zoneMongoRepository) UpsertZones(ctx context.Context, branchId string, data []domain.Zone) ([]domain.ZoneWithSpaceGroups, int, error) {
	zonesId := make([]string, 0)
	spaceGroupsId := make([]string, 0)

	for _, zone := range data {
		zonesId = append(zonesId, zone.UUID)

		for _, spaceGroup := range zone.SpaceGroups {
			spaceGroupsId = append(spaceGroupsId, spaceGroup)
		}

		filter := bson.M{
			"$and": []bson.M{
				{"branch_uuid": branchId},
				{"uuid": zone.UUID},
			},
		}

		countZone, err := repo.Collection.CountDocuments(ctx, filter)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		if len(zone.SpaceGroups) == 0 {
			zone.SpaceGroups = []string{}
		}

		if countZone == 0 {
			_, err := repo.Collection.InsertOne(ctx, zone)
			if err != nil {
				return nil, http.StatusInternalServerError, err
			}
		} else {
			updatedAt := time.Now().UnixMicro()
			zone.UpdatedAt = &updatedAt
			_, err := repo.Collection.UpdateOne(ctx, bson.M{"uuid": zone.UUID}, bson.M{"$set": zone})
			if err != nil {
				return nil, http.StatusInternalServerError, err
			}
		}
	}

	// Delete zones that not in the request
	filter := bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": bson.M{"$nin": zonesId}},
		},
	}

	_, err := repo.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	filter = bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": bson.M{"$in": zonesId}},
		},
	}

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
		return nil, http.StatusInternalServerError, err
	}

	var zones []domain.ZoneWithSpaceGroups
	if err = cursor.All(ctx, &zones); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// find all space groups that not related to any zone, then make special zone for it
	filter = bson.M{
		"$and": []bson.M{
			{"branch_uuid": branchId},
			{"uuid": bson.M{"$nin": spaceGroupsId}},
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
		return nil, http.StatusInternalServerError, err
	}

	var spaceGroups []domain.SpaceGroup
	if err = cursor.All(ctx, &spaceGroups); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	notInZone.SpaceGroups = spaceGroups
	zones = append(zones, notInZone)

	return zones, http.StatusOK, nil
}
