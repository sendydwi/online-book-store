package model

type UserModel struct {
	UserId   string `json:"user_id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
