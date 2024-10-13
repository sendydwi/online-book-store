package entity

import "time"

type Product struct {
	ProductId      int       `json:"product_id" gorm:"primary_key"`
	ISBN           string    `json:"isbn" gorm:"unique"`
	Title          string    `json:"title"`
	Subtitle       string    `json:"subtitle"`
	Author         string    `json:"author"`
	PublishTime    time.Time `json:"publish_time"`
	Publisher      string    `json:"publisher"`
	TotalPage      int       `json:"total_page"`
	Description    string    `json:"description"`
	TotalStock     int       `json:"total_stock"`
	AvailableStock int       `json:"available_stock"`
	IsActive       bool      `json:"is_active"`
	OnHoldStock    int       `json:"on_hold_stock"`
	Price          float32   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
}
