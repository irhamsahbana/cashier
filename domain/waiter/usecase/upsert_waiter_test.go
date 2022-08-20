package usecase_test

import (
	"lucy/cashier/domain"
	"lucy/cashier/domain/waiter/mocks"
	"lucy/cashier/domain/waiter/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpsertWaiter(t *testing.T) {
	var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
	var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
	createdAtUnix := createdAt.UTC().UnixMicro()

	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
	updatedAtUnix := updatedAt.UTC().UnixMicro()

	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
	deletedAtUnix := deletedAt.UTC().UnixMicro()

	var request = domain.WaiterUpsertrequest{
		UUID: "a7b53fe3-8eca-4520-811c-beb641809eaf",
		Name: "Rem",
		CreatedAt: createdAtString,
	}

	var mockRepoInputOutput = domain.Waiter{
		UUID: request.UUID,
		Name: request.Name,
		CreatedAt: createdAtUnix,
	}

	mockRepo.On("UpsertWaiter", ctx, &mockRepoInputOutput).Return(&mockRepoInputOutput, http.StatusOK, nil)
	result, code, err := testUsecase.UpsertWaiter(ctx, &request)

	assert.NotNil(t, result)
	assert.Equal(t, http.StatusOK, code)
	assert.Nil(t, err)

	t.Run("should return error if uuid is not valid uuid", func(t *testing.T) {
		requestWithInvalidUUID := request
		requestWithInvalidUUID.UUID = ""

		_, code, err := testUsecase.UpsertWaiter(ctx, &requestWithInvalidUUID)

		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.Error(t, err)
	})

	t.Run("should return error if created_at is not valid RFC3999", func(t *testing.T) {
		requestWithInvalidCreatedAt := request
		requestWithInvalidCreatedAt.CreatedAt = "2022-08-13T04:06:16.312xzyZ"

		result, code, err := testUsecase.UpsertWaiter(ctx, &requestWithInvalidCreatedAt)

		assert.Nil(t, result)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.Error(t, err)
	})

	t.Run("should return error if name of waiter is empty", func(t *testing.T) {
		requestWithEmptyName := request
		requestWithEmptyName.Name = ""

		result, code, err := testUsecase.UpsertWaiter(ctx, &requestWithEmptyName)

		assert.Nil(t, result)
		assert.Equal(t, http.StatusUnprocessableEntity, code)
		assert.Error(t, err)
	})

	t.Run("should return error if got something wrong with repository layer", func(t *testing.T) {
		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

		mockRepo.On("UpsertWaiter", ctx, &mockRepoInputOutput).Return(nil, http.StatusInternalServerError, errors.New("something wrong"))
		result, code, err := testUsecase.UpsertWaiter(ctx, &request)

		assert.Nil(t, result)
		assert.Equal(t, http.StatusInternalServerError, code)
		assert.Error(t, err)
	})

	t.Run("should convert last active of waiter from unix time to time.Time if exists", func(t *testing.T) {
		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

		mockRepoOutputWithLastActiveExists := mockRepoInputOutput
		mockRepoOutputWithLastActiveExists.LastActive = &createdAtUnix

		mockRepo.On("UpsertWaiter", ctx, &mockRepoInputOutput).Return(&mockRepoOutputWithLastActiveExists, http.StatusOK, nil)
		result, code, err := testUsecase.UpsertWaiter(ctx, &request)

		assert.NotNil(t, result)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)

		lastActiveResult := result.LastActive
		assert.Equal(t, createdAt, *lastActiveResult)
	})

	t.Run("should convert updated at of waiter from unix time to time.Time if exists", func(t *testing.T) {
		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

		mockRepoOutputWithUpdatedAtExists := mockRepoInputOutput
		mockRepoOutputWithUpdatedAtExists.UpdatedAt = &updatedAtUnix

		mockRepo.On("UpsertWaiter", ctx, &mockRepoInputOutput).Return(&mockRepoOutputWithUpdatedAtExists, http.StatusOK, nil)
		result, code, err := testUsecase.UpsertWaiter(ctx, &request)

		assert.NotNil(t, result)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)

		updatedAtResult := result.UpdatedAt
		assert.Equal(t, updatedAt, *updatedAtResult)
	})

	t.Run("should convert deleted at of waiter from unix time to time.Time if exists", func(t *testing.T) {
		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

		mockRepoOutputWithDeletedAtExists := mockRepoInputOutput
		mockRepoOutputWithDeletedAtExists.DeletedAt = &deletedAtUnix

		mockRepo.On("UpsertWaiter", ctx, &mockRepoInputOutput).Return(&mockRepoOutputWithDeletedAtExists, http.StatusOK, nil)
		result, code, err := testUsecase.UpsertWaiter(ctx, &request)

		assert.NotNil(t, result)
		assert.Equal(t, http.StatusOK, code)
		assert.Nil(t, err)
	})
}