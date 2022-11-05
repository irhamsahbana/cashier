package http

import (
	"fmt"
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	"lucy/cashier/lib/validator"
	"time"
)

func ValidateInsertInvoiceRequest(req *dto.InvoiceInsertRequest) customtype.Message {
	msg := customtype.Message{}

	if err := validator.Uuid(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", err.Error(), msg)
	}

	// customer
	if req.Customer != nil {
		if req.Customer.Name == "" {
			msg = helper.AddMessage("customer.name", "customer name must not be empty", msg)
		}

		if req.Customer.Phone != nil {
			if len(*req.Customer.Phone) > 15 {
				msg = helper.AddMessage("customer.phone", "customer phone must be less than 15 characters", msg)
			}
		}

		if req.Customer.Address != nil {
			if len(*req.Customer.Address) > 255 {
				msg = helper.AddMessage("customer.address", "customer address must be less than 255 characters", msg)
			}
		}
	}
	// -- customer

	// Payments
	for payIndex, payment := range req.Payments {
		if err := validator.Uuid(payment.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.uuid", payIndex), err.Error(), msg)
		}

		if err := validator.Uuid(payment.OrderGroupUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.order_group_uuid", payIndex), err.Error(), msg)
		}

		// payment method
		if err := validator.Uuid(payment.PaymentMethod.PaymentMethodUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.payment_method_uuid", payIndex), err.Error(), msg)
		}

		groups := []string{"cash", "edc", "e-wallet", "qris", "delivery"}
		if !validator.In(groups, payment.PaymentMethod.Group) {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.group", payIndex), "payment method group must be either cash, edc, e-wallet, qris, or delivery", msg)
		}

		if len(payment.PaymentMethod.Name) == 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.name", payIndex), "payment method name must not be empty", msg)
		}

		// payment method fee
		var feeCount int

		if payment.PaymentMethod.Fee.Fixed != 0 {
			feeCount++
		}

		if payment.PaymentMethod.Fee.Percent != 0 {
			feeCount++
		}

		if feeCount > 1 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.fee", payIndex), "payment method fee must be either fixed or percent", msg)
		}

		if payment.PaymentMethod.Fee.Fixed < 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.fee.fixed", payIndex), "payment method fee fixed must not be less than 0", msg)
		}

		if payment.PaymentMethod.Fee.Percent < 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.payment_method.fee.percent", payIndex), "payment method fee percent must not be less than 0", msg)
		}
		// -- payment method fee

		if _, err := time.Parse(time.RFC3339, payment.CreatedAt); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.created_at", payIndex), err.Error(), msg)
		}
		// --payment method

		// employee shift
		if err := validator.Uuid(payment.EmployeeShift.EmployeeShiftUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.employee_shift.employee_shift_uuid", payIndex), err.Error(), msg)
		}

		if err := validator.Uuid(payment.EmployeeShift.UserUUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.employee_shift.user_uuid", payIndex), err.Error(), msg)
		}

		if len(payment.EmployeeShift.Name) == 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.employee_shift.name", payIndex), "employee shift name must not be empty", msg)
		}
		// -- employee shift

		if payment.Total < 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.total", payIndex), "payment total must not be negative", msg)
		}

		if payment.Fee < 0 {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.fee", payIndex), "payment fee must not be negative", msg)
		}

		_, err := time.Parse(time.RFC3339Nano, payment.CreatedAt)
		if err != nil {
			msg = helper.AddMessage(fmt.Sprintf("payments.%d.created_at", payIndex), err.Error(), msg)
		}
	}
	// -- Payments

	// order groups (must not be empty)
	if len(req.OrderGroups) == 0 {
		msg = helper.AddMessage("order_groups", "order groups must not be empty", msg)
	}

	for ogIndex, og := range req.OrderGroups {
		var typeCount int

		if og.Delivery != nil {
			typeCount++
		}
		if og.Queue != nil {
			typeCount++
		}
		if og.SpaceUUID != nil {
			typeCount++
		}

		if typeCount > 1 || typeCount == 0 {
			msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.type", ogIndex), "order group type must be either delivery, queue, or space", msg)
		}

		if err := validator.Uuid(og.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.uuid", ogIndex), err.Error(), msg)
		}

		if len(og.Orders) == 0 {
			msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders", ogIndex), "orders must not be empty", msg)
		}

		// validate delivery
		if og.Delivery != nil {
			if err := validator.Uuid(og.Delivery.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.uuid", ogIndex), err.Error(), msg)
			}

			if og.Delivery.Partner == "" {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.partner", ogIndex), "delivery partner must not be empty", msg)
			}

			if og.Delivery.Driver == "" {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.driver", ogIndex), "delivery driver must not be empty", msg)
			}

			// if _, err := time.Parse(time.RFC3339Nano, og.Delivery.CreatedAt); err != nil {
			// 	msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.created_at", ogIndex), err.Error(), msg)
			// }

			if og.Delivery.ScheduledAt != nil {
				if _, err := time.Parse(time.RFC3339Nano, *og.Delivery.ScheduledAt); err != nil {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.scheduled_at", ogIndex), err.Error(), msg)
				}
			}

			if og.Delivery.Customer.Name == "" {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.delivery.customer.name", ogIndex), "delivery customer name must not be empty", msg)
			}
		}

		// validate queue
		if og.Queue != nil {
			if err := validator.Uuid(og.Queue.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.queue.uuid", ogIndex), err.Error(), msg)
			}

			if og.Queue.Customer.Name == "" {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.queue.customer.name", ogIndex), "queue customer name must not be empty", msg)
			}

			if og.Queue.ScheduledAt != nil {
				if _, err := time.Parse(time.RFC3339, *og.Queue.ScheduledAt); err != nil {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.queue.scheduled_at", ogIndex), err.Error(), msg)
				}
			}
		}

		// validate order
		for orderIndex, order := range og.Orders {
			if err := validator.Uuid(order.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.uuid", ogIndex, orderIndex), err.Error(), msg)
			}

			if err := validator.Uuid(order.Item.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.item.uuid", ogIndex, orderIndex), err.Error(), msg)
			}

			if order.Item.Quantity < 0 {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.item.quantity", ogIndex, orderIndex), "order item quantity must not be negative", msg)
			}

			if _, err := time.Parse(time.RFC3339Nano, order.CreatedAt); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.created_at", ogIndex, orderIndex), err.Error(), msg)
			}

			// order modifiers
			for modIndex, modifier := range order.Modifiers {
				if err := validator.Uuid(modifier.UUID); err != nil {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.modifiers.%d.uuid", ogIndex, orderIndex, modIndex), err.Error(), msg)
				}

				if modifier.Quantity < 1 {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.modifiers.%d.quantity", ogIndex, orderIndex, modIndex), "order modifier quantity must more than 0", msg)
				}
			}

			// order waiter
			if order.Waiter != nil {
				if err := validator.Uuid(order.Waiter.UUID); err != nil {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.waiter.uuid", ogIndex, orderIndex), err.Error(), msg)
				}

				if err := validator.Uuid(order.Waiter.BranchUUID); err != nil {
					msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.orders.%d.waiter.branch_uuid", ogIndex, orderIndex), err.Error(), msg)
				}
			}
		}

		// validate taxes
		for taxIndex, tax := range og.Taxes {
			if err := validator.Uuid(tax.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.taxes.%d.uuid", ogIndex, taxIndex), err.Error(), msg)
			}

			if tax.Value < 0 {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.taxes.%d.value", ogIndex, taxIndex), "tax value must not be negative", msg)
			}

			if tax.Value > 100 {
				msg = helper.AddMessage(fmt.Sprintf("order_groups.%d.taxes.%d.value", ogIndex, taxIndex), "tax value must not be more than 100", msg)
			}
		}
	}
	// -- order groups

	// credit contracts
	for ccIndex, cc := range req.CreditContracts {
		if err := validator.Uuid(cc.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("credit_contracts.%d.uuid", ccIndex), err.Error(), msg)
		}

		if len(cc.Note) > 255 {
			msg = helper.AddMessage(fmt.Sprintf("credit_contracts.%d.note", ccIndex), "credit contract note must not be more than 255 characters", msg)
		}

		if _, err := time.Parse(time.RFC3339Nano, cc.DueBy); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("credit_contracts.%d.due_by", ccIndex), err.Error(), msg)
		}

		if _, err := time.Parse(time.RFC3339Nano, cc.CreatedAt); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("credit_contracts.%d.created_at", ccIndex), err.Error(), msg)
		}
	}

	return msg
}
