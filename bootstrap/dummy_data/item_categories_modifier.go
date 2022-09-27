package dummydata

import (
	"go.mongodb.org/mongo-driver/bson"
)

// modifier for coffee based
func modifierGroups1() bson.A {
	data := bson.A{
		modifierGroup1(),
		modifierGroup2(),
	}

	return data
}

func modifierGroup1() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "2d09500f-eed9-471d-956e-d15f34cbd60f"},
			{Key: "name", Value: "Topping"},
			{Key: "modifiers", Value: modifierGroup1Modifiers()},
			{Key: "max_quantity", Value: 1},
			{Key: "min_quantity", Value: 1},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}

func modifierGroup1Modifiers() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "8d76153a-dadd-42d4-9238-b57a25547d17"},
			{Key: "name", Value: "Regal"},
			{Key: "price", Value: 5000},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "5f95cd4f-a249-47c0-942d-9e61e83fa39f"},
			{Key: "name", Value: "Rum"},
			{Key: "price", Value: 7000},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "54d83c45-e751-47ee-939a-cfd1657f19ca"},
			{Key: "name", Value: "Caramel"},
			{Key: "price", Value: 8000},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}

func modifierGroup2() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "064532a1-4c00-4d96-9342-824710d4be56"},
			{Key: "name", Value: "sugar"},
			{Key: "modifiers", Value: modifierGroup2Modifiers()},
			{Key: "max_quantity", Value: 1},
			{Key: "min_quantity", Value: 1},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}

func modifierGroup2Modifiers() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "86f390b1-5c82-4394-8d4d-6622d9a538ef"},
			{Key: "name", Value: "Less"},
			{Key: "price", Value: 0},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "05e76ea5-4c1f-4e1f-9d0a-49ac0e46a3af"},
			{Key: "name", Value: "Normal"},
			{Key: "price", Value: 0},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "73365ec8-574f-4a8f-b24e-b73708180efd"},
			{Key: "name", Value: "More"},
			{Key: "price", Value: 1500},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}
