package model

type UserModel struct {
	UserId   string `json:"user_id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
