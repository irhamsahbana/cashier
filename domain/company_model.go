package domain

import "time"

type CompanyModel struct {
	UUID      string        `bson:"uuid" json:"uuid"`
	Name      string        `bson:"name" json:"name"`
	Branches  []BranchModel `bson:"branches,omitempty" json:"branches,omitempty"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time    `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time    `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type BranchModel struct {
	UUID              string               `bson:"uuid" json:"uuid"`
	Company           CompanyModel         `bson:"company" json:"company"`
	UniqueIndentifier *string              `bson:"unique_identifier,omitempty" json:"unique_identifier,omitempty"`
	Name              string               `bson:"name" json:"name"`
	Preferences       []BranchPreference   `bson:"preferences" json:"preferences"`
	PaymentMethods    []PaymentMethodModel `bson:"payment_methods" json:"payment_methods"`
	Timezone          string               `bson:"timezone" json:"timezone"`
	Tax               float64              `bson:"tax" json:"tax"`
	Tip               float64              `bson:"tip" json:"tip"`
	Phone             *string              `bson:"phone,omitempty" json:"phone,omitempty"`
	Instagram         *string              `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Facebook          *string              `bson:"facebook,omitempty" json:"facebook,omitempty"`
	Twitter           *string              `bson:"twitter,omitempty" json:"twitter,omitempty"`
	GoogleMaps        *string              `bson:"google_maps,omitempty" json:"google_maps,omitempty"`
	Public            bool                 `bson:"public" json:"public"`
	Password          string               `bson:"password" json:"password"`
	CreatedAt         time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt         *time.Time           `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt         *time.Time           `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
