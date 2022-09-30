package usecase

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"
)

func (u *spaceGroupUsecase) FindSpaceGroups(c context.Context, branchId string, withTrashed bool) ([]domain.SpaceGroupResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	result, code, err := u.spaceGroupRepo.FindSpaceGroups(ctx, branchId, withTrashed)
	if err != nil {
		return nil, code, err
	}

	var resp []domain.SpaceGroupResponse

	// sg -> space group, s -> space
	for _, sg := range result {
		var data domain.SpaceGroupResponse
		var spaces []domain.SpaceResponse

		data.UUID = sg.UUID
		data.BranchUUID = sg.BranchUUID
		data.Name = sg.Name
		data.Code = sg.Code
		data.Shape = sg.Shape
		data.Pax = sg.Pax
		data.Reservable = sg.Reservable
		data.Disabled = sg.Disabled
		data.CreatedAt = time.UnixMicro(sg.CreatedAt).UTC()
		if sg.UpdatedAt != nil {
			dataUpdatedAt := time.UnixMicro(*sg.UpdatedAt).UTC()
			data.UpdatedAt = &dataUpdatedAt
		}
		if sg.DeletedAt != nil {
			dataDeletedAt := time.UnixMicro(*sg.DeletedAt).UTC()
			data.DeletedAt = &dataDeletedAt
		}

		if len(sg.Spaces) > 0 {
			for _, s := range sg.Spaces {
				var dataSpace domain.SpaceResponse

				dataSpace.UUID = s.UUID
				dataSpace.Number = s.Number
				dataSpace.Occupied = s.Occupied
				dataSpace.Description = s.Description
				dataSpace.CreatedAt = time.UnixMicro(s.CreatedAt).UTC()
				dataSpace.CreatedAt = time.UnixMicro(s.CreatedAt).UTC()
				if s.UpdatedAt != nil {
					dataUpdatedAt := time.UnixMicro(*s.UpdatedAt).UTC()
					dataSpace.UpdatedAt = &dataUpdatedAt
				}
				if s.DeletedAt != nil {
					dataDeletedAt := time.UnixMicro(*s.DeletedAt).UTC()
					dataSpace.DeletedAt = &dataDeletedAt
				}

				spaces = append(spaces, dataSpace)
			}
		}

		data.Spaces = spaces
		resp = append(resp, data)
	}

	return resp, http.StatusOK, nil
}
