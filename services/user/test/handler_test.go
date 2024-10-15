package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/services/user"
	mock_user "github.com/sendydwi/online-book-store/services/user/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_RegisterUser_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_user.NewMockUserServiceInterface(ctrl)
	handler := user.UserHandler{Svc: mockSvc}

	createTestContext := func(jsonPayload string) (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		return c, w
	}

	t.Run("invalid_json_request", func(t *testing.T) {
		c, w := createTestContext(`{invalid}}`)
		handler.RegisterUser(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("email_already_used", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().RegisterUser("email@example.com", "this_password").Return(user.ErrEmailAlreadyUsed).Times(1),
		)

		handler.RegisterUser(c)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "email already used")
	})

	t.Run("internal_server_error", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().RegisterUser("email@example.com", "this_password").Return(errors.New("service error")).Times(1),
		)

		handler.RegisterUser(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("success", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().RegisterUser("email@example.com", "this_password").Return(nil).Times(1),
		)

		handler.RegisterUser(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})
}

func Test_LoginUser_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_user.NewMockUserServiceInterface(ctrl)
	handler := user.UserHandler{Svc: mockSvc}

	createTestContext := func(jsonPayload string) (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		return c, w
	}

	t.Run("invalid_json_request", func(t *testing.T) {
		c, w := createTestContext(`{invalid}}`)
		handler.LoginUser(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("email_not_exist", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().Login("email@example.com", "this_password").Return("", user.ErrUserNotExist).Times(1),
		)

		handler.LoginUser(c)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "wrong email or password")
	})

	t.Run("wrong_password", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().Login("email@example.com", "this_password").Return("", user.ErrWrongPassword).Times(1),
		)

		handler.LoginUser(c)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "wrong email or password")
	})

	t.Run("internal_server_error", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().Login("email@example.com", "this_password").Return("", errors.New("service error")).Times(1),
		)

		handler.LoginUser(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("success", func(t *testing.T) {
		jsonPayload := `{"email":"email@example.com","password":"this_password"}`
		c, w := createTestContext(jsonPayload)

		gomock.InOrder(
			mockSvc.EXPECT().Login("email@example.com", "this_password").Return("valid_token", nil).Times(1),
		)

		handler.LoginUser(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "valid_token")
	})
}
