package user

import (
	"errors"
	"fmt"

	"github.com/sendydwi/online-book-store/services/user/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) RegisterUser(user entity.User) error {
	err := r.DB.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Model(entity.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}
	return &user, nil
}
