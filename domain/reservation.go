package domain

type Reservation struct {
	UUID        string     `bson:"uuid"`
	Name        string     `bson:"name"`
	Phone       string     `bson:"phone"`
	Pax         int32      `bson:"pax"`
	Password    string     `bson:"password"`
	OrderGroup  OrderGroup `bson:"order_group"`
	ScheduledAt int64      `bson:"scheduled_at"`
	CreatedAt   int64      `bson:"created_at"`
	UpdatedAt   *int64     `bson:"updated_at"`
}
