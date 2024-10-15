package product_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	testutils "github.com/sendydwi/online-book-store/commons"
	"github.com/sendydwi/online-book-store/services/product"
	"github.com/sendydwi/online-book-store/services/product/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_GetProductById_Repository(t *testing.T) {
	mockDB, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}
	repo := &product.ProductRepository{DB: mockDB}

	productId := 1
	product := entity.Product{ProductId: productId, Title: "Sample Product", IsActive: true}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE product_id = $1 AND is_active = $2 ORDER BY "products"."product_id" LIMIT $3`)).
			WithArgs(productId, true, 1).
			WillReturnRows(sqlmock.NewRows([]string{"product_id", "title", "is_active"}).
				AddRow(product.ProductId, product.Title, product.IsActive))

		result, err := repo.GetProductById(productId)

		assert.NoError(t, err)
		assert.Equal(t, &product, result)
	})

	t.Run("not_found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE product_id = $1 AND is_active = $2 ORDER BY "products"."product_id" LIMIT $3`)).
			WithArgs(productId, true, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := repo.GetProductById(productId)

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE product_id = $1 AND is_active = $2 ORDER BY "products"."product_id" LIMIT $3`)).
			WithArgs(productId, true, 1).
			WillReturnError(errors.New("database error"))

		result, err := repo.GetProductById(productId)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetProductList(t *testing.T) {
	mockDB, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}
	repo := &product.ProductRepository{DB: mockDB}

	// Case 1: Products found
	page, size := 1, 10
	products := []entity.Product{
		{ProductId: 1, Title: "Sample Product 1", IsActive: true},
		{ProductId: 2, Title: "Sample Product 2", IsActive: true},
	}
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" LIMIT $1`)).
			WithArgs(size).
			WillReturnRows(sqlmock.NewRows([]string{"product_id", "title", "is_active"}).
				AddRow(products[0].ProductId, products[0].Title, products[0].IsActive).
				AddRow(products[1].ProductId, products[1].Title, products[1].IsActive))

		result, err := repo.GetProductList(page, size)

		assert.NoError(t, err)
		assert.Equal(t, &products, result)
	})

	t.Run("success_page_3", func(t *testing.T) {
		page3 := 3
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" LIMIT $1 OFFSET $2`)).
			WithArgs(size, (page3-1)*size).
			WillReturnRows(sqlmock.NewRows([]string{"product_id", "title", "is_active"}).
				AddRow(products[0].ProductId, products[0].Title, products[0].IsActive).
				AddRow(products[1].ProductId, products[1].Title, products[1].IsActive))

		result, err := repo.GetProductList(page3, size)

		assert.NoError(t, err)
		assert.Equal(t, &products, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" LIMIT $1`)).
			WithArgs(size).
			WillReturnError(errors.New("database error"))

		result, err := repo.GetProductList(page, size)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
