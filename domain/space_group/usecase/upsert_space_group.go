package usecase

import (
	"context"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) UpsertSpaceGroup(c context.Context, branchId string, req *domain.SpaceGroupUpsertRequest) (*domain.SpaceGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err := validator.IsUUID(req.UUID); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, req.CreatedAt)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	if err := validateSpaceGroupShape(req.Shape); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	fmt.Println(branchId)
	var data domain.SpaceGroup
	data.UUID = req.UUID
	data.BranchUUID = branchId
	data.Name = req.Name
	data.Code = req.Code
	data.Pax = req.Pax
	data.Shape = req.Shape
	data.CreatedAt = createdAt.UTC().UnixMicro()

	result, code, err := u.spaceGroupRepo.UpsertSpaceGroup(ctx, &data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.SpaceGroupResponse
	resp.UUID = result.UUID
	resp.BranchUUID = result.BranchUUID
	resp.Name = result.Name
	resp.Code = result.Code
	resp.Shape = result.Shape
	resp.Pax = result.Pax
	resp.CreatedAt = time.UnixMicro(result.CreatedAt).UTC()
	if result.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*result.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}
	if result.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*result.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	return &resp, http.StatusOK, nil
}

func validateSpaceGroupShape(shape domain.SpaceGroupShape) error {
	switch shape {
	case "circle":
		return nil
	case "square":
		return nil
	default:
		return domain.ErrSpaceGroupShapeInvalid
	}
}
