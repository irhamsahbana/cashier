package usecase

import (
	"context"
	"errors"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *orderUsecase) UpsertActiveOrder(c context.Context, branchId string, req *dto.OrderGroupUpsertRequest) (*dto.OrderGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err := validateUpsertActiveOrderRequest(req)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data domain.OrderGroup
	data.BranchUUID = branchId
	orderDTOToDomain_UpsertActiveOrder(&data, req)

	result, code, err := u.orderRepo.UpsertActiveOrder(ctx, branchId, &data)
	if err != nil {
		return nil, code, err
	}

	var resp dto.OrderGroupResponse
	orderDomainToDTO_upsertActiveOrder(&resp, result)

	return &resp, code, nil
}

// validate request
func validateUpsertActiveOrderRequest(req *dto.OrderGroupUpsertRequest) error {
	// validate if order group is not double type (must be either delivery, queue, space, or none of them(quick order))
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
		return errors.New("order group type is double")
	}

	err := validator.IsUUID(req.UUID)
	if err != nil {
		return errors.New("order group uuid is not valid")
	}

	if len(req.Orders) == 0 {
		return errors.New("order group must have at least one order")
	}

	// validate delivery
	if req.Delivery != nil {
		err = validator.IsUUID(req.Delivery.UUID)
		if err != nil {
			return errors.New("invalid delivery uuid")
		}

		if req.Delivery.Number == "" {
			return errors.New("delivery number is required")
		}

		if req.Delivery.Partner == "" {
			return errors.New("delivery partner is required")
		}

		if req.Delivery.Driver == "" {
			return errors.New("delivery driver is required")
		}

		if _, err := time.Parse(time.RFC3339Nano, req.Delivery.CreatedAt); err != nil {
			return errors.New(fmt.Sprintf("invalid delivery created at: %s", err.Error()))
		}

		if req.Delivery.Customer.Name == "" {
			return errors.New("delivery customer name is required")
		}
	}

	// validate queue
	if req.Queue != nil {
		err = validator.IsUUID(req.Queue.UUID)
		if err != nil {
			return errors.New("invalid queue uuid")
		}

		if req.Queue.Number == "" {
			return errors.New("queue number is required")
		}

		if req.Queue.Customer.Name == "" {
			return errors.New("queue customer name is required")
		}

		if req.Queue.PromisedAt != nil {
			_, err := time.Parse(time.RFC3339, *req.Queue.PromisedAt)
			if err != nil {
				return errors.New("invalid queue promised at")
			}
		}
	}

	// validate order
	for orderIndex, order := range req.Orders {
		err = validator.IsUUID(order.UUID)
		if err != nil {
			return errors.New(fmt.Sprintf("order %d uuid is not valid", orderIndex))
		}

		err = validator.IsUUID(order.Item.UUID)
		if err != nil {
			return errors.New(fmt.Sprintf("order %d item uuid is not valid", orderIndex))
		}

		if order.Item.Quantity <= 0 {
			return errors.New(fmt.Sprintf("order %d item quantity is not valid, must be 0 or more", orderIndex))
		}

		_, err := time.Parse(time.RFC3339Nano, order.CreatedAt)
		if err != nil {
			return errors.New(fmt.Sprintf("order %d created at is not valid: %v", orderIndex, err))
		}

		// order modifiers
		for modIndex, modifier := range order.Modifiers {
			err = validator.IsUUID(modifier.UUID)
			if err != nil {
				return errors.New(fmt.Sprintf("order %d modifier %d uuid is not valid: %v", orderIndex, modIndex, err))
			}

			if modifier.Quantity < 0 {
				return errors.New(fmt.Sprintf("order %d modifier %d quantity is not valid: %v", orderIndex, modIndex, err))
			}
		}

		// order waiter
		if order.Waiter != nil {
			err = validator.IsUUID(order.Waiter.UUID)
			if err != nil {
				return errors.New(fmt.Sprintf("order %d waiter uuid is not valid: %v", orderIndex, err))
			}

			err = validator.IsUUID(order.Waiter.BranchUUID)
			if err != nil {
				return errors.New(fmt.Sprintf("order %d waiter branch uuid is not valid: %v", orderIndex, err))
			}
		}
	}

	return nil
}

