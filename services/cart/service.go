package cart

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	apicart "github.com/sendydwi/online-book-store/api/cart"
	"github.com/sendydwi/online-book-store/services/cart/entity"
	"github.com/sendydwi/online-book-store/services/product"
	"gorm.io/gorm"
)

type CartServiceInterface interface {
	UpdateCartItem(updateRequest apicart.CartUpdateRequest, userId string) error
	GetCurrentCart(userId string) (*entity.Cart, error)
	GetCartItem(userId string) (*apicart.GetCartResponse, error)
	UpdateCartStatusToOrdered(userId string) error
}

type Service struct {
	Repo       CartRepositoryInterface
	ProductSvc product.ProductServiceInterface
}

func (s *Service) UpdateCartItem(updateRequest apicart.CartUpdateRequest, userId string) error {
	cart, err := s.GetCurrentCart(userId)
	if err != nil {
		return err
	}

	cartItemUpdate := entity.CartItem{
		CartId:    cart.CartId,
		ProductId: updateRequest.ProductId,
		Quantity:  updateRequest.Quantity,
		CreatedAt: time.Time{},
		CreatedBy: "application",
		UpdatedAt: time.Time{},
		UpdatedBy: "application",
	}

	err = s.Repo.UpdateCartItem(cartItemUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCurrentCart(userId string) (*entity.Cart, error) {
	cart, err := s.Repo.GetCurrentActiveCart(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = &entity.Cart{
				CartId:    uuid.NewString(),
				UserId:    userId,
				Status:    entity.CartStatusActive,
				CreatedAt: time.Time{},
				CreatedBy: "application",
				UpdatedAt: time.Time{},
				UpdatedBy: "application",
			}
			err := s.Repo.CreateActiveCart(*cart)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return cart, nil
}

func (s *Service) GetCartItem(userId string) (*apicart.GetCartResponse, error) {
	cart, err := s.GetCurrentCart(userId)
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

	wg := sync.WaitGroup{}
	wg.Add(len(cartItems))
	queue := make(chan *apicart.CartItemResponse)
	cartItemResponse := []apicart.CartItemResponse{}
	totalPrice := float32(0)

	for _, item := range cartItems {
		go func(item entity.CartItem) {
			defer wg.Done()
			book, err := s.ProductSvc.GetProductById(item.ProductId)
			if err != nil {
				queue <- nil
				return
			}

			queue <- &apicart.CartItemResponse{
				ProductId:     item.ProductId,
				Quantity:      item.Quantity,
				Price:         book.Price,
				SubtotalPrice: book.Price * float32(item.Quantity),
			}
		}(item)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for item := range queue {
			if item != nil {
				cartItemResponse = append(cartItemResponse, *item)
				totalPrice += item.SubtotalPrice
			}
		}
	}()
	wg.Wait()
	close(queue)
	<-done

	return &apicart.GetCartResponse{
		CartItems:  cartItemResponse,
		TotalPrice: totalPrice,
	}, nil
}

func (s *Service) UpdateCartStatusToOrdered(userId string) error {
	cart, err := s.GetCurrentCart(userId)
	if err != nil {
		return err
	}

	cart.Status = entity.CartStatusOrdered

	err = s.Repo.UpdateCartStatus(*cart)
	if err != nil {
		return err
	}

	return nil
}
