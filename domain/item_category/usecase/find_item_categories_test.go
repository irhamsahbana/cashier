package usecase_test

// import (
// 	"errors"
// 	"testing"

// 	"lucy/cashier/domain/menu_category/mocks"
// 	"lucy/cashier/domain/menu_category/usecase"
// 	"net/http"
// 	"time"

// 	"lucy/cashier/domain"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestFindMenuCategories_normalCase(t *testing.T) {
// 	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
// 	createdAtUnix := createdAt.UTC().UnixMicro()

// 	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
// 	updatedAtUnix := updatedAt.UTC().UnixMicro()

// 	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
// 	deletedAtUnix := deletedAt.UTC().UnixMicro()

// 	menuCategories := []domain.MenuCategory{
// 		{
// 			UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
// 			BranchUUID: "56ede4f9-91a6-4189-a13d-c9cc3e767d53",
// 			Name: "Coffee Base",
// 			CreatedAt: createdAtUnix,
// 		},
// 		{
// 			UUID: "09f04bb0-2bd7-4af0-9dc5-c59af7f938de",
// 			BranchUUID: "5ece4124-e588-404f-a068-1c3c503768fe",
// 			Name: "Juice",
// 			CreatedAt: createdAtUnix,
// 			UpdatedAt: &updatedAtUnix,
// 		},
// 	}

// 	mockMenuCategoryRepository.On("FindMenuCategories", ctx, false).Return(menuCategories, http.StatusOK, nil)

// 	resp, code, err := testMenuCategoryUsecase.FindMenuCategories(ctx, false)

// 	assert.NotNil(t, resp)
// 	assert.Equal(t, code, http.StatusOK)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 2, len(resp))
// 	assert.Equal(t, createdAt, resp[0].CreatedAt)

// 	t.Run("should get deleted resources too when withTrashed parameter equal true", func(t *testing.T) {
// 		var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 		var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 		menuCategories := []domain.MenuCategory{
// 			{
// 				UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
// 				BranchUUID: "56ede4f9-91a6-4189-a13d-c9cc3e767d53",
// 				Name: "Coffee Base",
// 				CreatedAt: createdAtUnix,
// 			},
// 			{
// 				UUID: "09f04bb0-2bd7-4af0-9dc5-c59af7f938de",
// 				BranchUUID: "5ece4124-e588-404f-a068-1c3c503768fe",
// 				Name: "Juice",
// 				Menus: []domain.Menu{
// 					{
// 						UUID: "4fac7aca-3346-4d0c-abeb-058ec1e29164",
// 						Name: "Mango",
// 						Price: 18000,
// 						CreatedAt: createdAtUnix,
// 						UpdatedAt: &updatedAtUnix,
// 					},
// 					{
// 						UUID: "89717271-5606-4a45-9b29-876c69d37770",
// 						Name: "Orange",
// 						Price: 15000,
// 						CreatedAt: createdAtUnix,
// 						DeletedAt: &deletedAtUnix,
// 					},
// 				},
// 				CreatedAt: createdAtUnix,
// 				UpdatedAt: &updatedAtUnix,
// 			},
// 			{
// 				UUID: "216b1552-715e-43f4-a9a4-e8ce4e658434",
// 				BranchUUID: "d5ad0c27-bea7-4844-828f-7dd7675cf8e1",
// 				Name: "Side dishes",
// 				CreatedAt: createdAtUnix,
// 				DeletedAt: &deletedAtUnix,
// 			},
// 		}
// 		mockMenuCategoryRepository.On("FindMenuCategories", ctx, true).Return(menuCategories, http.StatusOK, nil)
// 		resp, _, _ := testMenuCategoryUsecase.FindMenuCategories(ctx, true)

// 		assert.Equal(t, &deletedAt, resp[2].DeletedAt)
// 	})
// }

// func TestFindMenuCategories_ErrorWhenSomethingWrongWithMarshalingOrDatabase(t *testing.T) {
// 	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 	mockMenuCategoryRepository.On("FindMenuCategories", ctx, false).Return(nil, http.StatusInternalServerError, errors.New("mongo: request timeout"))

// 	resp, code, err := testMenuCategoryUsecase.FindMenuCategories(ctx, false)

// 	assert.Nil(t, resp)
// 	assert.Equal(t, http.StatusInternalServerError, code)
// 	assert.NotNil(t, err)
// }
