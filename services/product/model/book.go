package model

import "time"

type BookModel struct {
	BookId         int       `json:"book_id" gorm:"primary_key"`
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
	OnHoldStock    int       `json:"on_hold_stock"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	LastUpdatedAt  time.Time `json:"last_updated_at"`
	LastUpdatedBy  time.Time `json:"last_updated_by"`
}
