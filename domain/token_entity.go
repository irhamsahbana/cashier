package domain

type TokenableType string

const (
	UserTokenableType     TokenableType = "user"
	EmployeeTokenableType TokenableType = "employee"
)

type Token struct {
	UUID          string        `bson:"uuid"`
	TokenableUUID string        `bson:"tokenable_uuid"`
	TokenableType TokenableType `bson:"tokenable_type"`
	AccessToken   string        `bson:"access_token"`
	RefreshToken  string        `bson:"refresh_token"`
}

type TokenUsecaseContract interface {
}

type TokenRepositoryContract interface {
}
