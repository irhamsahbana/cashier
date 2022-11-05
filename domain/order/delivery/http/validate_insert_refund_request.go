package http

import (
	"fmt"
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	valid "lucy/cashier/lib/validator"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ValidateInsertRefundRequest(req *dto.RefundInsertRequest) customtype.Message {
	msg := customtype.Message{}

	if err := valid.Uuid(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", err.Error(), msg)
	}

	if err := valid.Uuid(req.InvoiceUUID); err != nil {
		msg = helper.AddMessage("invoice_uuid", err.Error(), msg)
	}

	if err := valid.Exists("invoices", bson.M{"uuid": req.InvoiceUUID}); err != nil {
		msg = helper.AddMessage("invoice_uuid", fmt.Sprintf("invoice uuid not found => %s", err.Error()), msg)
	}

	// total
	if req.Total <= 0 {
		msg = helper.AddMessage("total", "total must be greater than 0", msg)
	}
	// -- total

	// employee
	if err := valid.Uuid(req.EmployeeShift.EmployeeShiftUUID); err != nil {
		msg = helper.AddMessage("employee_shift.employee_shift_uuid", err.Error(), msg)
	}

	if err := valid.Uuid(req.EmployeeShift.UserUUID); err != nil {
		msg = helper.AddMessage("employee_shift.user_uuid", err.Error(), msg)
	}

	if req.EmployeeShift.Name == "" {
		msg = helper.AddMessage("employee_shift.name", "employee shift name must not be empty", msg)
	}

	if _, err := time.Parse(time.RFC3339Nano, req.CreatedAt); err != nil {
		msg = helper.AddMessage("created_at", err.Error(), msg)
	}
	// -- employee

	// order refunds
	for i, orderRefund := range req.OrderRefunds {
		if err := valid.Uuid(orderRefund.OrderUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("order_refunds.%d.order_uuid", i), err.Error(), msg)
		}

		// order group uuid
		if err := valid.Uuid(orderRefund.OrderGroupUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("order_refunds.%d.order_group_uuid", i), err.Error(), msg)
		}

		filter := bson.M{"order_groups.uuid": orderRefund.OrderGroupUUID}
		if err := valid.Exists("invoices", filter); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("order_refunds.%d.order_group_uuid", i), fmt.Sprintf("order group uuid not found => %s", err.Error()), msg)
		}
		// -- order group uuid

		filter = bson.M{"order_groups.orders.uuid": orderRefund.OrderUUID}
		if err := valid.Exists("invoices", filter); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("order_refunds.%d.order_uuid", i), fmt.Sprintf("order uuid not found => %s", err.Error()), msg)
		}

		if orderRefund.Qty <= 0 {
			msg = helper.AddMessage(fmt.Sprintf("order_refunds.%d.quantity", i), "quantity must be greater than 0", msg)
		}
	}
	// -- order refunds

	return msg
}
