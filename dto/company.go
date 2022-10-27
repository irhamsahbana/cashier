package dto

import (
	customtype "lucy/cashier/lib/custom_type"
	"time"
)

type CompanyResponse struct {
	UUID      string           `json:"uuid"`
	Name      string           `json:"name"`
	Branches  []BranchResponse `json:"branches,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt *time.Time       `json:"updated_at"`
	DeletedAt *time.Time       `json:"deleted_at"`
}

type BranchResponse struct {
	UUID              string                        `json:"uuid"`
	UniqueIndentifier *string                       `json:"unique_identifier"`
	Name              string                        `json:"name"`
	Company           CompanyResponse               `json:"company"`
	Preferences       []customtype.BranchPreference `json:"preferences"`
	PaymentMethods    []PaymentMethodResponse       `json:"payment_methods"`
	Taxes             []TaxResponse                 `json:"taxes"`
	Tips              []TipResponse                 `json:"tips"`
	Discounts         []BranchDiscountResponse      `json:"discounts"`
	Employees         []UserResponse                `json:"employees"`
	Address           AddressResponse               `json:"address"`
	SocialMedia       SocialMediaResponse           `json:"social_media"`
	FeePreference     FeePreferenceResponse         `json:"fee_preference"`
	Phone             string                        `json:"phone"`
	Timezone          string                        `json:"timezone"`
	Public            bool                          `json:"public"`
	CreatedAt         time.Time                     `json:"created_at"`
	UpdatedAt         *time.Time                    `json:"updated_at"`
	DeletedAt         *time.Time                    `json:"deleted_at"`
}

type PaymentMethodResponse struct {
	UUID        string                        `json:"uuid"`
	EntryUUID   *string                       `json:"entry_uuid"`
	Group       customtype.PaymentMethodGroup `json:"group"`
	Name        string                        `json:"name"`
	Fee         PaymentMethodFeeResponse      `json:"fee"`
	Description string                        `json:"description"`
	Disabled    bool                          `json:"disabled"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   *time.Time                    `json:"updated_at"`
	DeletedAt   *time.Time                    `json:"deleted_at"`
}

type PaymentMethodFeeResponse struct {
	Fixed   float64 `json:"fixed"`
	Percent float64 `json:"percent"`
}

type BranchDiscountResponse struct {
	UUID        string     `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Fixed       float64    `json:"fixed"`
	Percentage  float64    `json:"percentage"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type AddressResponse struct {
	Province   string `json:"province"`
	City       string `json:"city"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
}

type SocialMediaResponse struct {
	Facebook   *string           `json:"facebook"`
	Twitter    *string           `json:"twitter"`
	Tiktok     *string           `json:"tiktok"`
	Instagram  *string           `json:"instagram"`
	GoogleMaps *string           `json:"google_maps"`
	Whatsapp   *WhatsappResponse `json:"whatsapp"`
}

type TaxResponse struct {
	UUID        string     `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Value       float64    `json:"value"`
	IsDefault   bool       `json:"is_default"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type TipResponse struct {
	UUID        string     `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Value       float64    `json:"value"`
	IsDefault   bool       `json:"is_default"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type FeePreferenceResponse struct {
	Service     FeeResponse `json:"service"`
	Queue       FeeResponse `json:"queue"`
	Reservation FeeResponse `json:"reservation"`
	Gojek       FeeResponse `json:"gojek"`
	Grab        FeeResponse `json:"grab"`
	Shopee      FeeResponse `json:"shopee"`
	Maxim       FeeResponse `json:"maxim"`
	Private     FeeResponse `json:"private"`
}

type FeeResponse struct {
	Nominal    *float64 `json:"nominal"`
	Percentage *float64 `json:"percentage"`
}

type WhatsappResponse struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
