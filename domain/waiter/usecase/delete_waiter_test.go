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

// func TestDeleteWaiter(t *testing.T) {
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

// 	mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(&mockRepoOutput, http.StatusOK, nil)
// 	result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

// 	assert.NotNil(t, result)
// 	assert.Equal(t, http.StatusOK, code)
// 	assert.Nil(t, err)

// 	t.Run("should return error when something wrong in repository layer", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(nil, http.StatusInternalServerError, errors.New("something wrong with database"))

// 		result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

// 		assert.Nil(t, result)
// 		assert.Equal(t, http.StatusInternalServerError, code)
// 		assert.Error(t, err)
// 	})

// 	t.Run("should return deleted success (status ok) when no resource was deleted", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(nil, http.StatusNotFound, nil)

// 		result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

// 		assert.Nil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("should convert last active of waiter from unix time to time.Time if exists", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepoOutputWithLastActiveExists := mockRepoOutput
// 		mockRepoOutputWithLastActiveExists.LastActive = &createdAtUnix

// 		mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(&mockRepoOutputWithLastActiveExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

// 		assert.NotNil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)

// 		lastActive := result.LastActive
// 		assert.Equal(t, createdAt, *lastActive)
// 	})

// 	t.Run("should convert updated at of waiter from unix time to time.Time if exists", func(t *testing.T) {
// 		var mockRepo = &mocks.MockWaiterRepository{Mock: mock.Mock{}}
// 		var testUsecase = usecase.NewWaiterUsecase(mockRepo, timeoutContext)

// 		mockRepoOutputWithUpdatedAtExists := mockRepoOutput
// 		mockRepoOutputWithUpdatedAtExists.UpdatedAt = &updatedAtUnix

// 		mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(&mockRepoOutputWithUpdatedAtExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

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

// 		mockRepo.On("DeleteWaiter", ctx, request.UUID).Return(&mockRepoOutputWithDeletedAtExists, http.StatusOK, nil)
// 		result, code, err := testUsecase.DeleteWaiter(ctx, request.UUID)

// 		assert.NotNil(t, result)
// 		assert.Equal(t, http.StatusOK, code)
// 		assert.Nil(t, err)

// 		deletedAtResult := result.DeletedAt
// 		assert.Equal(t, deletedAt, *deletedAtResult)
// 	})
// }
