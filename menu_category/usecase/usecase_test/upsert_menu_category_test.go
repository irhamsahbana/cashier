package usecase_test

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/menu_category/mocks"
	"lucy/cashier/menu_category/usecase"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = &mocks.MockMenuCategoryRepository{Mock: mock.Mock{}}
var testUsecase = usecase.NewMenuCategoryUsecase(mockRepo, time.Duration(5) * time.Second)
func TestUpsertMenuCategoryValidateUUID(t *testing.T) {
	// begin preparation
	ctx := context.Background()

	request := domain.MenuCategoryUpsertRequest{
		UUID: "74c4a96b-b19c-4c32-9b94-d13f533144fe",
		Name: "",
		CreatedAt: "2014-06-18T15:00:00.000000Z",
	}

	var menucategory domain.MenuCategory
	menucategory.UUID = "something"

	mockRepo.On("UpsertMenuCategory").Return(&menucategory, http.StatusOK, nil)

	resp, _, _ := testUsecase.UpsertMenuCategory(ctx, &request)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, resp.Name, "")
}

func TestValidateCode(t *testing.T) {

}