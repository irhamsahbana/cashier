package http

import (
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"time"
)

func ValidateCashEntryInsertRequest(req *dto.CashEntryInsertRequest) customtype.Message {
	msg := customtype.Message{}

	_, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		msg = helper.AddMessage("created_at", err.Error(), msg)
	}

	return msg
}
