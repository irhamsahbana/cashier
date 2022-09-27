package dummydata

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionBranchDiscount(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "f682490c-5a43-43f3-baae-e057eb423b1b"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Discount 1"},
			{Key: "description", Value: "bla bla bla"},
			{Key: "fixed", Value: nil},
			{Key: "percentage", Value: 12.5},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f682490c-5a43-43f3-baae-e057eb423b1c"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Discount 2"},
			{Key: "description", Value: "bli bli bli"},
			{Key: "fixed", Value: 10000},
			{Key: "percentage", Value: nil},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
	}

	_, err := coll.InsertMany(nil, data)
	if err != nil {
		log.Fatal("seed branch_discounts", err)
	}
}
