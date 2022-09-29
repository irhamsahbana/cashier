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
			{Key: "email", Value: "tony@stark.industries.com"},
			{Key: "password", Value: "$2a$10$vvI.Ym4N5zzqk6QuT.BCPO9A81WhnhhCaG9zDDHK2MwWyRLq2pmPO"}, // password: password
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "3ff9da64-1a97-438f-903a-e2ed99617a25"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "role_uuid", Value: "645ff83a-8cc5-4fe2-affb-ac79cc434535"}, // admin cashier
			{Key: "name", Value: "Steve Rogers"},
			{Key: "email", Value: "steve@rogers.com"},
			{Key: "password", Value: "$2a$10$3Vp2sksag/DtVRPTLUT2ueEa1.BCr7epx/BzZZLrQiou7GIol4EIC"}, // password: password
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f38a126a-a3e6-445a-85d6-9d92823213a9"},
			{Key: "branch_uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
			{Key: "role_uuid", Value: "f6118477-a4d4-45d9-a7ff-36bd17be0bb0"}, // cashier
			{Key: "name", Value: "Bruce Banner"},
			{Key: "email", Value: "bruce@banner.com"},
			{Key: "password", Value: "$2a$10$HEDVVS1DQ3R9INQyFnlUlu.Eg/w.CBupsYATBTixV/3fRLz9edEoC"}, // password: password
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
