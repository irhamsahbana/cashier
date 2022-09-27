package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionUserRole(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "2e60cc6e-afe5-412a-89bf-a85bbe166c2e"},
			{Key: "name", Value: "Owner"},
			{Key: "power", Value: 100},
			{Key: "created_at", Value: 1660403045123456},
			// {Key: "updated_at", Value: 1660403045123456},
			// {Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "144f16b6-3593-4331-9c64-51bbcf69f0c2"},
			{Key: "name", Value: "Head Accountant"},
			{Key: "power", Value: 40},
			{Key: "created_at", Value: 1660403045123456},
		},
		bson.D{
			{Key: "uuid", Value: "645ff83a-8cc5-4fe2-affb-ac79cc434535"},
			{Key: "name", Value: "Admin Cashier"},
			{Key: "power", Value: 40},
			{Key: "created_at", Value: 1660403045123456},
		},
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}
