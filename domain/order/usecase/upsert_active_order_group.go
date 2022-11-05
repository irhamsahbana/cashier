package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"time"
)

func (u *orderUsecase) UpsertActiveOrder(c context.Context, branchId string, req *dto.OrderGroupUpsertRequest) (*dto.OrderGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	var data domain.OrderGroup
	data.BranchUUID = branchId
	DTOtoDomain_UpsertActiveOrder(&data, req)

	result, code, err := u.orderRepo.UpsertActiveOrder(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.OrderGroupResponse
	DomainToDTO_UpsertActiveOrder(&resp, result)

	return &resp, code, nil
}

// DTO to Domain
func DTOtoDomain_UpsertActiveOrder(data *domain.OrderGroup, req *dto.OrderGroupUpsertRequest) {
	data.UUID = req.UUID
	data.CreatedBy = req.CreatedBy

	// delivery
	if req.Delivery != nil {
		var customer domain.Customer
		customer.Name = req.Delivery.Customer.Name
		customer.Phone = req.Delivery.Customer.Phone
		customer.Address = req.Delivery.Customer.Address

		var schedludedAtTime *int64
		if req.Delivery.ScheduledAt != nil {
			scheduledAt, _ := time.Parse(time.RFC3339Nano, *req.Delivery.ScheduledAt)
			scheduledAtUnixMicro := scheduledAt.UnixMicro()
			schedludedAtTime = &scheduledAtUnixMicro
		}

		var delivery domain.Delivery
		data.Delivery = &delivery
		data.Delivery.UUID = req.Delivery.UUID
		data.Delivery.Number = req.Delivery.Number
		data.Delivery.Partner = req.Delivery.Partner
		data.Delivery.Driver = req.Delivery.Driver
		data.Delivery.Customer = customer
		data.Delivery.ScheduledAt = schedludedAtTime
	}

	// queue
	if req.Queue != nil {
		var customer domain.Customer
		customer.Name = req.Queue.Customer.Name
		customer.Phone = req.Queue.Customer.Phone
		customer.Address = req.Queue.Customer.Address

		var ScheduledAtMicro *int64
		if req.Queue.ScheduledAt != nil {
			scheduledAt, _ := time.Parse(time.RFC3339Nano, *req.Queue.ScheduledAt)
			scheduled := scheduledAt.UnixMicro()
			ScheduledAtMicro = &scheduled
		}
		var queue domain.Queue
		data.Queue = &queue
		data.Queue.UUID = req.Queue.UUID
		data.Queue.Number = req.Queue.Number
		data.Queue.Customer = customer
		data.Queue.ScheduledAt = ScheduledAtMicro

	}

	// space
	data.SpaceUUID = req.SpaceUUID

	// orders
	orders := []domain.Order{}
	for _, order := range req.Orders {
		// waiters
		var waiter domain.WaiterOrder
		if order.Waiter != nil {
			waiter.UUID = order.Waiter.UUID
			waiter.BranchUUID = order.Waiter.BranchUUID
			waiter.Name = order.Waiter.Name
		}

		// order modifiers
		modifiers := []domain.ModifierOrder{}
		for _, modifier := range order.Modifiers {
			var mod domain.ModifierOrder

			mod.UUID = modifier.UUID
			mod.Name = modifier.Name
			mod.Price = modifier.Price
			mod.Quantity = modifier.Quantity

			modifiers = append(modifiers, mod)
		}

		// item
		var item domain.ItemOrder
		item.UUID = order.Item.UUID
		item.Name = order.Item.Name
		item.Label = order.Item.Label
		item.Price = order.Item.Price
		item.Quantity = order.Item.Quantity

		// discount
		discounts := []domain.DiscountOrder{}
		for _, discount := range order.Discounts {
			var discountOrder domain.DiscountOrder
			discountOrder.UUID = discount.UUID
			discountOrder.Name = discount.Name
			discountOrder.Fixed = discount.Fixed
			discountOrder.Percent = discount.Percent

			discounts = append(discounts, discountOrder)
		}

		createdAt, _ := time.Parse(time.RFC3339Nano, order.CreatedAt)
		orders = append(orders, domain.Order{
			UUID:      order.UUID,
			Item:      item,
			Modifiers: modifiers,
			Discounts: discounts,
			Waiter:    &waiter,
			Note:      order.Note,
			CreatedAt: createdAt.UnixMicro(),
		})
	}
	data.Orders = orders

	// taxes
	taxes := []domain.TaxOrderGroup{}
	for _, tax := range req.Taxes {
		var taxOrder domain.TaxOrderGroup
		taxOrder.UUID = tax.UUID
		taxOrder.Name = tax.Name
		taxOrder.Value = tax.Value

		taxes = append(taxes, taxOrder)
	}
	data.Taxes = taxes

	createdAt, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)
	data.CreatedAt = createdAt.UnixMicro()
}

