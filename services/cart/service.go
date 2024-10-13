package cart

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	apicart "github.com/sendydwi/online-book-store/api/cart"
	"github.com/sendydwi/online-book-store/services/cart/entity"
	"github.com/sendydwi/online-book-store/services/product"
	"gorm.io/gorm"
)

type Service struct {
	Repo       CartRepository
	ProductSvc product.Service
}

func (s *Service) UpdateCartItem(updateRequest apicart.CartUpdateRequest, userId string) error {
	cart, err := s.getCurrentCart(userId)
	if err != nil {
		return err
	}

	cartItemUpdate := entity.CartItem{
		CartId:    cart.CartId,
		ProductId: updateRequest.ProductId,
		Quantity:  updateRequest.Quantity,
	}

	err = s.Repo.UpdateCartItem(cartItemUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) getCurrentCart(userId string) (*entity.Cart, error) {
	cart, err := s.Repo.GetCurrentActiveCart(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = &entity.Cart{
				CartId: uuid.NewString(),
				UserId: userId,
				Status: entity.CartStatusActive,
			}
			err := s.Repo.CreateActiveCart(*cart)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return cart, err
}

func (s *Service) GetCartItem(userId string) (*apicart.GetCartResponse, error) {
	cart, err := s.getCurrentCart(userId)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return &apicart.GetCartResponse{
			CartItems:  []apicart.CartItemResponse{},
			TotalPrice: 0,
		}, nil
	}

	cartItems, err := s.Repo.GetCartItemByCartId(cart.CartId)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	cartItemResponse := []apicart.CartItemResponse{}
	for _, item := range cartItems {
		wg.Add(1)
		go func(item entity.CartItem) {
			defer wg.Done()
			book, err := s.ProductSvc.Repo.GetProductById(item.ProductId)
			if err != nil {
				return
			}

			cartItemResponse = append(cartItemResponse, apicart.CartItemResponse{
				ProductId:     item.ProductId,
				Quantity:      item.Quantity,
				Price:         book.Price,
				SubtotalPrice: book.Price * float32(item.Quantity),
			})

		}(item)
	}
	wg.Wait()

	totalPrice := float32(0)
	for _, item := range cartItemResponse {
		totalPrice += item.SubtotalPrice
	}

	return &apicart.GetCartResponse{
		CartItems:  cartItemResponse,
		TotalPrice: totalPrice,
	}, nil
}
