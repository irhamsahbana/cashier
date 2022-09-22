package domain

type InvoiceOrderGroup struct {
	UUID            string           `bson:"uuid"`
	BranchUUID      string           `bson:"branch_uuid"`
	Payments        []Payment        `bson:"payments"`
	OrderGroups     []OrderGroup     `bson:"order_groups"`
	CreditContracts []CreditContract `bson:"credit_contracts"`
	Customer        Customer         `bson:"customer"`
	CreatedBy       string           `bson:"created_by"`
	Tax             float32          `bson:"tax"`
	Tip             float32          `bson:"tip"`
	Completed       bool             `bson:"completed"`
	CreatedAt       int64            `bson:"created_at"`
}

type CreditContract struct {
	UUID      string `bson:"uuid"`
	Number    string `bson:"number"`
	Note      string `bson:"note"`
	DueBy     int64  `bson:"due_by"`
	CreatedAt int64  `bson:"created_at"`
}

type Payment struct {
	UUID                 string `bson:"uuid"`
	PaymentMethod        `bson:"payment_method"`
	EmployeeShiftInvoice `bson:"employee_shift"`
}

type PaymentMethodInvoice struct {
	PaymentMethodUUID string  `bson:"payment_method_uuid"`
	Group             string  `bson:"group"`
	Name              string  `bson:"name"`
	Fee               float32 `bson:"fee"`
}

type EmployeeShiftInvoice struct {
	EmployeeShiftUUID string `bson:"employee_shift_uuid"`
	UserUUID          string `bson:"user_uuid"`
	Name              string `bson:"name"`
}
