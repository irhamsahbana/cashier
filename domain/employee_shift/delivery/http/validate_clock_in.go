package http

import (
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
	"time"
)

func ValidateClockInRequest(req *dto.EmployeeShiftClockInRequest) customtype.Message {
	msg := customtype.Message{}

	if err := validator.IsUUID(req.UUID); err != nil {
		helper.AddMessage("uuid", err.Error(), msg)
	}

	if err := validator.IsUUID(req.UserUUID); err != nil {
		helper.AddMessage("user_uuid", err.Error(), msg)
	}

	if req.SupportingUUID != nil {
		if err := validator.IsUUID(*req.SupportingUUID); err != nil {
			helper.AddMessage("supporting_uuid", err.Error(), msg)
		}
	}

	if req.StartCash != nil && req.SupportingUUID != nil {
		helper.AddMessage("start_cash", "start_cash and supporting_uuid field cannot be set at the same time", msg)
		helper.AddMessage("supporting_uuid", "start_cash and supporting_uuid field cannot be set at the same time", msg)
	}

	if req.SupportingUUID == nil && req.StartCash == nil {
		helper.AddMessage("start_cash", "start_cash field is required if supporting_uuid is null", msg)
	}

	_, err := time.Parse(time.RFC3339Nano, req.StartTime)
	if err != nil {
		helper.AddMessage("start_time", err.Error(), msg)
	}

	return msg
}
