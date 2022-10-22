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
	orderDTOToDomain_UpsertActiveOrder(&data, req)

	result, code, err := u.orderRepo.UpsertActiveOrder(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.OrderGroupResponse
	orderDomainToDTO(&resp, result)

	return &resp, code, nil
}

// DTO to Domain
func orderDTOToDomain_UpsertActiveOrder(data *domain.OrderGroup, req *dto.OrderGroupUpsertRequest) {
	data.UUID = req.UUID
	data.CreatedBy = req.CreatedBy

	// delivery
	if req.Delivery != nil {
		customer := domain.Customer{
			Name:    req.Delivery.Customer.Name,
			Phone:   req.Delivery.Customer.Phone,
			Address: req.Delivery.Customer.Address,
		}

		deliveryCreatedAt, _ := time.Parse(time.RFC3339Nano, req.Delivery.CreatedAt)
		var schedludedAtTime *int64
		if req.Delivery.ScheduledAt != nil {
			scheduledAt, _ := time.Parse(time.RFC3339Nano, *req.Delivery.ScheduledAt)
			scheduledAtUnix := scheduledAt.UnixMicro()
			schedludedAtTime = &scheduledAtUnix
		}

		data.Delivery = &domain.Delivery{
			UUID:        req.Delivery.UUID,
			Number:      req.Delivery.Number,
			Partner:     req.Delivery.Partner,
			Driver:      req.Delivery.Driver,
			Customer:    customer,
			ScheduledAt: schedludedAtTime,
			CreatedAt:   deliveryCreatedAt.UnixMicro(),
		}
	}

	// queue
	if req.Queue != nil {
		customer := domain.Customer{
			Name:    req.Queue.Customer.Name,
			Phone:   req.Queue.Customer.Phone,
			Address: req.Queue.Customer.Address,
		}

		var ScheduledAtMicro *int64
		if req.Queue.ScheduledAt != nil {
			scheduledAt, _ := time.Parse(time.RFC3339Nano, *req.Queue.ScheduledAt)
			scheduled := scheduledAt.UnixMicro()
			ScheduledAtMicro = &scheduled
		}
		createdAt, _ := time.Parse(time.RFC3339Nano, req.Queue.CreatedAt)
		data.Queue = &domain.Queue{
			UUID:        req.Queue.UUID,
			Number:      req.Queue.Number,
			Customer:    customer,
			ScheduledAt: ScheduledAtMicro,
			CreatedAt:   createdAt.UnixMicro(),
		}
	}

	// space
	data.SpaceUUID = req.SpaceUUID

	// orders
	orders := []domain.Order{}
	for _, order := range req.Orders {
		// waiters
		var waiter *domain.WaiterOrder
		if order.Waiter != nil {
			waiter = &domain.WaiterOrder{
				UUID:       order.Waiter.UUID,
				BranchUUID: order.Waiter.BranchUUID,
				Name:       order.Waiter.Name,
			}
		}

		// order modifiers
		modifiers := []domain.ModifierOrder{}
		for _, modifier := range order.Modifiers {
			modifiers = append(modifiers, domain.ModifierOrder{
				UUID:     modifier.UUID,
				Name:     modifier.Name,
				Price:    modifier.Price,
				Quantity: modifier.Quantity,
			})
		}

		// item
		item := domain.ItemOrder{
			UUID:     order.Item.UUID,
			Name:     order.Item.Name,
			Label:    order.Item.Label,
			Price:    order.Item.Price,
			Quantity: order.Item.Quantity,
		}

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
			Waiter:    waiter,
			Note:      order.Note,
			CreatedAt: createdAt.UnixMicro(),
		})
	}
	data.Orders = orders

	// taxes
	taxes := []domain.TaxOrderGroup{}
	for _, tax := range req.Taxes {
		taxes = append(taxes, domain.TaxOrderGroup{
			UUID:  tax.UUID,
			Name:  tax.Name,
			Value: tax.Value,
		})
	}
	data.Taxes = taxes

	createdAt, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)
	data.CreatedAt = createdAt.UnixMicro()
}

// Domain to DTO
func orderDomainToDTO(resp *dto.OrderGroupResponse, data *domain.OrderGroup) {
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
		var deliveryUpdatedAt *time.Time
		if data.Delivery.UpdatedAt != nil {
			updatedAt := time.UnixMicro(*data.Delivery.UpdatedAt).UTC()
			deliveryUpdatedAt = &updatedAt
		}
		var deliveryDeletedAt *time.Time
		if data.Delivery.DeletedAt != nil {
			deletedAt := time.UnixMicro(*data.Delivery.DeletedAt).UTC()
			deliveryDeletedAt = &deletedAt
		}

		resp.Delivery = &dto.DeliveryResponse{
			UUID:        data.Delivery.UUID,
			Number:      data.Delivery.Number,
			Partner:     data.Delivery.Partner,
			Driver:      data.Delivery.Driver,
			Customer:    customer,
			ScheduledAt: deliveryScheduledAt,
			CreatedAt:   time.UnixMicro(data.Delivery.CreatedAt).UTC(),
			UpdatedAt:   deliveryUpdatedAt,
			DeletedAt:   deliveryDeletedAt,
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
		createdAt := time.UnixMicro(data.Queue.CreatedAt)
		resp.Queue = &dto.Queue{
			UUID:        data.Queue.UUID,
			Number:      data.Queue.Number,
			Customer:    customer,
			ScheduledAt: scheduledAt,
			CreatedAt:   createdAt.Format(time.RFC3339Nano),
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
			UUID:      order.UUID,
			Item:      item,
			Modifiers: modifiers,
			Discounts: discounts,
			Waiter:    waiter,
			Note:      order.Note,
			CreatedAt: time.UnixMicro(order.CreatedAt).Format(time.RFC3339Nano),
			UpdatedAt: orderUpdatedAt,
			DeletedAt: orderDeletedAt,
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