// Domain to DTO
func DomainToDTO_UpsertActiveOrder(resp *dto.OrderGroupResponse, data *domain.OrderGroup) {
	resp.UUID = data.UUID
	resp.BranchUUID = data.BranchUUID
	resp.CreatedBy = data.CreatedBy

	resp.CreatedAt = time.UnixMicro(data.CreatedAt).UTC()
	resp.CancelReason = data.CancelReason
	if data.UpdatedAt != nil {
		dataUpdatedAt := time.UnixMicro(*data.UpdatedAt).UTC()
		resp.UpdatedAt = &dataUpdatedAt
	}
	if data.DeletedAt != nil {
		dataDeletedAt := time.UnixMicro(*data.DeletedAt).UTC()
		resp.DeletedAt = &dataDeletedAt
	}

	// delivery
	if data.Delivery != nil {
		customer := dto.Customer{
			Name:    data.Delivery.Customer.Name,
			Phone:   data.Delivery.Customer.Phone,
			Address: data.Delivery.Customer.Address,
		}

		var deliveryScheduledAt *time.Time
		if data.Delivery.ScheduledAt != nil {
			scheduledAt := time.UnixMicro(*data.Delivery.ScheduledAt).UTC()
			deliveryScheduledAt = &scheduledAt
		}

		resp.Delivery = &dto.DeliveryResponse{
			UUID:        data.Delivery.UUID,
			Number:      data.Delivery.Number,
			Partner:     data.Delivery.Partner,
			Driver:      data.Delivery.Driver,
			Customer:    customer,
			ScheduledAt: deliveryScheduledAt,
		}
	}

	// queue
	if data.Queue != nil {
		customer := dto.Customer{
			Name:    data.Queue.Customer.Name,
			Phone:   data.Queue.Customer.Phone,
			Address: data.Queue.Customer.Address,
		}

		var scheduledAt *string
		if data.Queue.ScheduledAt != nil {
			scheduledAtTime := time.UnixMicro(*data.Queue.ScheduledAt)
			promisedAtStr := scheduledAtTime.Format(time.RFC3339Nano)
			scheduledAt = &promisedAtStr
		}
		// createdAt := time.UnixMicro(data.Queue.CreatedAt)
		resp.Queue = &dto.Queue{
			UUID:        data.Queue.UUID,
			Number:      data.Queue.Number,
			Customer:    customer,
			ScheduledAt: scheduledAt,
			// CreatedAt:   createdAt.Format(time.RFC3339Nano),
		}
	}

	// space
	resp.SpaceUUID = data.SpaceUUID

	// orders
	orders := []dto.OrderResponse{}
	for _, order := range data.Orders {
		var waiter *dto.WaiterOrder
		if order.Waiter != nil {
			waiter = &dto.WaiterOrder{
				UUID:       order.Waiter.UUID,
				BranchUUID: order.Waiter.BranchUUID,
				Name:       order.Waiter.Name,
			}
		}

		// order modifiers
		modifiers := []dto.ModifierOrder{}
		for _, modifier := range order.Modifiers {
			modifiers = append(modifiers, dto.ModifierOrder{
				UUID:     modifier.UUID,
				Name:     modifier.Name,
				Quantity: modifier.Quantity,
				Price:    modifier.Price,
			})
		}

		// order discounts
		discounts := []dto.DiscountOrder{}
		for _, discount := range order.Discounts {
			discounts = append(discounts, dto.DiscountOrder{
				UUID:    discount.UUID,
				Name:    discount.Name,
				Fixed:   discount.Fixed,
				Percent: discount.Percent,
			})
		}

		// item
		var item dto.ItemOrder
		item.UUID = order.Item.UUID
		item.Name = order.Item.Name
		item.Label = order.Item.Label
		item.Price = order.Item.Price
		item.Quantity = order.Item.Quantity

		// append order
		var orderUpdatedAt *time.Time
		var orderDeletedAt *time.Time
		if order.UpdatedAt != nil {
			updatedAt := time.UnixMicro(*order.UpdatedAt).UTC()
			orderUpdatedAt = &updatedAt
		}
		if order.DeletedAt != nil {
			deletedAt := time.UnixMicro(*order.DeletedAt).UTC()
			orderDeletedAt = &deletedAt
		}
		orders = append(orders, dto.OrderResponse{
			UUID:        order.UUID,
			Item:        item,
			Modifiers:   modifiers,
			Discounts:   discounts,
			Waiter:      waiter,
			Note:        order.Note,
			RefundedQty: order.RefundedQty,
			CreatedAt:   time.UnixMicro(order.CreatedAt).Format(time.RFC3339Nano),
			UpdatedAt:   orderUpdatedAt,
			DeletedAt:   orderDeletedAt,
		})
	}
	resp.Orders = orders

	// taxes
	taxes := []dto.TaxOrderGroup{}
	for _, tax := range data.Taxes {
		taxes = append(taxes, dto.TaxOrderGroup{
			UUID:  tax.UUID,
			Name:  tax.Name,
			Value: tax.Value,
		})
	}
	resp.Taxes = taxes

	resp.CreatedAt = time.UnixMicro(data.CreatedAt).UTC()
	if data.UpdatedAt != nil {
		dataUpdatedAt := time.UnixMicro(*data.UpdatedAt).UTC()
		resp.UpdatedAt = &dataUpdatedAt
	}
	if data.DeletedAt != nil {
		dataDeletedAt := time.UnixMicro(*data.DeletedAt).UTC()
		resp.DeletedAt = &dataDeletedAt
	}
}
