package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDatabaseIndexes(ctx context.Context, client *mongo.Client, dbName string) {
	var collections []string = []string{
		"users",
		"user_roles",
		"tokens",

		"item_categories",
		"space_groups",
		"zones",
		"waiters",

		"employee_shifts",

		"files",
	}

	//delete all indexes first
	for _, collection := range collections {
		color.Yellow(fmt.Sprintf("deleting indexes from %s", collection) + " collection ...")
		_, err := client.Database(dbName).Collection(collection).Indexes().DropAll(ctx)
		if err != nil {
			color.Red("MongoDB: " + err.Error() + " on collection " + collection)
			log.Fatal(err)
		}
	}

	//create indexes
	for _, collection := range collections {
		switch collection {
		case "users":
			color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
				{
					Keys:    bson.M{"email": 1},
					Options: options.Index().SetUnique(true),
				},
			})

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

			if err != nil {
				color.Red("MongoDB: " + err.Error() + " on collection " + collection)
				log.Fatal(err)
			}

		case "employee_shifts":
			color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
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

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

			if err != nil {
				color.Red("MongoDB: " + err.Error() + " on collection " + collection)
				log.Fatal(err)
			}

		case "item_categories":
			color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

			if err != nil {
				color.Red("MongoDB: " + err.Error() + " on collection " + collection)
				log.Fatal(err)
			}

		case "space_groups":
			color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
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

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

			if err != nil {
				color.Red("MongoDB: " + err.Error() + " on collection " + collection)
				log.Fatal(err)
			}

		case "zones":
			color.Cyan(fmt.Sprintf("creating indexes for %s", collection) + " collection ...")
			res, err := client.Database(dbName).Collection(collection).Indexes().CreateMany(ctx, []mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "branch_uuid", Value: 1},
						{Key: "uuid", Value: 1},
					},
					Options: options.Index().SetUnique(true),
				},
			})

			for _, index := range res {
				color.Green(fmt.Sprintf("index %s created", index))
			}

			if err != nil {
				color.Red("MongoDB: " + err.Error() + " on collection " + collection)
				log.Fatal(err)
			}
		}
	}
}
