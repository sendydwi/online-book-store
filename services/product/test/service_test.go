package product_test

import (
	"errors"
	"testing"
	"time"

	"github.com/sendydwi/online-book-store/services/product"
	"github.com/sendydwi/online-book-store/services/product/entity"
	mock_product "github.com/sendydwi/online-book-store/services/product/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_GetProductById_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_product.NewMockProductRepositoryInterface(ctrl)

	service := product.Service{
		Repo: mockRepo,
	}

	// Case 1: Product found
	productEntity := &entity.Product{
		ProductId:      1,
		ISBN:           "978-3-16-148410-0",
		Author:         "John Doe",
		Description:    "A book about testing",
		Title:          "Testing in Go",
		Subtitle:       "A Comprehensive Guide",
		Publisher:      "Tech Publishers",
		PublishTime:    time.Now(),
		AvailableStock: 10,
		Price:          29.99,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetProductById(1).Return(productEntity, nil).Times(1),
		)

		response, err := service.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, productEntity.ProductId, response.ProductDetail.ProductId)
		assert.Equal(t, productEntity.ISBN, response.ProductDetail.ISBN)
	})

	t.Run("product_not_found", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetProductById(2).Return(nil, gorm.ErrRecordNotFound).Times(1),
		)

		response, err := service.GetProductById(2)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.True(t, errors.Is(err, product.ErrProductNotFound))
	})

	t.Run("database_error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetProductById(3).Return(nil, errors.New("database error")).Times(1),
		)

		response, err := service.GetProductById(3)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "database error", err.Error())
	})
}

func Test_GetProductList_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_product.NewMockProductRepositoryInterface(ctrl)

	service := product.Service{
		Repo: mockRepo,
	}

	// Case 1: Product list found
	products := []entity.Product{
		{
			ProductId:      1,
			ISBN:           "978-3-16-148410-0",
			Author:         "John Doe",
			Description:    "A book about testing",
			Title:          "Testing in Go",
			Subtitle:       "A Comprehensive Guide",
			Publisher:      "Tech Publishers",
			PublishTime:    time.Now(),
			AvailableStock: 10,
			Price:          29.99,
		},
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetProductList(1, 10).Return(&products, nil).Times(1),
		)

		response, err := service.GetProductList(1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, len(response.ProductList))
		assert.Equal(t, products[0].ProductId, response.ProductList[0].ProductDetail.ProductId)
	})

	t.Run("database_error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetProductList(1, 10).Return(nil, errors.New("database error")).Times(1),
		)

		response, err := service.GetProductList(1, 10)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "database error", err.Error())
	})
}
