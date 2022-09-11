package usecase

import (
	"lucy/cashier/domain"
	"time"
)

type employeeShiftUsecase struct {
	employeeShiftRepo domain.EmployeeShiftRepositoryContract
	contextTimeout    time.Duration
}

func NewEmployeeShiftUsecase(repo domain.EmployeeShiftRepositoryContract, timeout time.Duration) domain.EmployeeShiftUsecaseContract {
	return &employeeShiftUsecase{
		employeeShiftRepo: repo,
		contextTimeout:    timeout,
	}
}
