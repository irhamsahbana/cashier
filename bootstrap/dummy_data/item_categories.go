package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionItemCategory(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "73302ff5-6c97-40ef-813a-94b4a906a065"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Coffee Based"},
			{Key: "modifier_groups", Value: modifierGroups1()},
			{Key: "items", Value: items1()},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "56377297-fe33-4a38-a2d4-afe6d97fc2ad"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Main Course"},
			{Key: "modifier_groups", Value: bson.A{}},
			{Key: "items", Value: items2()},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}
