package order_test

import (
	"errors"
	"testing"

	apicart "github.com/sendydwi/online-book-store/api/cart"
	apiorder "github.com/sendydwi/online-book-store/api/order"
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	mock_cart "github.com/sendydwi/online-book-store/services/cart/mock"
	"github.com/sendydwi/online-book-store/services/order"
	"github.com/sendydwi/online-book-store/services/order/entity"
	mock_order "github.com/sendydwi/online-book-store/services/order/mock"
	mock_product "github.com/sendydwi/online-book-store/services/product/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateOrder_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_order.NewMockOrderRepositoryInterface(ctrl)
	mockProductSvc := mock_product.NewMockProductServiceInterface(ctrl)
	mockCartSvc := mock_cart.NewMockCartServiceInterface(ctrl)

	service := &order.Service{
		CartSvc:    mockCartSvc,
		ProductSvc: mockProductSvc,
		Repo:       mockRepo,
	}

	userId := "user123"
	request := apiorder.CreateOrderRequest{
		Address: "123 Test St",
	}

	cartResponse := &apicart.GetCartResponse{
		TotalPrice: 100,
		CartItems: []apicart.CartItemResponse{
			{ProductId: 1, SubtotalPrice: 100, Quantity: 1},
		},
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockCartSvc.EXPECT().GetCartItem(userId).Return(cartResponse, nil).Times(1),
			mockProductSvc.EXPECT().GetProductById(1).Return(&apiproduct.ProductResponse{
				ProductDetail: apiproduct.ProductDetail{},
				Stock:         10,
				Price:         100,
			}, nil).Times(1),
			mockRepo.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(nil).Times(1),
			mockCartSvc.EXPECT().UpdateCartStatusToOrdered(userId).Return(nil).Times(1),
		)

		err := service.CreateOrder(userId, request)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		gomock.InOrder(
			mockCartSvc.EXPECT().GetCartItem(userId).Return(nil, errors.New("cart not found")).Times(1),
		)

		err := service.CreateOrder(userId, request)

		assert.Error(t, err)
		assert.EqualError(t, err, "cart not found")
	})
}

func Test_GetOrderDetail_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_order.NewMockOrderRepositoryInterface(ctrl)
	mockProductSvc := mock_product.NewMockProductServiceInterface(ctrl)
	mockCartSvc := mock_cart.NewMockCartServiceInterface(ctrl)

	service := &order.Service{
		CartSvc:    mockCartSvc,
		ProductSvc: mockProductSvc,
		Repo:       mockRepo,
	}

	orderId := "order123"
	userId := "user123"
	order := &entity.Order{
		OrderId: orderId,
		UserId:  userId,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetOrderById(orderId).Return(order, nil).Times(1),
			mockRepo.EXPECT().GetOrderItemByOrderId(orderId).Return([]entity.OrderItem{
				{OrderId: orderId, ProductId: 1},
			}, nil).Times(1),
		)

		response, err := service.GetOrderDetail(orderId, userId)

		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetOrderById(orderId).Return(nil, errors.New("order not found")).Times(1),
		)

		response, err := service.GetOrderDetail(orderId, userId)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.EqualError(t, err, "order not found")
	})

}

func Test_GetOrderHistories_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_order.NewMockOrderRepositoryInterface(ctrl)
	mockProductSvc := mock_product.NewMockProductServiceInterface(ctrl)
	mockCartSvc := mock_cart.NewMockCartServiceInterface(ctrl)

	service := &order.Service{
		CartSvc:    mockCartSvc,
		ProductSvc: mockProductSvc,
		Repo:       mockRepo,
	}

	userId := "user123"
	orders := []entity.Order{
		{OrderId: "order1", UserId: userId},
		{OrderId: "order2", UserId: userId},
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetOrderByUserId(userId).Return(orders, nil).Times(1),
		)

		response, err := service.GetOrderHistories(userId, 1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.OrderHistories, 2)
	})

	t.Run("error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetOrderByUserId(userId).Return(nil, errors.New("database error")).Times(1),
		)

		response, err := service.GetOrderHistories(userId, 1, 10)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.EqualError(t, err, "database error")
	})
}
