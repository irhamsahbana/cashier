package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionUser(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "a2b53413-8920-4f44-b4f5-bea3189f92ec"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "role_uuid", Value: "2e60cc6e-afe5-412a-89bf-a85bbe166c2e"},
			{Key: "name", Value: "Tony Stark"},
			{Key: "email", Value: "tony@starkindustry.com"},
			{Key: "password", Value: "$2a$10$vvI.Ym4N5zzqk6QuT.BCPO9A81WhnhhCaG9zDDHK2MwWyRLq2pmPO"}, // password: password
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}
