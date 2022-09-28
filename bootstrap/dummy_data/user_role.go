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
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "d8569961-fae4-4338-a617-0b663af3ae12"},
			{Key: "name", Value: "Branch Owner"},
			{Key: "power", Value: 80},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f69f078a-eb79-42b8-81a6-9ad3b52e47a9"},
			{Key: "name", Value: "Manager"},
			{Key: "power", Value: 60},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "144f16b6-3593-4331-9c64-51bbcf69f0c2"},
			{Key: "name", Value: "Head Accountant"},
			{Key: "power", Value: 40},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "645ff83a-8cc5-4fe2-affb-ac79cc434535"},
			{Key: "name", Value: "Admin Cashier"},
			{Key: "power", Value: 40},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f818c877-603b-4e7c-9aa9-a760d10b053d"},
			{Key: "name", Value: "Stock Overseer"},
			{Key: "power", Value: 40},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "a81d48da-b583-4034-9aec-9ea6eca562f6"},
			{Key: "name", Value: "Accountant"},
			{Key: "power", Value: 20},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f6118477-a4d4-45d9-a7ff-36bd17be0bb0"},
			{Key: "name", Value: "Cashier"},
			{Key: "power", Value: 20},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "b94557e6-d115-4235-b9c2-93f06406d2aa"},
			{Key: "name", Value: "Stock Keeper"},
			{Key: "power", Value: 20},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "dda76633-d516-46e3-8c36-7bfae9366df5"},
			{Key: "name", Value: "Customer"},
			{Key: "power", Value: 0},
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
