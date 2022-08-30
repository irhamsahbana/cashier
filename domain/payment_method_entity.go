package domain

import "time"

type PaymentMethodGroup string

const (
	PaymentMethodGroup_CASH     PaymentMethodGroup = "cash"
	PaymentMethodGroup_DELIVERY PaymentMethodGroup = "delivery"
	PaymentMethodGroup_EDC      PaymentMethodGroup = "edc"
	PaymentMethodGroup_EWALLET  PaymentMethodGroup = "e-wallet"
	PaymentMethodGroup_QRIS     PaymentMethodGroup = "qris"
)

type PaymentMethod struct {
	UUID        string             `bson:"uuid"`
	Group       PaymentMethodGroup `bson:"group"`
	Name        string             `bson:"name"`
	Fee         float64            `bson:"fee"`
	Description *string            `bson:"description,omitempty"`
	Disabled    bool               `bson:"disabled"`
	CreatedAt   int64              `bson:"created_at"`
	UpdatedAt   *int64             `bson:"updated_at,omitempty"`
	DeletedAt   *int64             `bson:"deleted_at,omitempty"`
}

type PaymentMethodModel struct {
	UUID        string             `bson:"uuid" json:"uuid"`
	Group       PaymentMethodGroup `bson:"group" json:"group"`
	Name        string             `bson:"name" json:"name"`
	Fee         float64            `bson:"fee" json:"fee"`
	Description *string            `bson:"description,omitempty" json:"description,omitempty"`
	Disabled    bool               `bson:"disabled" json:"disabled"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
