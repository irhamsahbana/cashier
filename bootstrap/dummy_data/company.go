package dummydata

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func collectionCompany(coll *mongo.Collection) {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "fd58bfcf-e95e-4cfa-8789-7fcc9d5e046c"},
			{Key: "name", Value: "Stark Industry"},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
			{Key: "branches", Value: bson.A{
				bson.D{
					{Key: "uuid", Value: "b0e91f44-2481-4eea-899b-93387a828155"},
					{Key: "name", Value: "Stark Industry - Head Office"},
					{Key: "address", Value: bson.D{
						{Key: "province", Value: "Sulawesi Selatan"},
						{Key: "city", Value: "Kota Makassar"},
						{Key: "street", Value: "Jl. Jendral Sudirman No. 1"},
						{Key: "postal_code", Value: "90245"},
					}},
					{Key: "preferences", Value: bson.A{
						"queues", "deliveries", "spaces",
					}},
					{Key: "social_media", Value: bson.D{
						{Key: "facebook", Value: "https://www.facebook.com/starkindustry"},
						{Key: "twitter", Value: "https://www.twitter.com/starkindustry"},
						{Key: "tiktok", Value: "https://www.tiktok.com/starkindustry"},
						{Key: "instagram", Value: "https://www.instagram.com/starkindustry"},
						{Key: "google_maps", Value: "https://www.google.com/maps/starkindustry"},
						{Key: "whatsapp", Value: bson.D{
							{Key: "country_code", Value: "62"},
							{Key: "number", Value: "81234567890"},
						}},
					}},
					{Key: "fee_preference", Value: feePreference()},
					{Key: "taxes", Value: bson.A{
						bson.D{
							{Key: "uuid", Value: "be8c37d0-e5f8-4c5d-aef6-c0dbc4311f1f"},
							{Key: "name", Value: "Pajak Restoran"},
							{Key: "description", Value: "Pajak Restoran"},
							{Key: "value", Value: 11},
							{Key: "is_default", Value: true},
							{Key: "created_at", Value: 1660403045123456},
							{Key: "updated_at", Value: 1660403045123456},
							{Key: "deleted_at", Value: nil},
						},
					}},
					{Key: "tips", Value: bson.A{
						bson.D{
							{Key: "uuid", Value: "28429b57-df0d-4e57-92ab-14dc24648458"},
							{Key: "name", Value: "Tip Waiter"},
							{Key: "description", Value: "Tip Waiter"},
							{Key: "value", Value: 11},
							{Key: "is_default", Value: false},
							{Key: "created_at", Value: 1660403045123456},
							{Key: "updated_at", Value: 1660403045123456},
							{Key: "deleted_at", Value: nil},
						},
					}},
					{Key: "payment_methods", Value: paymentMethods()},
					{Key: "phone", Value: "14045"},
					{Key: "timezone", Value: "Asia/Makassar"},
					{Key: "public", Value: true},
					{Key: "created_at", Value: 1660403045123456},
					{Key: "updated_at", Value: 1660403045123456},
					{Key: "deleted_at", Value: nil},
				},
			},
			},
		}}

	_, err := coll.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func feePreference() bson.D {
	data := bson.D{
		{Key: "service", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: 11},
		}},
		{Key: "queue", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "reservation", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "gojek", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "grab", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "shopee", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "maxim", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
		{Key: "private", Value: bson.D{
			{Key: "nominal", Value: nil},
			{Key: "percentage", Value: nil},
		}},
	}

	return data
}

func paymentMethods() bson.A {
	data := bson.A{
		bson.D{
			{Key: "uuid", Value: "981fddcb-8e10-42ba-a77a-850ae0169c56"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "cash"},
			{Key: "name", Value: "Cash"},
			{Key: "description", Value: "Pembayaran Tunai"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "78fdd1c4-d3d6-491e-8a96-9e046ec06c21"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "edc"},
			{Key: "name", Value: "EDC BCA"},
			{Key: "description", Value: "EDC BCA"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "d8f666c1-107b-41ad-b1ed-a8e24c864f30"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "edc"},
			{Key: "name", Value: "EDC BNI"},
			{Key: "description", Value: "EDC BNI"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f1be1299-69b2-4c23-978a-a749af2df96e"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "edc"},
			{Key: "name", Value: "EDC BRI"},
			{Key: "description", Value: "EDC BRI"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "bec37cae-cf2e-41c8-83e9-736782563fd8"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "delivery"},
			{Key: "name", Value: "Gojek"},
			{Key: "description", Value: ""},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "bfe4c6fa-9007-498d-962e-de944e3eae49"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "delivery"},
			{Key: "name", Value: "Grab"},
			{Key: "description", Value: ""},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "4d15a9e7-87d4-44f2-a98d-6bf458d873f5"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "delivery"},
			{Key: "name", Value: "Shopee"},
			{Key: "description", Value: ""},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "c586a1d3-a3ff-4660-bea5-417342b720cd"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "e-wallet"},
			{Key: "name", Value: "GoPay"},
			{Key: "description", Value: "Gojek Pay"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "5246ea32-fa8e-490a-bb85-e25a25a4fc26"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "e-wallet"},
			{Key: "name", Value: "Shopee Pay"},
			{Key: "description", Value: "Shopee E-Wallet"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
		bson.D{
			{Key: "uuid", Value: "f44bd63f-b061-4592-8257-5402c7d6304a"},
			{Key: "entry_uuid", Value: nil},
			{Key: "group", Value: "e-wallet"},
			{Key: "name", Value: "OVO"},
			{Key: "description", Value: "OVO E-Wallet"},
			{Key: "fee", Value: bson.D{
				{Key: "fixed", Value: 0.00},
				{Key: "percent", Value: 0.00},
			}},
			{Key: "disabled", Value: false},
			{Key: "created_at", Value: 1660403045123456},
			{Key: "updated_at", Value: 1660403045123456},
			{Key: "deleted_at", Value: nil},
		},
	}

	return data
}
