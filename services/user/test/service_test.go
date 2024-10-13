package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sendydwi/online-book-store/services/user"
	"github.com/sendydwi/online-book-store/services/user/entity"
	mock_user "github.com/sendydwi/online-book-store/services/user/mocks"
)

func Test_Login_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_user.NewMockUserRepositoryInterface(ctrl)

	service := user.Service{
		Repo: mockRepo,
	}

	gomock.InOrder(
		mockRepo.EXPECT().GetUserByEmail("email").Return(&entity.User{
			UserId:   "123",
			Email:    "email",
			Password: "$2a$10$HD8D6eK7/DV5gpueVyRJKueCOsAw9SWuiL4Z6ojU6eo7sLpHC3OtK",
		}, nil).Times(1),
	)

	token, err := service.Login("email", "this_password")
	if err != nil {
		t.Fatal("login should not return error", err)
	}

	if token == "" {
		t.Fatal("token should not return empty", err)
	}
}
