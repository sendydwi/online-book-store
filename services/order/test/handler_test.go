package order_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	apiorder "github.com/sendydwi/online-book-store/api/order"
	"github.com/sendydwi/online-book-store/services/order"
	mock_order "github.com/sendydwi/online-book-store/services/order/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateOrder_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_order.NewMockOrderServiceInterface(ctrl)
	orderHandler := &order.OrderHandler{Svc: mockService}

	createTestContext := func(jsonPayload string) (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		c.Request = req
		c.Set("userId", "user-id")
		return c, w
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockService.EXPECT().CreateOrder("user-id", gomock.Any()).Return(nil).Times(1),
		)
		requestBody := `{"address": "123 Test St"}`

		ctx, w := createTestContext(requestBody)
		orderHandler.CreateOrder(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "created")
	})

	t.Run("bad_request", func(t *testing.T) {
		requestBody := `{invalid}}`

		ctx, w := createTestContext(requestBody)
		orderHandler.CreateOrder(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		gomock.InOrder(
			mockService.EXPECT().CreateOrder("user-id", gomock.Any()).Return(errors.New("service error")).Times(1),
		)
		requestBody := `{"address": "123 Test St"}`

		ctx, w := createTestContext(requestBody)
		orderHandler.CreateOrder(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func Test_GetOrderDetail_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_order.NewMockOrderServiceInterface(ctrl)
	orderHandler := &order.OrderHandler{Svc: mockService}

	createTestContext := func() (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/orders/123", nil)
		c.Request = req
		c.Set("userId", "user-id")
		return c, w
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockService.EXPECT().GetOrderDetail("123", "user-id").Return(&apiorder.GetOrderDetailResponse{}, nil).Times(1),
		)

		ctx, w := createTestContext()
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "123",
			},
		}
		orderHandler.GetOrderDetail(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("bad_request", func(t *testing.T) {
		ctx, w := createTestContext()

		orderHandler.GetOrderDetail(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockService.EXPECT().GetOrderDetail("123", "user-id").Return(nil, errors.New("service error")).Times(1),
		)

		ctx, w := createTestContext()
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "123",
			},
		}
		orderHandler.GetOrderDetail(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func Test_GetOrderHistories_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_order.NewMockOrderServiceInterface(ctrl)
	orderHandler := &order.OrderHandler{Svc: mockService}

	createTestContext := func(url string) (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		c.Request = req
		c.Set("userId", "user-id")
		return c, w
	}

	t.Run("success", func(t *testing.T) {
		mockService.EXPECT().GetOrderHistories("user-id", 1, 10).Return(&apiorder.GetOrderHistoryResponse{}, nil)

		ctx, w := createTestContext("/orders?page=1&size=10")
		orderHandler.GetOrderHistories(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})

	t.Run("bad_request_page", func(t *testing.T) {
		ctx, w := createTestContext("/orders?page=abs&size=10")
		orderHandler.GetOrderHistories(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("bad_request_size", func(t *testing.T) {
		ctx, w := createTestContext("/orders?page=1&size=xyz")
		orderHandler.GetOrderHistories(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		mockService.EXPECT().GetOrderHistories("user-id", 1, 10).Return(nil, errors.New("service error"))

		ctx, w := createTestContext("/orders?page=1&size=10")
		orderHandler.GetOrderHistories(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
