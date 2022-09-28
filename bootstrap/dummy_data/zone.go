package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionZone(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "d8b781e6-e7c6-4701-91be-daba256cf1dc"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Lantai 1"},
			{Key: "description", Value: "Indoor"},
			{Key: "space_groups", Value: bson.A{
				"c9c4c16b-34b5-4b7a-94a4-8c24cc4b5656",
				"267dcd73-d055-44bc-8782-cd91fb934953",
			}},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
		},

		bson.D{
			{Key: "uuid", Value: "1c4229d8-c690-4aef-b115-8e8c57042408"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Lantai 2"},
			{Key: "description", Value: "Indoor"},
			{Key: "space_groups", Value: bson.A{
				"e4eeaac6-b345-44c5-9e89-2c5edb46a421",
			}},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
		},
	}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}
