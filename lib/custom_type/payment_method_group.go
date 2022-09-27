package customtype

type PaymentMethodGroup string

const (
	PaymentMethodGroup_CASH     PaymentMethodGroup = "cash"
	PaymentMethodGroup_DELIVERY PaymentMethodGroup = "delivery"
	PaymentMethodGroup_EDC      PaymentMethodGroup = "edc"
	PaymentMethodGroup_EWALLET  PaymentMethodGroup = "e-wallet"
	PaymentMethodGroup_QRIS     PaymentMethodGroup = "qris"
)
