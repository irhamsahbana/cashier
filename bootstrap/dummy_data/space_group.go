package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionSpaceGroup(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "c9c4c16b-34b5-4b7a-94a4-8c24cc4b5656"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Meja Bulat"},
			{Key: "spaces", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "5c102c7c-da1e-4f1f-9988-d46c0de0280e"},
					{Key: "number", Value: 1},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 1"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
				bson.D{
					{Key: "uuid", Value: "547a5166-fdf5-452a-a2b8-ed2b4439ada9"},
					{Key: "number", Value: 2},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 2"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
			}},
			{Key: "code", Value: "MB"},
			{Key: "shape", Value: "circle"},
			{Key: "length", Value: 1},
			{Key: "pax", Value: 10},
			{Key: "reserveable", Value: true},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},

		bson.D{
			{Key: "uuid", Value: "267dcd73-d055-44bc-8782-cd91fb934953"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Meja Kotak"},
			{Key: "spaces", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "e34890f6-44f6-43d3-a1e1-98a8461b1b86"},
					{Key: "number", Value: 3},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 3"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
				bson.D{
					{Key: "uuid", Value: "209e2c1d-3bc8-4f94-bf75-abaf4b8a51ff"},
					{Key: "number", Value: 4},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 4"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
			}},
			{Key: "code", Value: "MK"},
			{Key: "shape", Value: "square"},
			{Key: "length", Value: 2},
			{Key: "pax", Value: 20},
			{Key: "reserveable", Value: true},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},

		bson.D{
			{Key: "uuid", Value: "e4eeaac6-b345-44c5-9e89-2c5edb46a421"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "name", Value: "Meeting Room"},
			{Key: "spaces", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "8b3b8dd7-bf5f-4cdc-9c1d-1001e02747ec"},
					{Key: "number", Value: 1},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 1"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
				bson.D{
					{Key: "uuid", Value: "be47014f-6504-4df1-9105-bf530e2296ce"},
					{Key: "number", Value: 2},
					{Key: "occupied", Value: false},
					{Key: "description", Value: "Meja 2"},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
			}},
			{Key: "code", Value: "MR"},
			{Key: "shape", Value: "circle"},
			{Key: "length", Value: 1},
			{Key: "pax", Value: 10},
			{Key: "reserveable", Value: true},
			{Key: "disabled", Value: false},
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
