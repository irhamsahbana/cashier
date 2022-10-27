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
	PaymentMethods    []PaymentMethod               `bson:"payment_methods"`
	Taxes             []Tax                         `bson:"taxes"`
	Tips              []Tip                         `bson:"tips"`
	Address           Address                       `bson:"address"`
	FeePreference     FeePreference                 `bson:"fee_preference"`
	SocialMedia       SocialMedia                   `bson:"social_media"`
	Phone             string                        `bson:"phone"`
	Timezone          string                        `bson:"timezone"`
	Public            bool                          `bson:"public"`
	Password          string                        `bson:"password"`
	CreatedAt         int64                         `bson:"created_at"`
	UpdatedAt         *int64                        `bson:"updated_at"`
	DeletedAt         *int64                        `bson:"deleted_at"`
}

type PaymentMethod struct {
	UUID        string                        `bson:"uuid"`
	EntryUUID   *string                       `bson:"entry_uuid"`
	Group       customtype.PaymentMethodGroup `bson:"group"`
	Name        string                        `bson:"name"`
	Fee         PaymentMethodFee              `bson:"fee"`
	Description string                        `bson:"description"`
	Disabled    bool                          `bson:"disabled"`
	CreatedAt   int64                         `bson:"created_at"`
	UpdatedAt   *int64                        `bson:"updated_at"`
	DeletedAt   *int64                        `bson:"deleted_at"`
}

type PaymentMethodFee struct {
	Fixed   float64 `bson:"fixed"`
	Percent float64 `bson:"percent"`
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
	Nominal    *float64 `bson:"nominal"`
	Percentage *float64 `bson:"percentage"`
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
	Value       float64 `bson:"value"`
	IsDefault   bool    `bson:"is_default"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at"`
	DeletedAt   *int64  `bson:"deleted_at"`
}

type Tip struct {
	UUID        string  `bson:"uuid"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Value       float64 `bson:"value"`
	IsDefault   bool    `bson:"is_default"`
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
