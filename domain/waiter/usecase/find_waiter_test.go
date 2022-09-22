package usecase_test

// import (
// 	"errors"
// 	"lucy/cashier/domain"
// 	"lucy/cashier/domain/waiter/mocks"
// 	"lucy/cashier/domain/waiter/usecase"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestFindWaiter(t *testing.T) {
// 	var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 	var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
// 	createdAtUnix := createdAt.UTC().UnixMicro()

// 	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
// 	updatedAtUnix := updatedAt.UTC().UnixMicro()

// 	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
// 	deletedAtUnix := deletedAt.UTC().UnixMicro()

// 	var request = domain.WaiterUpsertrequest{
// 		UUID: "a7b53fe3-8eca-4520-811c-beb641809eaf",
// 		Name: "Rem",
// 		CreatedAt: createdAtString,
// 	}

// 	var mockRepoOutput = domain.Waiter{
// 		UUID: request.UUID,
// 		Name: request.Name,
// 		CreatedAt: createdAtUnix,
// 	}

// 	mockRepo.On("FindWaiter", ctx, request.UUID, true).Return(&mockRepoOutput, http.StatusOK, nil)
// 	result, code, err := testUsecase.FindWaiter(ctx, request.UUID, true)

// 	assert.NotNil(t, result)
// 	assert.Equal(t, http.StatusOK, code)
// 	assert.Nil(t, err)

// 	t.Run("should return error if something wrong with repository or waiter not found", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepo.On("FindWaiter", ctx, request.UUID, true).Return(nil, http.StatusNotFound, errors.New("Waiter not found"))
// 		result, code, err := testUsecase.FindWaiter(ctx, request.UUID, true)

// 		assert.Nil(t, result)
// 		assert.Equal(t, http.StatusNotFound, code)
// 		assert.Error(t, err)
// 	})

// 	t.Run("should convert last active of waiter from unix time to time.Time if exists", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepoOutputWithLastActiveExists := mockRepoOutput
// 		mockRepoOutputWithLastActiveExists.LastActive = &createdAtUnix

// 		mockRepo.On("FindWaiter", ctx, request.UUID, true).Return(&mockRepoOutputWithLastActiveExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.FindWaiter(ctx, request.UUID, true)

// 		assert.NotNil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)

// 		lastActiveResult := result.LastActive
// 		assert.Equal(t, createdAt, *lastActiveResult)
// 	})

// 	t.Run("should convert updated at of waiter from unix time to time.Time if exists", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepoOutputWithUpdatedAtExists := mockRepoOutput
// 		mockRepoOutputWithUpdatedAtExists.UpdatedAt = &updatedAtUnix

// 		mockRepo.On("FindWaiter", ctx, request.UUID, true).Return(&mockRepoOutputWithUpdatedAtExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.FindWaiter(ctx, request.UUID, true)

// 		assert.NotNil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)

// 		updatedAtResult := result.UpdatedAt
// 		assert.Equal(t, updatedAt, *updatedAtResult)
// 	})

// 	t.Run("should convert deleted at of waiter from unix time to time.Time if exists", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepoOutputWithDeletedAtExists := mockRepoOutput
// 		mockRepoOutputWithDeletedAtExists.DeletedAt = &deletedAtUnix

// 		mockRepo.On("FindWaiter", ctx, request.UUID, true).Return(&mockRepoOutputWithDeletedAtExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.FindWaiter(ctx, request.UUID, true)

// 		assert.NotNil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)

// 		deletedAtResult := result.DeletedAt
// 		assert.Equal(t, deletedAt, *deletedAtResult)
// 	})
// }
