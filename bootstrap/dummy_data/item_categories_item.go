package dummydata

import "go.mongodb.org/mongo-driver/bson"

func items1() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "427bba53-66a0-4a16-a402-a7b81e2b8b41"},
			{Key: "name", Value: "Cappuccino"},
			{Key: "price", Value: 23000},
			{Key: "label", Value: "Original"},
			{Key: "variants", Value: bson.A{}},
			{Key: "description", Value: "Cappuccino is a coffee-based drink that originated in Italy, and is traditionally prepared with steamed milk foam (microfoam)."},
			{Key: "image_path", Value: nil},
			{Key: "public", Value: true},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "acf6a93c-471e-4e81-bd17-a1850cb54fd6"},
			{Key: "name", Value: "Latte"},
			{Key: "price", Value: 25000},
			{Key: "label", Value: ""},
			{Key: "variants", Value: bson.A{}},
			{Key: "description", Value: "Latte is a coffee-based drink made primarily from espresso and steamed milk. It is typically smaller in volume than a cappuccino."},
			{Key: "image_path", Value: nil},
			{Key: "public", Value: true},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}

func items2() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "b67a71bb-1a3a-4e14-9701-60a173e383b6"},
			{Key: "name", Value: "Tim Tuna"},
			{Key: "price", Value: 33000},
			{Key: "label", Value: "Kuah"},
			{Key: "variants", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "3de19cf9-3ae5-4e4c-933b-f190ba6c1962"},
					{Key: "label", Value: "Spesial"},
					{Key: "price", Value: 38000},
					{Key: "image_path", Value: nil},
					{Key: "public", Value: true},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: nil},
				},
			}},
			{Key: "description", Value: ""},
			{Key: "image_path", Value: nil},
			{Key: "public", Value: true},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},

		bson.D{
			{Key: "uuid", Value: "ee9a152f-ed71-4f18-a3fc-82a1dab7251b"},
			{Key: "name", Value: "Tuna Goreng"},
			{Key: "price", Value: 33000},
			{Key: "label", Value: "Polos"},
			{Key: "variants", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "7d64a0c7-8a6c-42f3-b8cc-02b74f540967"},
					{Key: "label", Value: "Saos Woku"},
					{Key: "price", Value: 38000},
					{Key: "image_path", Value: nil},
					{Key: "public", Value: true},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: nil},
				},
			}},
			{Key: "description", Value: ""},
			{Key: "image_path", Value: nil},
			{Key: "public", Value: true},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},

		bson.D{
			{Key: "uuid", Value: "051dacd7-efef-4809-b177-0b6560ddee26"},
			{Key: "name", Value: "Shasimi Tuna Katsu"},
			{Key: "price", Value: 65000},
			{Key: "label", Value: ""},
			{Key: "variants", Value: bson.A{}},
			{Key: "description", Value: ""},
			{Key: "image_path", Value: nil},
			{Key: "public", Value: true},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: nil},
		},
	}

	return data
}
