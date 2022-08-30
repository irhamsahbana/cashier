package domain

import "time"

type UserModel struct {
	UUID                   string                       `bson:"uuid" json:"uuid"`
	Branch                 BranchModel                  `bson:"branch" json:"branch"`
	Role                   UserRoleModel                `bson:"role" json:"role"`
	Name                   string                       `bson:"name" json:"name"`
	Email                  string                       `bson:"email" json:"email"`
	EmailVerifiedAt        *int64                       `bson:"email_verified_at,omitempty" json:"email_verified_at,omitempty"`
	EmailVerificationCodes []EmailVerificationCodeModel `bson:"email_verification_code,omitempty" json:"email_verification_code,omitempty"`
	Password               string                       `bson:"password" json:"password"`
	Phone                  *string                      `bson:"phone,omitempty" json:"phone,omitempty"`
	Whatsapp               *string                      `bson:"whatsapp,omitempty" json:"whatsapp,omitempty"`
	ProfileUrl             *string                      `bson:"profile_url,omitempty" json:"profile_url,omitempty"`
	Tokens                 []Token                      `bson:"tokens,omitempty" json:"tokens,omitempty"`
	CreatedAt              time.Time                    `bson:"created_at" json:"created_at"`
	UpdatedAt              *time.Time                   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt              *time.Time                   `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
