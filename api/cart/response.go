package apicart

type GetCartResponse struct {
	CartItems  []CartItemResponse `json:"cart_item"`
	TotalPrice float32            `json:"total_price"`
}

type CartItemResponse struct {
	ProductId     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float32 `json:"price"`
	SubtotalPrice float32 `json:"subtotal_price"`
}
