package usecase

import (
	"context"
	"lucy/cashier/domain"
)

func (u *spaceGroupUsecase) FindSpaceGroups(c context.Context, withTrashed bool) ([]domain.SpaceGroupResponse, int, error) {
	panic("not implemented")
}
