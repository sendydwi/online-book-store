package order

import (
	"github.com/sendydwi/online-book-store/services/order/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) CreateOrder(order entity.Order, orderItem []*entity.OrderItem) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&order).Error
		if err != nil {
			return err
		}

		err = tx.Save(orderItem).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) DeleteOrder(order entity.Order) error {

	err := r.DB.Delete(&entity.OrderItem{}, "order_id = ?", order.OrderId).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&entity.Order{}, "order_id = ?", order.OrderId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetOrderById(orderId string) (*entity.Order, error) {
	var orderEntity entity.Order
	err := r.DB.Where("order_id = ?", orderId).First(&orderEntity).Error

	if err != nil {
		return nil, err
	}

	return &orderEntity, nil
}

func (r *OrderRepository) GetOrderItemByOrderId(orderId string) ([]entity.OrderItem, error) {
	var orderItemEntities []entity.OrderItem
	err := r.DB.Where("order_id = ?", orderId).Find(&orderItemEntities).Error

	if err != nil {
		return nil, err
	}

	return orderItemEntities, nil
}

func (r *OrderRepository) GetOrderByUserId(userId string) ([]entity.Order, error) {
	var orderEntities []entity.Order
	err := r.DB.Where("user_id = ?", userId).Find(&orderEntities).Error

	if err != nil {
		return nil, err
	}

	return orderEntities, nil
}
