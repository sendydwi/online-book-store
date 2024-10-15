package product_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/product"
	mock_product "github.com/sendydwi/online-book-store/services/product/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_GetProductDetail_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_product.NewMockProductServiceInterface(ctrl)
	handler := product.ProductHandler{Svc: mockSvc}

	router.GET("/products/:id", handler.GetProductDetail)
	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetProductById(1).Return(&apiproduct.ProductResponse{
				ProductDetail: apiproduct.ProductDetail{},
				Stock:         10,
				Price:         100,
			}, nil).Times(1),
		)

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var response api.GenericResponseWithData
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	t.Run("not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("product_id_invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal_service_error", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetProductById(1).Return(nil, errors.New("service error")).Times(1),
		)

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

}

func Test_GetProductList_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mock_product.NewMockProductServiceInterface(ctrl)
	handler := product.ProductHandler{Svc: mockSvc}

	router.GET("/products", handler.GetProductList)

	t.Run("success", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetProductList(1, 10).Return(&apiproduct.ProductListResponse{ProductList: []apiproduct.ProductResponse{}}, nil).Times(1),
		)

		req := httptest.NewRequest(http.MethodGet, "/products?page=1&size=10", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response api.GenericResponseWithData
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	t.Run("page_invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?page=abc&size=10", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("size_invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?page=1&size=xyz", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal_service_error", func(t *testing.T) {
		gomock.InOrder(
			mockSvc.EXPECT().GetProductList(1, 10).Return(nil, errors.New("service error")).Times(1),
		)

		req := httptest.NewRequest(http.MethodGet, "/products?page=1&size=10", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
