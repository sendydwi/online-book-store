package apiorder

import "time"

type GetOrderHistoryResponse struct {
	OrderHistories []Order `json:"order_histories"`
}

type GetOrderDetailResponse struct {
	Order      Order       `json:"order"`
	OrderItems []OrderItem `json:"order_items"`
}

type Order struct {
	ID              string  `json:"order_id" gorm:"primary_key"`
	UserId          string  `json:"user_id"`
	Status          string  `json:"status"`
	TotalPrice      float32 `json:"total_price"`
	DeliveryAddress string  `json:"address"`
}

type OrderItem struct {
	ID              int             `json:"order_item_id"`
	OrderId         string          `json:"order_id"`
	ProductId       int             `json:"product_id"`
	SubtotalPrice   float32         `json:"subtotal_price"`
	Quantity        int             `json:"quantity"`
	ProductSnapshot ProductSnapshot `json:"product_snapshot"`
}

type ProductSnapshot struct {
	ISBN           string    `json:"isbn" `
	Title          string    `json:"title"`
	Subtitle       string    `json:"subtitle"`
	Author         string    `json:"author"`
	PublishTime    time.Time `json:"publish_time"`
	Publisher      string    `json:"publisher"`
	TotalPage      int       `json:"total_page"`
	Description    string    `json:"description"`
	AvailableStock int       `json:"available_stock"`
	Price          float32   `json:"price"`
}
