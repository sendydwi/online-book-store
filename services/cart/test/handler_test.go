package cart_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apicart "github.com/sendydwi/online-book-store/api/cart"
	"github.com/sendydwi/online-book-store/services/cart"
	mock_cart "github.com/sendydwi/online-book-store/services/cart/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_UpdateCartItem_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_cart.NewMockCartServiceInterface(ctrl)
	handler := cart.CartHandler{Svc: mockSvc}

	reqBody := apicart.CartUpdateRequest{
		ProductId: 1,
		Quantity:  2,
	}

	jsonValue, _ := json.Marshal(reqBody)
	createTestContext := func(jsonPayload string) (*gin.Context, *httptest.ResponseRecorder) {
		req, _ := http.NewRequest(http.MethodPut, "/cart", strings.NewReader(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req
		ctx.Set("userId", "user-id")
		return ctx, w
	}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().UpdateCartItem(reqBody, "user-id").Return(nil).Times(1),
		)

		ctx, w := createTestContext(string(jsonValue))
		handler.UpdateCartItem(ctx)

		// Validate response
		assert.Equal(t, http.StatusOK, w.Code)
		var response api.GenericResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response.Message)
	})

	t.Run("invalid_json_request", func(t *testing.T) {
		ctx, w := createTestContext("{invalid}}")
		handler.UpdateCartItem(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed_update_cart", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().UpdateCartItem(reqBody, "user-id").Return(errors.New("service error")).Times(1),
		)

		ctx, w := createTestContext(string(jsonValue))
		handler.UpdateCartItem(ctx)

		// Validate response
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func Test_GetCart_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_cart.NewMockCartServiceInterface(ctrl)
	handler := cart.CartHandler{Svc: mockSvc}

	// Create test context

	createTestContext := func() (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/cart", nil)
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req
		ctx.Set("userId", "user-id")
		return ctx, w
	}

	// Mock the GetCartItem response
	mockCartItemResponse := apicart.GetCartResponse{}

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetCartItem("user-id").Return(&mockCartItemResponse, nil).Times(1),
		)

		ctx, w := createTestContext()
		handler.GetCart(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		var response api.GenericResponseWithData
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("failed", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetCartItem("user-id").Return(nil, errors.New("service error")).Times(1),
		)

		ctx, w := createTestContext()
		handler.GetCart(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

// func TestGetCart_Fail(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	// Mock service
// 	mockSvc := new(mocks.CartService)
// 	handler := handlers.CartHandler{Svc: mockSvc}

// 	// Create test context
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/cart", nil)
// 	ctx, _ := gin.CreateTestContext(w)
// 	ctx.Request = req
// 	ctx.Set("userId", "mockedUserId")

// 	// Mock the GetCartItem to return an error
// 	mockSvc.On("GetCartItem", "mockedUserId").Return(nil, errors.New("service error"))

// 	// Call the handler
// 	handler.GetCart(ctx)

// 	// Validate response
// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	mockSvc.AssertExpectations(t)
// }
