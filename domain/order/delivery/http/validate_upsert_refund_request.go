package http

func ValidateUpsertRefundRequest() {
	// // refunds (optional)
	// for refIndex, refund := range req.Refunds {
	// 	if err := validator.IsUUID(refund.UUID); err != nil {
	// 		msg = helper.AddMessage(fmt.Sprintf("refunds.%d.uuid", refIndex), err.Error(), msg)
	// 	}

	// 	if err := validator.IsUUID(refund.OrderGroupUUID); err != nil {
	// 		msg = helper.AddMessage(fmt.Sprintf("refunds.%d.order_group_uuid", refIndex), err.Error(), msg)
	// 	}

	// 	if err := validator.IsUUID(refund.EmployeeShift.EmployeeShiftUUID); err != nil {
	// 		msg = helper.AddMessage(fmt.Sprintf("refunds.%d.employee_shift.employee_shift_uuid", refIndex), err.Error(), msg)
	// 	}

	// 	if err := validator.IsUUID(refund.EmployeeShift.UserUUID); err != nil {
	// 		msg = helper.AddMessage(fmt.Sprintf("refunds.%d.employee_shift.user_uuid", refIndex), err.Error(), msg)
	// 	}

	// 	if len(refund.EmployeeShift.Name) == 0 {
	// 		msg = helper.AddMessage(fmt.Sprintf("refunds.%d.employee_shift.name", refIndex), "employee shift name must not be empty", msg)
	// 	}
	// }
	// // -- refunds
}
