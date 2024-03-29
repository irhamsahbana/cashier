package http

import (
	"fmt"
	"lucy/cashier/dto"
	customtype "lucy/cashier/lib/custom_type"
	"lucy/cashier/lib/helper"
	valid "lucy/cashier/lib/validator"
	"time"
)

func validateUpserOrderRequest(req *dto.OrderGroupUpsertRequest) customtype.Message {
	// validate if order group is not double type (must be either delivery, queue, space, or none of them(quick order))
	msg := customtype.Message{}
	var typeCount int

	if req.Delivery != nil {
		typeCount++
	}
	if req.Queue != nil {
		typeCount++
	}
	if req.SpaceUUID != nil {
		typeCount++
	}

	if typeCount > 1 || typeCount == 0 {
		msg = helper.AddMessage("type", "order group type must be either delivery, queue, or space", msg)
	}

	if err := valid.Uuid(req.UUID); err != nil {
		msg = helper.AddMessage("uuid", err.Error(), msg)
	}

	if len(req.Orders) == 0 {
		msg = helper.AddMessage("orders", "orders must not be empty", msg)
	}

	// validate delivery
	if req.Delivery != nil {
		if err := valid.Uuid(req.Delivery.UUID); err != nil {
			msg = helper.AddMessage("delivery.uuid", err.Error(), msg)
		}

		if req.Delivery.Partner == "" {
			msg = helper.AddMessage("delivery.partner", "delivery partner must not be empty", msg)
		}

		if req.Delivery.Driver == "" {
			msg = helper.AddMessage("delivery.driver", "delivery driver must not be empty", msg)
		}

		// if _, err := time.Parse(time.RFC3339Nano, req.Delivery.CreatedAt); err != nil {
		// 	msg = helper.AddMessage("delivery.created_at", err.Error(), msg)
		// }

		if req.Delivery.ScheduledAt != nil {
			if _, err := time.Parse(time.RFC3339Nano, *req.Delivery.ScheduledAt); err != nil {
				msg = helper.AddMessage("delivery.scheduled_at", err.Error(), msg)
			}
		}

		if req.Delivery.Customer.Name == "" {
			msg = helper.AddMessage("delivery.customer.name", "delivery customer name must not be empty", msg)
		}
	}

	// validate queue
	if req.Queue != nil {
		if err := valid.Uuid(req.Queue.UUID); err != nil {
			msg = helper.AddMessage("queue.uuid", err.Error(), msg)
		}

		if req.Queue.Customer.Name == "" {
			msg = helper.AddMessage("queue.customer.name", "queue customer name must not be empty", msg)
		}

		if req.Queue.ScheduledAt != nil {
			if _, err := time.Parse(time.RFC3339, *req.Queue.ScheduledAt); err != nil {
				msg = helper.AddMessage("queue.scheduled_at", err.Error(), msg)
			}
		}
	}

	// validate order
	for orderIndex, order := range req.Orders {
		if err := valid.Uuid(order.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("orders.%d.uuid", orderIndex), err.Error(), msg)
		}

		if err := valid.Uuid(order.Item.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("orders.%d.item.uuid", orderIndex), err.Error(), msg)
		}

		if order.Item.Quantity < 0 {
			msg = helper.AddMessage(fmt.Sprintf("orders.%d.item.quantity", orderIndex), "order item quantity must not be negative", msg)
		}

		if _, err := time.Parse(time.RFC3339Nano, order.CreatedAt); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("orders.%d.created_at", orderIndex), err.Error(), msg)
		}

		// order modifiers
		for modIndex, modifier := range order.Modifiers {
			if err := valid.Uuid(modifier.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("orders.%d.modifiers.%d.uuid", orderIndex, modIndex), err.Error(), msg)
			}

			if modifier.Quantity < 1 {
				msg = helper.AddMessage(fmt.Sprintf("orders.%d.modifiers.%d.quantity", orderIndex, modIndex), "order modifier quantity must more than 0", msg)
			}
		}

		// order waiter
		if order.Waiter != nil {
			if err := valid.Uuid(order.Waiter.UUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("orders.%d.waiter.uuid", orderIndex), err.Error(), msg)
			}

			if err := valid.Uuid(order.Waiter.BranchUUID); err != nil {
				msg = helper.AddMessage(fmt.Sprintf("orders.%d.waiter.branch_uuid", orderIndex), err.Error(), msg)
			}
		}
	}

	// validate taxes
	for taxIndex, tax := range req.Taxes {
		if err := valid.Uuid(tax.UUID); err != nil {
			msg = helper.AddMessage(fmt.Sprintf("taxes.%d.uuid", taxIndex), err.Error(), msg)
		}

		if tax.Value < 0 {
			msg = helper.AddMessage(fmt.Sprintf("taxes.%d.value", taxIndex), "tax value must not be negative", msg)
		}

		if tax.Value > 100 {
			msg = helper.AddMessage(fmt.Sprintf("taxes.%d.value", taxIndex), "tax value must not be more than 100", msg)
		}
	}

	return msg
}
