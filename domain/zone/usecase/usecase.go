package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type zoneUsecase struct {
	zoneRepo       domain.ZoneRepositoryContract
	contextTimeout time.Duration
}

func NewZoneUsecase(repo domain.ZoneRepositoryContract, timeout time.Duration) domain.ZoneUsecaseContract {
	return &zoneUsecase{
		zoneRepo:       repo,
		contextTimeout: timeout,
	}
}
