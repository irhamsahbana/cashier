package bootstrap

import (
	"context"
	"fmt"
	"log"
	dummydata "lucy/cashier/bootstrap/dummy_data"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func initMongoDatabaseDummyData(ctx context.Context, client *mongo.Client, dbName string) {
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

		// "employee_shifts",

		"files",
	}

	// delete documents from collections
	for _, collection := range collections {
		color.Yellow(fmt.Sprintf("deleting documents from %s", collection) + " collection ...")
		_, err := client.Database(dbName).Collection(collection).DeleteMany(ctx, bson.M{})
		if err != nil {
			color.Red("MongoDB: " + err.Error() + " on collection " + collection)
			log.Fatal(err)
		}
	}

	// seeding dummy data
	dummydata.Seed(client.Database(dbName))
}
