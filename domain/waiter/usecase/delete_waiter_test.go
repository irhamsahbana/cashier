package usecase_test

import (
	"lucy/cashier/domain"
	"lucy/cashier/domain/waiter/mocks"
	"lucy/cashier/domain/waiter/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteWaiter(t *testing.T) {
	var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
	var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	var request = domain.WaiterUpsertrequest{
		UUID: "a7b53fe3-8eca-4520-811c-beb641809eaf",
		Name: "Rem",
		CreatedAt: createdAtString,
	}

	var mockOutput = domain.Waiter{
		UUID: request.UUID,
		Name: request.Name,
		CreatedAt: createdAtUnix,
	}

	mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(&mockOutput, http.StatusOK, nil)

	resp, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)
}