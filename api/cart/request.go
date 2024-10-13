package apicart

type CartUpdateRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
