package entity

import "time"

type User struct {
	UserId    string    `json:"user_id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
