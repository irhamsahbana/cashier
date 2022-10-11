package http

import (
	"fmt"
	"lucy/cashier/dto"
	"lucy/cashier/lib/validator"
	"time"

	"github.com/gin-gonic/gin"
)

func validateUpserOrderRequest(req *dto.OrderGroupUpsertRequest) gin.H {
	// validate if order group is not double type (must be either delivery, queue, space, or none of them(quick order))
	msg := gin.H{}
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

	if typeCount > 1 {
		msg["type"] = "order group type must be either delivery, queue, space, or none of them(quick order)"
	}

	err := validator.IsUUID(req.UUID)
	if err != nil {
		msg["uuid"] = "order group uuid is not valid"
	}

	if len(req.Orders) == 0 {
		msg["orders"] = "orders must not be empty"
	}

	// validate delivery
	if req.Delivery != nil {
		err = validator.IsUUID(req.Delivery.UUID)
		if err != nil {
			msg["delivery.uuid"] = "delivery uuid is not valid"
		}

		if req.Delivery.Number == "" {
			msg["delivery.number"] = "delivery number must not be empty"
		}

		if req.Delivery.Partner == "" {
			msg["delivery.partner"] = "delivery partner must not be empty"
		}

		if req.Delivery.Driver == "" {
			msg["delivery.driver"] = "delivery driver must not be empty"
		}

		if _, err := time.Parse(time.RFC3339Nano, req.Delivery.CreatedAt); err != nil {
			msg["delivery.created_at"] = err.Error()
		}

		if req.Delivery.ScheduledAt != nil {
			if _, err := time.Parse(time.RFC3339Nano, *req.Delivery.ScheduledAt); err != nil {
				msg["delivery.scheduled_at"] = err.Error()
			}
		}

		if req.Delivery.Customer.Name == "" {
			msg["delivery.customer.name"] = "delivery customer name must not be empty"
		}
	}

	// validate queue
	if req.Queue != nil {
		err = validator.IsUUID(req.Queue.UUID)
		if err != nil {
			msg["queue.uuid"] = "queue uuid is not valid"
		}

		if req.Queue.Number == "" {
			msg["queue.number"] = "queue number must not be empty"
		}

		if req.Queue.Customer.Name == "" {
			msg["queue.customer.name"] = "queue customer name must not be empty"
		}

		if req.Queue.ScheduledAt != nil {
			_, err := time.Parse(time.RFC3339, *req.Queue.ScheduledAt)
			if err != nil {
				msg["queue.scheduled_at"] = err.Error()
			}
		}
	}

	// validate order
	for orderIndex, order := range req.Orders {
		err = validator.IsUUID(order.UUID)
		if err != nil {
			msg[fmt.Sprintf("orders.%d.uuid", orderIndex)] = "order uuid is not valid"
		}

		err = validator.IsUUID(order.Item.UUID)
		if err != nil {
			msg[fmt.Sprintf("orders.%d.item.uuid", orderIndex)] = "order item uuid is not valid"
		}

		if order.Item.Quantity < 0 {
			msg[fmt.Sprintf("orders.%d.item.quantity", orderIndex)] = "order item quantity must not be negative"
		}

		_, err := time.Parse(time.RFC3339Nano, order.CreatedAt)
		if err != nil {
			msg[fmt.Sprintf("orders.%d.created_at", orderIndex)] = err.Error()
		}

		// order modifiers
		for modIndex, modifier := range order.Modifiers {
			err = validator.IsUUID(modifier.UUID)
			if err != nil {
				msg[fmt.Sprintf("orders.%d.modifiers.%d.uuid", orderIndex, modIndex)] = "order modifier uuid is not valid"
			}

			if modifier.Quantity < 0 {
				msg[fmt.Sprintf("orders.%d.modifiers.%d.quantity", orderIndex, modIndex)] = "order modifier quantity must not be negative"
			}
		}

		// order waiter
		if order.Waiter != nil {
			err = validator.IsUUID(order.Waiter.UUID)
			if err != nil {
				msg[fmt.Sprintf("orders.%d.waiter.uuid", orderIndex)] = "order waiter uuid is not valid"
			}

			err = validator.IsUUID(order.Waiter.BranchUUID)
			if err != nil {
				msg[fmt.Sprintf("orders.%d.waiter.branch_uuid", orderIndex)] = "order waiter branch uuid is not valid"
			}
		}
	}
	return msg
}
