package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
	apiorder "github.com/sendydwi/online-book-store/api/order"
	"github.com/sendydwi/online-book-store/services/cart"
	"github.com/sendydwi/online-book-store/services/order/adapter"
	"github.com/sendydwi/online-book-store/services/order/entity"
	"github.com/sendydwi/online-book-store/services/product"
)

type OrderServiceInterface interface {
	CreateOrder(userId string, request apiorder.CreateOrderRequest) error
	GetOrderDetail(orderid, userId string) (*apiorder.GetOrderDetailResponse, error)
	GetOrderHistories(userId string, page, limit int) (*apiorder.GetOrderHistoryResponse, error)
}

type Service struct {
	Repo       OrderRepositoryInterface
	CartSvc    cart.CartServiceInterface
	ProductSvc product.ProductServiceInterface
}

func (s *Service) CreateOrder(userId string, request apiorder.CreateOrderRequest) error {
	cartItems, err := s.CartSvc.GetCartItem(userId)
	if err != nil {
		return err
	}

	order := entity.Order{
		OrderId:         uuid.NewString(),
		UserId:          userId,
		Status:          entity.OrderStatusWaitingForPayment,
		TotalPrice:      cartItems.TotalPrice,
		DeliveryAddress: request.Address,
		CreatedAt:       time.Time{},
		CreatedBy:       "application",
		UpdatedAt:       time.Time{},
		UpdatedBy:       "application",
	}

	orderItems := []*entity.OrderItem{}

	for _, item := range cartItems.CartItems {
		product, err := s.ProductSvc.GetProductById(item.ProductId)
		if err != nil {
			return err
		}

		snapshot := adapter.ConvertToProductSnapshot(product.ProductDetail, product.Stock, product.Price)

		orderItem := entity.OrderItem{
			OrderId:         order.OrderId,
			ProductId:       item.ProductId,
			SubtotalPrice:   item.SubtotalPrice,
			Quantity:        item.Quantity,
			ProductSnapshot: snapshot,
			CreatedAt:       time.Time{},
			CreatedBy:       "application",
			UpdatedAt:       time.Time{},
			UpdatedBy:       "application",
		}

		orderItems = append(orderItems, &orderItem)
	}

	err = s.Repo.CreateOrder(order, orderItems)
	if err != nil {
		return err
	}

	err = s.CartSvc.UpdateCartStatusToOrdered(userId)
	if err != nil {
		s.Repo.DeleteOrder(order)
		return err
	}

	return nil
}

func (s *Service) GetOrderDetail(orderid, userId string) (*apiorder.GetOrderDetailResponse, error) {
	order, err := s.Repo.GetOrderById(orderid)
	if err != nil {
		return nil, err
	}

	if order.UserId != userId {
		return nil, errors.New("order not found")
	}

	orderItems, err := s.Repo.GetOrderItemByOrderId(orderid)
	if err != nil {
		return nil, err
	}

	responseItem := []apiorder.OrderItem{}
	for _, item := range orderItems {
		apiItem := adapter.ConvertToApiOrderItem(item)
		responseItem = append(responseItem, apiItem)
	}

	response := apiorder.GetOrderDetailResponse{
		Order:      adapter.ConvertToApiOrder(*order),
		OrderItems: responseItem,
	}
	return &response, nil
}

func (s *Service) GetOrderHistories(userId string, page, limit int) (*apiorder.GetOrderHistoryResponse, error) {
	orders, err := s.Repo.GetOrderByUserId(userId)
	if err != nil {
		return nil, err
	}

	responseOrder := []apiorder.Order{}
	for _, order := range orders {
		apiOrder := adapter.ConvertToApiOrder(order)
		responseOrder = append(responseOrder, apiOrder)
	}
	return &apiorder.GetOrderHistoryResponse{
		OrderHistories: responseOrder,
	}, nil
}
