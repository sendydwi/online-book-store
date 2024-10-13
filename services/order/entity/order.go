package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	OrderId         string    `json:"order_id" gorm:"primary_key"`
	UserId          string    `json:"user_id"`
	Status          string    `json:"status"`
	TotalPrice      float32   `json:"total_price"`
	DeliveryAddress string    `json:"address"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       string    `json:"updated_by"`
}

type OrderItem struct {
	OrderItemId     int             `json:"order_item_id" gorm:"primary_key"`
	OrderId         string          `json:"order_id"`
	ProductId       int             `json:"product_id"`
	SubtotalPrice   float32         `json:"subtotal_price"`
	Quantity        int             `json:"quantity"`
	ProductSnapshot ProductSnapshot `json:"product_snapshot" gorm:"type:jsonb"`
	CreatedAt       time.Time       `json:"created_at"`
	CreatedBy       string          `json:"created_by"`
	UpdatedAt       time.Time       `json:"updated_at"`
	UpdatedBy       string          `json:"updated_by"`
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

func (p *ProductSnapshot) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan error: expected []byte, got %T", value)
	}

	return json.Unmarshal(bytes, p)
}

func (p ProductSnapshot) Value() (driver.Value, error) {
	return json.Marshal(p)
}

const (
	OrderStatusWaitingForPayment = "WAITING_FOR_PAYMENT"
	OrderStatisComplete          = "COMPLETE"
)
