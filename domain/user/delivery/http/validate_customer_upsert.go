package http

import (
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
	"net/mail"
	"time"
)

func validateUpsertCustomerRequest(req *dto.CustomerUpserRequest) customtype.Message {
	msg := customtype.Message{}

	// uuid
	if err := validator.IsUUID(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", "customer uuid is not valid", msg)
	}
	// -- uuid

	//name
	if req.Name == "" {
		msg = helper.AddMessage("name", "name must not be empty", msg)
	}

	if len(req.Name) > 100 {
		msg = helper.AddMessage("name", "name must not be more than 100 characters", msg)
	}

	if len(req.Name) < 3 {
		msg = helper.AddMessage("name", "name must not be less than 3 characters", msg)
	}
	// -- name

	// dob
	if req.Dob == "" {
		msg = helper.AddMessage("date_of_birth", "date of brith must not be empty", msg)
	}

	if _, err := time.Parse(time.RFC3339Nano, req.Dob); err != nil {
		msg = helper.AddMessage("date_of_birth", "date of birth must be in RFC3339Nano format", msg)
	}
	// -- dob

	// email
	if _, err := mail.ParseAddress(req.Email); err != nil {
		msg = helper.AddMessage("email", "email is not valid", msg)
	}
	// -- email

	// created_at
	if _, err := time.Parse(time.RFC3339Nano, req.CreatedAt); err != nil {
		msg = helper.AddMessage("created_at", "created at must be in RFC3339Nano format", msg)
	}
	// -- created_at

	return msg
}
