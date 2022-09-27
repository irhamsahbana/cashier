package dummydata

import (
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
)

func Seed(DB *mongo.Database) error {
	company := DB.Collection("companies")
	collectionCompany(company)
	color.Green("Dummy companies data seeded")

	branchDiscount := DB.Collection("branch_discounts")
	collectionBranchDiscount(branchDiscount)
	color.Green("Dummy branch_discounts data seeded")

	itemCategory := DB.Collection("item_categories")
	collectionItemCategory(itemCategory)
	color.Green("Dummy item_categories data seeded")

	userRole := DB.Collection("user_roles")
	collectionUserRole(userRole)
	color.Green("Dummy user_roles data seeded")

	user := DB.Collection("users")
	collectionUser(user)
	color.Green("Dummy users data seeded")

	return nil
}
