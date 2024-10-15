package user_test

import (
	"testing"

	"github.com/sendydwi/online-book-store/services/user"
	"github.com/sendydwi/online-book-store/services/user/entity"
	mock_user "github.com/sendydwi/online-book-store/services/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_RegisterUser_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepositoryInterface(ctrl)

	service := user.Service{
		Repo: mockRepo,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(nil, gorm.ErrRecordNotFound).Times(1),
			mockRepo.EXPECT().RegisterUser(gomock.Any()).Return(nil).Times(1),
		)

		err := service.RegisterUser("email@example.com", "this_password")

		assert.Nil(t, err, "register should return not error")
	})

	t.Run("email_already_exist", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(&entity.User{
				UserId:   "123",
				Email:    "email@example.com",
				Password: "$2a$10$HD8D6eK7/DV5gpueVyRJKueCOsAw9SWuiL4Z6ojU6eo7sLpHC3OtK",
			}, nil).Times(1),
		)

		err := service.RegisterUser("email@example.com", "this_password")

		assert.Equal(t, user.ErrEmailAlreadyUsed, err, "register should return error email already exist")
	})
	t.Run("error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(nil, gorm.ErrDuplicatedKey).Times(1),
		)

		err := service.RegisterUser("email@example.com", "this_password")

		assert.NotNil(t, err, "register should return error")
	})
}

func Test_Login_Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepositoryInterface(ctrl)

	service := user.Service{
		Repo: mockRepo,
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(&entity.User{
				UserId:   "123",
				Email:    "email@example.com",
				Password: "$2a$10$HD8D6eK7/DV5gpueVyRJKueCOsAw9SWuiL4Z6ojU6eo7sLpHC3OtK",
			}, nil).Times(1),
		)

		token, err := service.Login("email@example.com", "this_password")

		assert.Nil(t, err, "login should not return error")
		assert.NotEqual(t, "", token, "token should not return empty")
	})

	t.Run("email_not_found", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(nil, gorm.ErrRecordNotFound).Times(1),
		)

		_, err := service.Login("email@example.com", "this_password")

		assert.Equal(t, user.ErrUserNotExist, err, "register should return error user not exist")
	})

	t.Run("wrong_password", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(&entity.User{
				UserId:   "123",
				Email:    "email@example.com",
				Password: "$2a$10$HD8D6eK7/DV5gpueVyRJKueCOsAw9SWuiL4Z6ojU6eo7sLpHC3OtK",
			}, nil).Times(1),
		)

		_, err := service.Login("email@example.com", "this_password_false")

		assert.Equal(t, user.ErrWrongPassword, err, "register should return error wrong password")
	})

	t.Run("error", func(t *testing.T) {
		gomock.InOrder(
			mockRepo.EXPECT().GetUserByEmail("email@example.com").Return(nil, gorm.ErrEmptySlice).Times(1),
		)

		_, err := service.Login("email@example.com", "this_password")

		assert.NotNil(t, err, "register should return error but found error is nil")
	})
}
