package user

import (
	"fmt"

	"github.com/sendydwi/online-book-store/services/user/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	RegisterUser(user entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}
type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) RegisterUser(user entity.User) error {
	fmt.Println(user)
	err := r.DB.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Model(entity.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
