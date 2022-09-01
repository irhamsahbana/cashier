package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type spaceGroupUsecase struct {
	spaceGroupRepo domain.SpaceGroupRepositoryContract
	contextTimeout time.Duration
}

func NewSpaceGroupUsecase(repo domain.SpaceGroupRepositoryContract, timeout time.Duration) domain.SpaceGroupUsecaseContract {
	return &spaceGroupUsecase{
		spaceGroupRepo: repo,
		contextTimeout: timeout,
	}
}
