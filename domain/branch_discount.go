package domain

import "context"

type BranchDiscount struct {
	UUID        string  `bson:"uuid"`
	BranchUUID  string  `bson:"branch_uuid"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Fixed       float32 `bson:"fixed"`
	Percentage  float32 `bson:"percentage"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at"`
	DeletedAt   *int64  `bson:"deleted_at"`
}

type BranchDiscountRepositoryContract interface {
	FindBranchDiscounts(ctx context.Context, branchUUID string) ([]BranchDiscount, int, error)
}
