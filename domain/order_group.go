package domain

import "time"

type OrderType string

const (
	DineInOrder OrderType = "dine-in"
	DeliveryOrder OrderType = "delivery"
	takeawayOrder OrderType = "takeaway"
)

type OrderGroup struct {
	ID			uint64		`json:"id"`
	UUID		string		`json:"uuid"`
	Type		OrderType	`json:"type"`
	Tax			float32		`json:"tax"`
	CompletedAt	*time.Time	`json:"completed_at"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}



