package domain

import (
	"context"
	"time"
)

type VerifiedableType string

const (
	VerifiedableType_USER     VerifiedableType = "user"
	VerifiedableType_EMPLOYEE VerifiedableType = "employee"
)

type EmailVerificationCode struct {
	UUID             string           `bson:"uuid"`
	VerfiedableUUID  string           `bson:"verifiedable_uuid"`
	VerifiedableType VerifiedableType `bson:"verifiedable_type"`
	Code             string           `bson:"code"`
	UsedAt           *int64           `bson:"used_at,omitempty"`
	CreatedAt        int64            `bson:"created_at"`
}

type EmailVerificationCodeModel struct {
	UUID             string           `bson:"uuid" json:"uuid"`
	VerfiedableUUID  string           `bson:"verifiedable_uuid" json:"verifiedable_uuid"`
	VerifiedableType VerifiedableType `bson:"verifiedable_type" json:"verifiedable_type"`
	Code             string           `bson:"code" json:"code"`
	UsedAt           *time.Time       `bson:"used_at,omitempty" json:"used_at,omitempty"`
	CreatedAt        time.Time        `bson:"created_at" json:"created_at"`
}

type EmailVerificationCodeRepositoryContract interface {
	FindEmailVerificationCode(ctx context.Context, id string, withTrashed bool) (*EmailVerificationCode, int, error)
}
