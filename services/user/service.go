package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sendydwi/online-book-store/services/user/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	RegisterUser(email, password string) error
	Login(email, password string) (string, error)
}

type Service struct {
	Repo UserRepositoryInterface
}

func (s *Service) RegisterUser(email, password string) error {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if user != nil {
		return ErrEmailAlreadyUsed
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.Repo.RegisterUser(entity.User{
		UserId:    uuid.NewString(),
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Time{},
		CreatedBy: "application",
		UpdatedAt: time.Time{},
		UpdatedBy: "application",
	})

	if err != nil {
		return errors.New("failed to register user. try again later")
	}

	return nil
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrUserNotExist
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrWrongPassword
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
