package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sendydwi/online-book-store/services/user/model"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo UserRepository
}

func (s *Service) RegisterUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	user, err := s.Repo.GetUserByEmail(email)
	// error checking still absurd
	if user != nil && err != nil {
		return errors.New("email already used")
	}

	err = s.Repo.RegisterUser(model.UserModel{
		UserId:   uuid.NewString(),
		Email:    email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return errors.New("failed to register user. try again later")
	}

	return nil
}
