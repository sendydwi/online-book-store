package entity

type Cart struct {
	CartId string `json:"cart_id" gorm:"primary_key"`
	UserId string `json:"user_id"`
	Status string `json:"status"`
}

type CartItem struct {
	CartItemId int    `json:"cart_item_id" gorm:"primary_key"`
	CartId     string `json:"cart_id"`
	ProductId  int    `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

const (
	CartStatusActive  = "ACTIVE"
	CartStatusOrdered = "ORDERED"
)
