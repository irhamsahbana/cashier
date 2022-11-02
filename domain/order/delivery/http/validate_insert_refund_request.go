package http

import (
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
	"time"
)

func ValidateInsertRefundRequest(req *dto.RefundInsertRequest) customtype.Message {
	msg := customtype.Message{}

	if err := validator.IsUUID(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", err.Error(), msg)
	}

	if err := validator.IsUUID(req.OrderGroupUUID); err != nil {
		msg = helper.AddMessage("order_group_uuid", err.Error(), msg)
	}

	if req.Total <= 0 {
		msg = helper.AddMessage("total", "total must be greater than 0", msg)
	}

	if err := validator.IsUUID(req.EmployeeShift.EmployeeShiftUUID); err != nil {
		msg = helper.AddMessage("employee_shift.employee_shift_uuid", err.Error(), msg)
	}

	if err := validator.IsUUID(req.EmployeeShift.UserUUID); err != nil {
		msg = helper.AddMessage("employee_shift.user_uuid", err.Error(), msg)
	}

	if req.EmployeeShift.Name == "" {
		msg = helper.AddMessage("employee_shift.name", "employee shift name must not be empty", msg)
	}

	if _, err := time.Parse(time.RFC3339Nano, req.CreatedAt); err != nil {
		msg = helper.AddMessage("created_at", err.Error(), msg)
	}

	return msg
}
