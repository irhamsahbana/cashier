package domain

import (
	"context"
	customtype "lucy/cashier/lib/custom_type"
)

type CompanyRepositoryContract interface {
	FindCompanyByUUID(ctx context.Context, companyId string, withTrashed bool) (*Company, int, error)
	FindCompanyByBranchUUID(ctx context.Context, branchId string, withTrashed bool) (*Company, int, error)
}

type Company struct {
	UUID      string   `bson:"uuid"`
	Name      string   `bson:"name"`
	Branches  []Branch `bson:"branches"`
	CreatedAt int64    `bson:"created_at"`
	UpdatedAt *int64   `bson:"updated_at"`
	DeletedAt *int64   `bson:"deleted_at"`
}

type Branch struct {
	UUID              string                        `bson:"uuid"`
	UniqueIndentifier *string                       `bson:"unique_identifier"`
	Name              string                        `bson:"name"`
	Preferences       []customtype.BranchPreference `bson:"preferences"`
	FeePreference     FeePreference                 `bson:"fee_preference"`
	PaymentMethods    []string                      `bson:"payment_methods"`
	Timezone          string                        `bson:"timezone"`
	Taxes             []Tax                         `bson:"taxes"`
	Tips              []Tip                         `bson:"tips"`
	Phone             string                        `bson:"phone"`
	Address           Address                       `bson:"address"`
	SocialMedia       SocialMedia                   `bson:"social_media"`
	Public            bool                          `bson:"public"`
	Password          string                        `bson:"password"`
	CreatedAt         int64                         `bson:"created_at"`
	UpdatedAt         *int64                        `bson:"updated_at"`
	DeletedAt         *int64                        `bson:"deleted_at"`
}

type Address struct {
	Province   string `bson:"province"`
	City       string `bson:"city"`
	Street     string `bson:"street"`
	PostalCode string `bson:"postal_code"`
}

type FeePreference struct {
	Service     Fee `bson:"service"`
	Queue       Fee `bson:"queue"`
	Reservation Fee `bson:"reservation"`
	Gojek       Fee `bson:"gojek"`
	Grab        Fee `bson:"grab"`
	Shopee      Fee `bson:"shopee"`
	Maxim       Fee `bson:"maxim"`
	Private     Fee `bson:"private"`
}

type Fee struct {
	Nominal    *float32 `bson:"nominal"`
	Percentage *float32 `bson:"percentage"`
}

type SocialMedia struct {
	Facebook   *string   `bson:"facebook"`
	Twitter    *string   `bson:"twitter"`
	Tiktok     *string   `bson:"tiktok"`
	Instagram  *string   `bson:"instagram"`
	GoogleMaps *string   `bson:"google_maps"`
	Whatsapp   *Whatsapp `bson:"whatsapp"`
}

type Tax struct {
	UUID        string  `bson:"uuid"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Value       float32 `bson:"value"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at"`
	DeletedAt   *int64  `bson:"deleted_at"`
}

type Tip struct {
	UUID        string  `bson:"uuid"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Value       float32 `bson:"value"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at"`
	DeletedAt   *int64  `bson:"deleted_at"`
}

type Whatsapp struct {
	CountryCode string `bson:"country_code"`
	Number      string `bson:"number"`
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
