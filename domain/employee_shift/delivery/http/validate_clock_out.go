package http

import (
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
	"time"
)

func validateClockOutRequest(req *dto.EmployeeShiftClockOutRequest) customtype.Message {
	msg := customtype.Message{}
	if err := validator.IsUUID(req.UUID); err != nil {
		helper.AddMessage("uuid", err.Error(), msg)
	}

	if _, err := time.Parse(time.RFC3339Nano, req.EndTime); err != nil {
		helper.AddMessage("end_time", err.Error(), msg)
	}

	if req.EndCash != nil {
		if *req.EndCash < 0 {
			helper.AddMessage("end_cash", "invalid end cash field", msg)
		}
	}

	return msg
}
