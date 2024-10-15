package cart_test

import (
	"errors"
	"testing"

	apicart "github.com/sendydwi/online-book-store/api/cart"
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/cart"
	"github.com/sendydwi/online-book-store/services/cart/entity"
	mock_cart "github.com/sendydwi/online-book-store/services/cart/mock"
	mock_product "github.com/sendydwi/online-book-store/services/product/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_UpdateCartItem_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_cart.NewMockCartRepositoryInterface(ctrl)

	service := cart.Service{
		Repo: mockRepo,
	}

	// Test data
	userId := "user123"
	updateRequest := apicart.CartUpdateRequest{
		ProductId: 1,
		Quantity:  2,
	}
	cart := entity.Cart{
		CartId: "cart123",
		UserId: userId,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(&cart, nil).Times(1),
			mockRepo.EXPECT().UpdateCartItem(gomock.Any()).Return(nil).Times(1),
		)

		err := service.UpdateCartItem(updateRequest, userId)

		assert.NoError(t, err)
	})

	t.Run("service_error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(nil, errors.New("service error")).Times(1),
		)

		err := service.UpdateCartItem(updateRequest, userId)

		assert.Error(t, err)
	})
}

func Test_GetCurrentCart_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_cart.NewMockCartRepositoryInterface(ctrl)

	service := cart.Service{
		Repo: mockRepo,
	}

	// Test data
	userId := "user123"
	cart := &entity.Cart{
		CartId: "cart123",
		UserId: userId,
	}

	t.Run("cart_exist", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(cart, nil).Times(1),
		)

		result, err := service.GetCurrentCart(userId)

		assert.NoError(t, err)
		assert.Equal(t, cart, result)
	})

	t.Run("no_cart_exist", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(nil, gorm.ErrRecordNotFound).Times(1),
			mockRepo.EXPECT().CreateActiveCart(gomock.Any()).Return(nil).Times(1),
		)

		result, err := service.GetCurrentCart(userId)

		assert.NoError(t, err)
		assert.NotEmpty(t, result.CartId)
	})
}

func Test_GetCartItem_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_cart.NewMockCartRepositoryInterface(ctrl)
	mockProduct := mock_product.NewMockProductServiceInterface(ctrl)

	service := cart.Service{
		Repo:       mockRepo,
		ProductSvc: mockProduct,
	}

	// Test data
	userId := "user123"
	cart := &entity.Cart{CartId: "cart123"}
	cartItems := []entity.CartItem{
		{CartId: "cart123", ProductId: 1, Quantity: 1},
		{CartId: "cart123", ProductId: 2, Quantity: 2},
	}
	products := map[int]*apiproduct.ProductResponse{
		1: {Price: 10.0},
		2: {Price: 15.0},
	}

	expect := &apicart.GetCartResponse{CartItems: []apicart.CartItemResponse{
		{ProductId: 1, Quantity: 1, Price: 10, SubtotalPrice: 10},
		{ProductId: 2, Quantity: 2, Price: 15, SubtotalPrice: 30}},
		TotalPrice: 40}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(cart, nil).Times(1),
			mockRepo.EXPECT().GetCartItemByCartId(cart.CartId).Return(cartItems, nil).Times(1),
			mockProduct.EXPECT().GetProductById(1).Return(products[1], nil),
			mockProduct.EXPECT().GetProductById(2).Return(products[2], nil),
		)

		result, err := service.GetCartItem(userId)

		assert.NoError(t, err)
		assert.Equal(t, expect, result)

		assert.NoError(t, err)
		assert.Equal(t, float32(40.0), result.TotalPrice)
	})

	t.Run("service_error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(cart, errors.New("service error")).Times(1),
		)

		_, err := service.GetCartItem(userId)

		assert.Error(t, err)
	})
}

func Test_UpdateCartStatusToOrdered_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_cart.NewMockCartRepositoryInterface(ctrl)

	service := cart.Service{
		Repo: mockRepo,
	}

	userId := "user123"
	cart := &entity.Cart{
		CartId: "cart123",
		Status: entity.CartStatusOrdered,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(cart, nil).Times(1),
			mockRepo.EXPECT().UpdateCartStatus(*cart).Return(nil).Times(1),
		)

		err := service.UpdateCartStatusToOrdered(userId)

		assert.NoError(t, err)
		assert.Equal(t, entity.CartStatusOrdered, cart.Status)
	})

	t.Run("service_error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetCurrentActiveCart(userId).Return(nil, errors.New("service error")).Times(1),
		)

		err := service.UpdateCartStatusToOrdered(userId)

		assert.Error(t, err)
	})

}
