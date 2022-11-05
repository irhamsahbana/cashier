package usecase

import (
	"context"
	"errors"
	"fmt"
	"lucy/cashier/domain"
	"lucy/cashier/lib/validator"
	"net/http"
	"time"
)

func (u *zoneUsecase) UpsertZones(ctx context.Context, branchId string, req *domain.ZoneUpsertRequest) (*domain.ZonesResponse, int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if err := validateUpsertZonesRequest(*req); err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	var data []domain.Zone

	for _, zone := range req.Zones {
		createdAt, err := time.Parse(time.RFC3339, zone.CreatedAt)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}

		data = append(data, domain.Zone{
			UUID:        zone.UUID,
			BranchUUID:  branchId,
			Name:        zone.Name,
			Description: zone.Description,
			SpaceGroups: zone.SpaceGroups,
			CreatedAt:   createdAt.UnixMicro(),
		})
	}

	result, code, err := u.zoneRepo.UpsertZones(ctx, branchId, data)
	if err != nil {
		return nil, code, err
	}

	var resp domain.ZonesResponse

	for _, zone := range result {
		var zoneResp domain.ZoneResponse

		zoneResp.UUID = zone.UUID
		zoneResp.BranchUUID = zone.BranchUUID
		zoneResp.Name = zone.Name
		zoneResp.Description = zone.Description
		zoneResp.CreatedAt = time.UnixMicro(zone.CreatedAt).UTC()

		if zone.UpdatedAt != nil {
			zoneRespUpdatedAt := time.UnixMicro(*zone.UpdatedAt).UTC()
			zoneResp.UpdatedAt = &zoneRespUpdatedAt
		}

		for _, spaceGroup := range zone.SpaceGroups {
			var spaceGroupResp domain.SpaceGroupResponse

			spaceGroupResp.UUID = spaceGroup.UUID
			spaceGroupResp.BranchUUID = spaceGroup.BranchUUID
			spaceGroupResp.Name = spaceGroup.Name
			spaceGroupResp.Code = spaceGroup.Code
			spaceGroupResp.Shape = spaceGroup.Shape
			spaceGroupResp.Pax = spaceGroup.Pax
			spaceGroupResp.CreatedAt = time.UnixMicro(spaceGroup.CreatedAt).UTC()

			for _, space := range spaceGroup.Spaces {
				var spaceResp domain.SpaceResponse
				spaceResp.UUID = space.UUID
				spaceResp.Number = space.Number
				spaceResp.Occupied = space.Occupied
				spaceResp.CreatedAt = time.UnixMicro(space.CreatedAt).UTC()

				if space.UpdatedAt != nil {
					spaceRespUpdatedAt := time.UnixMicro(*space.UpdatedAt).UTC()
					spaceResp.UpdatedAt = &spaceRespUpdatedAt
				}

				if space.DeletedAt != nil {
					spaceRespDeletedAt := time.UnixMicro(*space.DeletedAt).UTC()
					spaceResp.DeletedAt = &spaceRespDeletedAt
				}

				spaceGroupResp.Spaces = append(spaceGroupResp.Spaces, spaceResp)
			}

			if len(spaceGroupResp.Spaces) == 0 {
				spaceGroupResp.Spaces = make([]domain.SpaceResponse, 0)
			}

			if spaceGroup.UpdatedAt != nil {
				spaceGroupRespUpdatedAt := time.UnixMicro(*spaceGroup.UpdatedAt).UTC()
				spaceGroupResp.UpdatedAt = &spaceGroupRespUpdatedAt
			}

			if spaceGroup.DeletedAt != nil {
				spaceGroupRespDeletedAt := time.UnixMicro(*spaceGroup.DeletedAt).UTC()
				spaceGroupResp.DeletedAt = &spaceGroupRespDeletedAt
			}

			zoneResp.SpaceGroups = append(zoneResp.SpaceGroups, spaceGroupResp)
		}

		if len(zone.SpaceGroups) == 0 {
			zoneResp.SpaceGroups = make([]domain.SpaceGroupResponse, 0)
		}

		resp.Zones = append(resp.Zones, zoneResp)
	}

	if len(result) == 0 {
		resp.Zones = make([]domain.ZoneResponse, 0)
	}

	return &resp, http.StatusOK, nil
}

func validateUpsertZonesRequest(req domain.ZoneUpsertRequest) error {
	spaceGroupIds := make(map[string]bool)

	for _, request := range req.Zones {
		if err := validator.Uuid(request.UUID); err != nil {
			return errors.New("invalid uuid")
		}

		for index, spaceGroup := range request.SpaceGroups {
			if err := validator.Uuid(spaceGroup); err != nil {
				return errors.New(fmt.Sprintf("invalid space group uuid at index %v (%v)", index, spaceGroup))
			}

			if _, ok := spaceGroupIds[spaceGroup]; ok {
				return errors.New(fmt.Sprintf("duplicate space group uuid at index %v (%v)", index, spaceGroup))
			}

			spaceGroupIds[spaceGroup] = true
		}
	}

	return nil
}
