package usecase

import (
	"context"
	"errors"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/lib/dto"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func NewOrderUsecase(repo domain.OrderRepositoryContract, timeout time.Duration) domain.OrderUsecaseContract {
	return &orderUsecase{
		orderRepo:      repo,
		contextTimeout: timeout,
	}
}

type orderUsecase struct {
	orderRepo      domain.OrderRepositoryContract
	contextTimeout time.Duration
}

func (u *orderUsecase) UpsertOrder(c context.Context, branchId string, req *dto.OrderGroupUpsertRequest) (*dto.OrderGroup, int, error) {
	panic("implement me")

	err := validateUpsertOrderRequest(req)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data *domain.OrderGroup
	data.BranchUUID = branchId

	return nil, 0, nil
}

func validateUpsertOrderRequest(req *dto.OrderGroupUpsertRequest) error {
	var typeCount int

	var errMsg map[string]interface{}

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
		errMsg["type"] = "Only one type is allowed"
	}

	err := validator.IsUUID(req.UUID)
	if err != nil {
		errMsg["uuid"] = "Invalid UUID"
	}

	if len(req.Orders) == 0 {
		errMsg["orders"] = "Orders cannot be empty"
	}

	// validate order
	for orderIndex, order := range req.Orders {
		err = validator.IsUUID(order.UUID)
		if err != nil {
			errMsg["orders.["+fmt.Sprint(orderIndex)+"].uuid"] = "Invalid UUID"
		}

		err = validator.IsUUID(order.Item.UUID)
		if err != nil {
			errMsg["orders.["+fmt.Sprint(orderIndex)+"].item.uuid"] = "Invalid UUID"
		}

		if order.Item.Quantity <= 0 {
			errMsg["orders.["+fmt.Sprint(orderIndex)+"]item.quantity"] = "Invalid quantity"
		}

		for _, modifier := range order.Modifiers {
			err = validator.IsUUID(modifier.UUID)
			if err != nil {
				return errors.New("invalid modifier uuid")
			}

			if modifier.Quantity <= 0 {
				return errors.New("modifier quantity must be greater than 0")
			}
		}

		if order.Waiter != nil {
			err = validator.IsUUID(order.Waiter.UUID)
			if err != nil {
				return errors.New("invalid waiter uuid")
			}

			err = validator.IsUUID(order.Waiter.BranchUUID)
			if err != nil {
				return errors.New("invalid waiter branch uuid")
			}
		}
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

		if req.Delivery.Customer.Name == "" {
			return errors.New("delivery customer name is required")
		}
	}

	return nil
}