// DTO to Domain
func orderDTOToDomain_UpsertActiveOrder(data *domain.OrderGroup, req *dto.OrderGroupUpsertRequest) {
	data.UUID = req.UUID
	data.CreatedBy = req.CreatedBy

	data.Tax = req.Tax
	data.Tip = req.Tip
	data.Completed = req.Completed

	// delivery
	if req.Delivery != nil {
		customer := domain.Customer{
			Name:    req.Delivery.Customer.Name,
			Phone:   req.Delivery.Customer.Phone,
			Address: req.Delivery.Customer.Address,
		}

		deliveryCreatedAt, _ := time.Parse(time.RFC3339Nano, req.Delivery.CreatedAt)
		data.Delivery = &domain.Delivery{
			UUID:      req.Delivery.UUID,
			Number:    req.Delivery.Number,
			Partner:   req.Delivery.Partner,
			Driver:    req.Delivery.Driver,
			Customer:  customer,
			CreatedAt: deliveryCreatedAt.UnixMicro(),
		}
	}

	// queue
	if req.Queue != nil {
		customer := domain.Customer{
			Name:    req.Queue.Customer.Name,
			Phone:   req.Queue.Customer.Phone,
			Address: req.Queue.Customer.Address,
		}

		var promisedAtMicro *int64
		if req.Queue.PromisedAt != nil {
			promisedAt, _ := time.Parse(time.RFC3339Nano, *req.Queue.PromisedAt)
			promised := promisedAt.UnixMicro()
			promisedAtMicro = &promised
		}
		createdAt, _ := time.Parse(time.RFC3339Nano, req.Queue.CreatedAt)
		data.Queue = &domain.Queue{
			UUID:       req.Queue.UUID,
			Number:     req.Queue.Number,
			Customer:   customer,
			PromisedAt: promisedAtMicro,
			CreatedAt:  createdAt.UnixMicro(),
		}
	}

	// space
	data.SpaceUUID = req.SpaceUUID

	// orders
	var orders []domain.Order
	for _, order := range req.Orders {
		var waiter *domain.WaiterOrder
		if order.Waiter != nil {
			waiter = &domain.WaiterOrder{
				UUID:       order.Waiter.UUID,
				BranchUUID: order.Waiter.BranchUUID,
				Name:       order.Waiter.Name,
			}
		}

		// order modifiers
		var modifiers []domain.ModifierOrder
		for _, modifier := range order.Modifiers {
			modifiers = append(modifiers, domain.ModifierOrder{
				UUID:     modifier.UUID,
				Name:     modifier.Name,
				Price:    modifier.Price,
				Quantity: modifier.Quantity,
			})
		}

		if len(modifiers) == 0 {
			modifiers = make([]domain.ModifierOrder, 0)
		}

		// item
		item := domain.ItemOrder{
			UUID:     order.Item.UUID,
			Name:     order.Item.Name,
			Label:    order.Item.Label,
			Price:    order.Item.Price,
			Quantity: order.Item.Quantity,
		}

		createdAt, _ := time.Parse(time.RFC3339Nano, order.CreatedAt)
		orders = append(orders, domain.Order{
			UUID:      order.UUID,
			Item:      item,
			Modifiers: modifiers,
			Waiter:    waiter,
			CreatedAt: createdAt.UnixMicro(),
		})
	}

	data.Orders = orders

	createdAt, _ := time.Parse(time.RFC3339Nano, req.CreatedAt)
	data.CreatedAt = createdAt.UnixMicro()
}

// Domain to DTO
func orderDomainToDTO_upsertActiveOrder(resp *dto.OrderGroupResponse, data *domain.OrderGroup) {
	resp.UUID = data.UUID
	resp.BranchUUID = data.BranchUUID
	resp.CreatedBy = data.CreatedBy

	resp.Tax = data.Tax
	resp.Tip = data.Tip
	resp.Completed = data.Completed
	resp.CreatedAt = time.UnixMicro(data.CreatedAt).UTC()
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
			UUID:      data.Delivery.UUID,
			Number:    data.Delivery.Number,
			Partner:   data.Delivery.Partner,
			Driver:    data.Delivery.Driver,
			Customer:  customer,
			CreatedAt: time.UnixMicro(data.Delivery.CreatedAt).UTC(),
			UpdatedAt: deliveryUpdatedAt,
			DeletedAt: deliveryDeletedAt,
		}
	}

	// queue
	if data.Queue != nil {
		customer := dto.Customer{
			Name:    data.Queue.Customer.Name,
			Phone:   data.Queue.Customer.Phone,
			Address: data.Queue.Customer.Address,
		}

		var promisedAt *string
		if data.Queue.PromisedAt != nil {
			promisedAtTime := time.UnixMicro(*data.Queue.PromisedAt)
			promisedAtStr := promisedAtTime.Format(time.RFC3339Nano)
			promisedAt = &promisedAtStr
		}
		createdAt := time.UnixMicro(data.Queue.CreatedAt)
		resp.Queue = &dto.Queue{
			UUID:       data.Queue.UUID,
			Number:     data.Queue.Number,
			Customer:   customer,
			PromisedAt: promisedAt,
			CreatedAt:  createdAt.Format(time.RFC3339Nano),
		}
	}

	// space
	resp.SpaceUUID = data.SpaceUUID

	// orders
	var orders []dto.OrderResponse
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
		var modifiers []dto.ModifierOrder
		for _, modifier := range order.Modifiers {
			modifiers = append(modifiers, dto.ModifierOrder{
				UUID:     modifier.UUID,
				Name:     modifier.Name,
				Quantity: modifier.Quantity,
				Price:    modifier.Price,
			})
		}

		if len(modifiers) == 0 {
			modifiers = make([]dto.ModifierOrder, 0)
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
			Waiter:    waiter,
			CreatedAt: time.UnixMicro(order.CreatedAt).Format(time.RFC3339Nano),
			UpdatedAt: orderUpdatedAt,
			DeletedAt: orderDeletedAt,
		})
	}
	resp.Orders = orders

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
