package cart

import (
	"errors"

	"github.com/sendydwi/online-book-store/services/cart/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository struct {
	DB *gorm.DB
}

func (r *CartRepository) GetCurrentActiveCart(userId string) (*entity.Cart, error) {
	var cart entity.Cart
	err := r.DB.Where("user_id = ? AND status=?", userId, entity.CartStatusActive).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) CreateActiveCart(cart entity.Cart) error {
	err := r.DB.Create(&cart).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) UpdateCartItem(cartItem entity.CartItem) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var existingCartItem entity.CartItem
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("cart_id = ? AND product_id = ?", cartItem.CartId, cartItem.ProductId).First(existingCartItem).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err == nil {
			cartItem.CartItemId = existingCartItem.CartItemId
		}

		if err := tx.Save(cartItem).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) GetCartItemByCartId(cartId string) ([]entity.CartItem, error) {
	var cartItems []entity.CartItem
	err := r.DB.Where("cart_id = ?", cartId).Find(&cartItems).Error

	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
