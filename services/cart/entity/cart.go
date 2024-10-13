package entity

import "time"

type Cart struct {
	CartId    string    `json:"cart_id" gorm:"primary_key"`
	UserId    string    `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type CartItem struct {
	CartItemId int       `json:"cart_item_id" gorm:"primary_key"`
	CartId     string    `json:"cart_id"`
	ProductId  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}

const (
	CartStatusActive  = "ACTIVE"
	CartStatusOrdered = "ORDERED"
)
