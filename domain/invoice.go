package domain

type Invoice struct {
	UUID            string                  `bson:"uuid"`
	BranchUUID      string                  `bson:"branch_uuid"`
	Customer        *Customer               `bson:"customer"`
	Payments        []InvoicePayment        `bson:"payments"`
	OrderGroups     []OrderGroup            `bson:"order_groups"`
	Refunds         []InvoiceRefund         `bson:"refunds"`
	CreditContracts []InvoiceCreditContract `bson:"credit_contracts"`
	TotalAmount     float64                 `bson:"total_amount"`
	TotalTax        float64                 `bson:"total_tax"`
	TotalDiscount   float64                 `bson:"total_discount"`
	TotalChange     float64                 `bson:"total_change"`
	TotalTip        float64                 `bson:"total_tip"`
	Note            *string                 `bson:"note"`
	CompletedAt     *int64                  `bson:"completed_at"`
	CreatedAt       int64                   `bson:"created_at"`
	UpdatedAt       *int64                  `bson:"updated_at"`
}

type InvoiceRefund struct {
	UUID           string               `bson:"uuid"`
	OrderGroupUUID string               `bson:"order_group_uuid"`
	EmployeeShift  InvoiceEmployeeShift `bson:"employee_shift"`
	Total          float64              `bson:"total"`
	CreatedAt      int64                `bson:"created_at"`
	UpdatedAt      *int64               `bson:"updated_at"`
	DeletedAt      *int64               `bson:"deleted_at"`
}

type InvoicePayment struct {
	UUID                 string               `bson:"uuid"`
	OrderGroupUUID       string               `bson:"order_group_uuid"`
	PaymentMethod        InvoicePaymentMethod `bson:"payment_method"`
	EmployeeShiftInvoice InvoiceEmployeeShift `bson:"employee_shift"`
	Total                float64              `bson:"total"`
	Fee                  float64              `bson:"fee"`
	CreatedAt            int64                `bson:"created_at"`
}

type InvoicePaymentMethod struct {
	PaymentMethodUUID string                  `bson:"payment_method_uuid"`
	Group             string                  `bson:"group"`
	Name              string                  `bson:"name"`
	Fee               InvoicePaymentMethodFee `bson:"fee"`
}

type InvoicePaymentMethodFee struct {
	Percent float64 `bson:"percent"`
	Fixed   float64 `bson:"fixed"`
}

type InvoiceEmployeeShift struct {
	EmployeeShiftUUID string `bson:"employee_shift_uuid"`
	UserUUID          string `bson:"user_uuid"`
	Name              string `bson:"name"`
}

type InvoiceCreditContract struct {
	UUID      string `bson:"uuid"`
	Number    string `bson:"number"`
	Note      string `bson:"note"`
	DueBy     int64  `bson:"due_by"`
	CreatedAt int64  `bson:"created_at"`
}

type InvoiceSpace struct {
	ZoneName    string `bson:"zone_name"`
	GroupCode   string `bson:"group_code"`
	SpaceNumber int    `bson:"space_number"`
}

type OrderRefundData struct {
	OrderGroupUUID string `bson:"order_group_uuid" json:"order_group_uuid"`
	OrderUUID      string `bson:"order_uuid" json:"order_uuid"`
	Qty            int64  `bson:"refunded_qty" json:"refunded_qty"`
}
