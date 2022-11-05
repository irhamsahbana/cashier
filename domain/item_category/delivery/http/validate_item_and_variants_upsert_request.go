package http

import (
	"fmt"
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
)

func ValidateItemVariantsUpsertRequest(req *dto.ItemAndVariantsUpsertRequest) customtype.Message {
	msg := customtype.Message{}

	if err := validator.Uuid(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", err.Error(), msg)
	}

	if len(req.Name) < 3 {
		msg = helper.AddMessage("name", "Name must be at least 3 characters", msg)
	}

	if req.Price < 0 {
		msg = helper.AddMessage("price", "Price must be greater than 0", msg)
	}

	// variants
	for vi, v := range req.Variants {
		if err := validator.Uuid(v.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("variants.%d.uuid", vi), err.Error(), msg)
		}

		if v.Price < 0 {
			msg = helper.AddMessage(fmt.Sprintf("variants.%d.price", vi), "Price must be greater than 0", msg)
		}

		if len(v.Label) < 3 {
			msg = helper.AddMessage(fmt.Sprintf("variants.%d.label", vi), "Label must be at least 3 characters", msg)
		}
	}

	return msg
}
