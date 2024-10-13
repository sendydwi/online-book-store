package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sendydwi/online-book-store/services/user/entity"
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

	err = s.Repo.RegisterUser(entity.User{
		UserId:   uuid.NewString(),
		Email:    email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return errors.New("failed to register user. try again later")
	}

	return nil
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.UserId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
