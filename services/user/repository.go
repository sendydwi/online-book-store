package user

import (
	"github.com/sendydwi/online-book-store/services/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) RegisterUser(user model.UserModel) error {
	err := r.DB.Model(model.UserModel{}).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	var user model.UserModel
	err := r.DB.Model(model.UserModel{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
