package usecase_test

import (
	"context"
	"lucy/cashier/domain"
	"time"
)

var ctx =  context.Background()
var timeoutContext = time.Duration(5) * time.Second

var normalUpsertRequest = domain.MenuCategoryUpsertRequest{
	UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
	Name: "Coffee Base",
	CreatedAt: "2022-08-13T04:06:16.312789Z",
}