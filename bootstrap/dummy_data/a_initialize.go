package dummydata

import (
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
)

func Seed(DB *mongo.Database) error {
	collectionCompany(DB.Collection("companies"))
	color.Green("Dummy companies data seeded")

	collectionBranchDiscount(DB.Collection("branch_discounts"))
	color.Green("Dummy branch_discounts data seeded")

	collectionSpaceGroup(DB.Collection("space_groups"))
	color.Green("Dummy space_groups data seeded")

	collectionZone(DB.Collection("zones"))
	color.Green("Dummy zones data seeded")

	collectionItemCategory(DB.Collection("item_categories"))
	color.Green("Dummy item_categories data seeded")

	collectionUserRole(DB.Collection("user_roles"))
	color.Green("Dummy user_roles data seeded")

	collectionUser(DB.Collection("users"))
	color.Green("Dummy users data seeded")

	return nil
}
