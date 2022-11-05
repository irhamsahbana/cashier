package bootstrap

import (
	"context"
	"fmt"
	"log"
	"lucy/cashier/lib/helper"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDatabaseIndexes(ctx context.Context, client *mongo.Client, dbName string) {
	var collections []string = []string{
		"companies",
		"branch_discounts",

		"users",
		"user_roles",
		"tokens",

		"item_categories",
		"space_groups",
		"zones",
		"waiters",

		"employee_shifts",

		"active_order_groups",
		"invoices",

		"files",
	}

	// get all collections
	collNames, err := client.Database(dbName).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		App.Log.Error(err)
		log.Fatal(err)
	}

	// create collection if not exists
	for _, collName := range collections {
		if !helper.Contain(collNames, collName) {
			err = client.Database(dbName).CreateCollection(ctx, collName)
			if err != nil {
				App.Log.Error(err)
				log.Fatal(err)
			}
			color.Cyan(fmt.Sprintf("Collection %s created", collName))
		}
	}

	// delete all indexes first
	for _, collection := range collections {
		color.Yellow(fmt.Sprintf("deleting indexes from %s", collection) + " collection ...")
		_, err := client.Database(dbName).Collection(collection).Indexes().DropAll(ctx)
		if err != nil {
			color.Red("MongoDB: " + err.Error() + " on collection " + collection)
			App.Log.Error(err)
			log.Fatal(err)
		}
	}

	// create indexes
	for _, collection := range collections {
		switch collection {
		case "users":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "email", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "phone", Value: 1},
					},
					// set unique index to non null phone
					Options: options.Index().SetUnique(true).SetPartialFilterExpression(
						bson.D{
							{Key: "phone", Value: bson.D{{Key: "$type", Value: "string"}}},
						},
					),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "employee_shifts":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "user_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "supporters.uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true).SetSparse(true),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "item_categories":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "space_groups":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "spaces.uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true).SetSparse(true),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "zones":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "active_order_groups":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)

		case "invoices":
			createCollectionIndex(collection)
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys: bson.D{
						{Key: "order_groups.uuid", Value: 1},
					},
				},
				{
					Keys: bson.D{
						{Key: "order_groups.orders.uuid", Value: 1},
					},
				},
			})
			errCollectionIndexingCheck(err, collection)
			notifyCollectionIndexesCreated(res)
		}
	}
}

func createCollectionIndex(collection string) {
	color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
}

func errCollectionIndexingCheck(err error, collection string) {
	if err != nil {
		color.Red("MongoDB: " + err.Error() + " on collection " + collection)
		App.Log.Error(err)
		log.Fatal(err)
	}
}

func notifyCollectionIndexesCreated(res []string) {
	for _, index := range res {
		color.Green(fmt.Sprintf("index %s created", index))
	}
}
