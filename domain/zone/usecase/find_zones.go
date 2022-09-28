package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *zoneUsecase) Zones(ctx context.Context, branchId string) (*domain.ZonesResponse, int, error) {
	result, statusCode, err := u.zoneRepo.Zones(ctx, branchId)
	if err != nil {
		return nil, statusCode, err
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

			if len(spaceGroup.Spaces) == 0 {
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
