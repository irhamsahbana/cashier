package dto

import "time"

type OrderGroupUpsertRequest struct {
	UUID      string    `bson:"uuid"`
	SpaceUUID *string   `bson:"space_uuid"`
	Delivery  *Delivery `bson:"delivery"`
	Queue     *Queue    `bson:"queue"`
	Orders    []Order   `bson:"orders"`
	CreatedBy string    `bson:"created_by"`
	Tax       float32   `bson:"tax"`
	Tip       float32   `bson:"tip"`
	Pending   *bool     `bson:"pending"`
	Completed bool      `bson:"completed"`
	CreatedAt string    `bson:"created_at"`
}

type OrderGroup struct {
	UUID       string    `bson:"uuid"`
	BranchUUID string    `bson:"branch_uuid"`
	SpaceUUID  *string   `bson:"space_uuid"`
	Delivery   *Delivery `bson:"delivery"`
	Queue      *Queue    `bson:"queue"`
	Orders     []Order   `bson:"orders"`
	CreatedBy  string    `bson:"created_by"`
	Tax        float32   `bson:"tax"`
	Tip        float32   `bson:"tip"`
	Pending    *bool     `bson:"pending"`
	Completed  bool      `bson:"completed"`
	CreatedAt  time.Time `bson:"created_at"`
}

type Order struct {
	UUID        string          `bson:"uuid"`
	Item        ItemOrder       `bson:"item"`
	Modifiers   []ModifierOrder `bson:"modifiers"`
	Waiter      *WaiterOrder    `bson:"waiter"`
	RefundedQty int32           `bson:"refunded_qty"`
	CreatedAt   time.Time       `bson:"created_at"`
}

type ItemOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Label    string  `bson:"label"`
	Price    float32 `bson:"price"`
	Quantity int     `bson:"quantity"`
}

type ModifierOrder struct {
	UUID     string  `bson:"uuid"`
	Name     string  `bson:"name"`
	Quantity int     `bson:"quantity"`
	Price    float32 `bson:"price"`
}

type WaiterOrder struct {
	UUID       string `bson:"uuid"`
	BranchUUID string `bson:"branch_uuid"`
	Name       string `bson:"name"`
}

type Delivery struct {
	UUID      string    `bson:"uuid"`
	Number    string    `bson:"number"`
	Partner   string    `bson:"partner"`
	Driver    string    `bson:"driver"`
	Customer  Customer  `bson:"customer"`
	CreatedAt time.Time `bson:"created_at"`
}

type Queue struct {
	UUID       string     `bson:"uuid"`
	Customer   Customer   `bson:"customer"`
	PromisedAt *time.Time `bson:"promised_at"`
	CreatedAt  time.Time  `bson:"created_at"`
}

type Customer struct {
	Name    string  `bson:"name"`
	Phone   *string `bson:"phone"`
	Address *string `bson:"address"`
}