package domain

type VerifiedableType string

const (
	UserVerifiedableType    VerifiedableType = "user"
	EmployeVerifiedableType VerifiedableType = "employee"
)

type EmailVerificationCode struct {
	UUID             string           `bson:"uuid"`
	VerfiedableUUID  string           `bson:"verifiedable_uuid"`
	VerifiedableType VerifiedableType `bson:"verifiedable_type"`
	Code             string           `bson:"code"`
	UsedAt           *int64           `bson:"used_at,omitempty"`
	CreatedAt        int64            `bson:"created_at"`
}
