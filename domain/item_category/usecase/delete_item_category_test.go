package usecase_test

// import (
// 	"errors"
// 	"lucy/cashier/domain"
// 	"lucy/cashier/domain/menu_category/mocks"
// 	"lucy/cashier/domain/menu_category/usecase"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestDeleteMenuCategory_normalCase(t *testing.T) {
// 	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 	createdAt, _ := time.Parse(time.RFC3339, createdAtString)
// 	createdAtUnix := createdAt.UTC().UnixMicro()

// 	updatedAt, _ := time.Parse(time.RFC3339, updatedAtString)
// 	updatedAtUnix := updatedAt.UTC().UnixMicro()

// 	deletedAt, _ := time.Parse(time.RFC3339, deletedAtString)
// 	deletedAtUnix := deletedAt.UTC().UnixMicro()

// 	menuCategory := domain.MenuCategory{
// 		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
// 		Name: "Coffee Base",
// 		CreatedAt: createdAtUnix,
// 		UpdatedAt: &updatedAtUnix,
// 		DeletedAt: &deletedAtUnix,
// 	}

// 	mockMenuCategoryRepository.Mock.On("DeleteMenuCategory", ctx, menuCategory.UUID).Return(&menuCategory, http.StatusOK, nil)

// 	resp, code, err := testMenuCategoryUsecase.DeleteMenuCategory(ctx, menuCategory.UUID)

// 	assert.NotNil(t, resp)
// 	assert.Equal(t, http.StatusOK, code)
// 	assert.Nil(t, err)
// }

// func TestDeleteMenuCategory_successWhenDeleteMenuCategoryThatDoesntExists(t *testing.T) {
// 	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 	var request = domain.MenuCategoryUpsertRequest{
// 		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
// 		Name: "Coffee Base",
// 		CreatedAt: createdAtString,
// 	}

// 	mockMenuCategoryRepository.Mock.On("DeleteMenuCategory", ctx, request.UUID).Return(nil, http.StatusNotFound, nil)

// 	resp, code, err := testMenuCategoryUsecase.DeleteMenuCategory(ctx, request.UUID)

// 	assert.Nil(t, resp)
// 	assert.Equal(t, http.StatusOK, code)
// 	assert.Nil(t, err)
// }

// func TestDeleteMenuCategory_errorWhenSomethingWrongWithMarshalingOrDatabase(t *testing.T) {
// 	var mockMenuCategoryRepository = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
// 	var testMenuCategoryUsecase = usecase.NewMenuCategoryUsecase(mockMenuCategoryRepository, timeoutContext)

// 	var request = domain.MenuCategoryUpsertRequest{
// 		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
// 		Name: "Coffee Base",
// 		CreatedAt: createdAtString,
// 	}

// 	mockMenuCategoryRepository.Mock.On("DeleteMenuCategory", ctx, request.UUID).Return(nil, http.StatusInternalServerError, errors.New("something wrong"))

// 	resp, code, err := testMenuCategoryUsecase.DeleteMenuCategory(ctx, request.UUID)

// 	assert.Nil(t, resp)
// 	assert.Equal(t, http.StatusInternalServerError, code)
// 	assert.NotNil(t, err)
// }

// 	// t.Run("Should throw error when get error 500", func(t *testing.T) {
//     //     testMenuCategoryUsecase.Mock.On("GetVendorCompanyById", mock.Anything, "login").
//     //         Return(nil, helper.Error(codes.Internal, "company.service.GetVendorCompanyById", errors.New("ERROR"))).Once()

//     //     _, err := service.CreateDeliveryOrderFTL(mockContext, &payload)

//     //     assert.Error(t, err)
//     // })

//     // t.Run("Should throw error when vendor equal nil or vendor id is empty string", func(t *testing.T) {
//     //     companyUsecaseMock.On("GetVendorCompanyById", mock.Anything, "login").
//     //         Return(nil, nil).Once()

//     //     _, err := service.CreateDeliveryOrderFTL(mockContext, &payload)

//     //     assert.Error(t, err)
//     // })
