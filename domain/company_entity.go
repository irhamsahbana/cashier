package domain

import (
	"context"
)

type BranchPreference string

const (
	BranchPreference_TAKEAWAYS  BranchPreference = "takeaways"
	BranchPreference_DELIVERIES BranchPreference = "deliveries"
	BranchPreference_SPACES     BranchPreference = "spaces"
)

type Company struct {
	UUID      string   `bson:"uuid"`
	Name      string   `bson:"name"`
	Branches  []Branch `bson:"branches,omitempty"`
	CreatedAt int64    `bson:"created_at"`
	UpdatedAt *int64   `bson:"updated_at,omitempty"`
	DeletedAt *int64   `bson:"deleted_at,omitempty"`
}

type Branch struct {
	UUID              string             `bson:"uuid"`
	UniqueIndentifier *string            `bson:"unique_identifier,omitempty"`
	Name              string             `bson:"name"`
	Preferences       []BranchPreference `bson:"preferences"`
	PaymentMethods    []string           `bson:"payment_methods"`
	Timezone          string             `bson:"timezone"`
	Tax               float64            `bson:"tax"`
	Tip               float64            `bson:"tip"`
	Phone             *string            `bson:"phone,omitempty"`
	Instagram         *string            `bson:"instagram,omitempty"`
	Facebook          *string            `bson:"facebook,omitempty"`
	Twitter           *string            `bson:"twitter,omitempty"`
	GoogleMaps        *string            `bson:"google_maps,omitempty"`
	Public            bool               `bson:"public"`
	Password          string             `bson:"password"`
	CreatedAt         int64              `bson:"created_at"`
	UpdatedAt         *int64             `bson:"updated_at,omitempty"`
	DeletedAt         *int64             `bson:"deleted_at,omitempty"`
}

type CompanyRepository interface {
	UpsertCompany(ctx context.Context, company *Company) (*Company, int, error)
	InsertBranch(ctx context.Context, companyId string, branch *Branch) (*Company, int, error)
	UpdateBranch(ctx context.Context, branchId string, branch *Branch) (*Company, int, error)

	FindCompany(ctx context.Context, companyId string, withTrashed bool) (*Company, int, error)
	FindBranch(ctx context.Context, branchId string, withTrashed bool) (*Company, int, error)

	DeleteCompany(ctx context.Context, companyId string) (*Company, int, error)
	DeleteBranch(ctx context.Context, branchId string) (*Company, int, error)
}
