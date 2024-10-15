package product_test

import (
	"testing"
	"time"

	"github.com/sendydwi/online-book-store/services/product/adapter"
	"github.com/sendydwi/online-book-store/services/product/entity"
	"github.com/stretchr/testify/assert"
)

func Test_ProductModelToProductResponse(t *testing.T) {
	product := entity.Product{
		ProductId:      1,
		ISBN:           "978-3-16-148410-0",
		Author:         "John Doe",
		Description:    "A book about testing",
		Title:          "Testing in Go",
		Subtitle:       "A Comprehensive Guide",
		Publisher:      "Tech Publishers",
		PublishTime:    time.Now(),
		AvailableStock: 10,
		Price:          29.99,
	}

	result := adapter.ProductModelToProductResponse(product)

	assert.Equal(t, product.ProductId, result.ProductDetail.ProductId)
	assert.Equal(t, product.ISBN, result.ProductDetail.ISBN)
	assert.Equal(t, product.Author, result.ProductDetail.Author)
	assert.Equal(t, product.Description, result.ProductDetail.Description)
	assert.Equal(t, product.Title, result.ProductDetail.Title)
	assert.Equal(t, product.Subtitle, result.ProductDetail.Subtitle)
	assert.Equal(t, product.Publisher, result.ProductDetail.Publisher)
	assert.Equal(t, product.PublishTime, result.ProductDetail.PublishTime)
	assert.Equal(t, product.AvailableStock, result.Stock)
	assert.Equal(t, product.Price, result.Price)
}

func Test_ProductModelListToProductResponseList(t *testing.T) {
	products := []entity.Product{
		{
			ProductId:      1,
			ISBN:           "978-3-16-148410-0",
			Author:         "John Doe",
			Description:    "A book about testing",
			Title:          "Testing in Go",
			Subtitle:       "A Comprehensive Guide",
			Publisher:      "Tech Publishers",
			PublishTime:    time.Now(),
			AvailableStock: 10,
			Price:          29.99,
		},
		{
			ProductId:      2,
			ISBN:           "978-1-234-56789-7",
			Author:         "Jane Doe",
			Description:    "Another book about testing",
			Title:          "Advanced Testing in Go",
			Subtitle:       "Best Practices",
			Publisher:      "Code Masters",
			PublishTime:    time.Now(),
			AvailableStock: 5,
			Price:          39.99,
		},
	}

	result := adapter.ProductModelListToProductResponseList(products)

	assert.Equal(t, len(products), len(result))

	assert.Equal(t, products[0].ProductId, result[0].ProductDetail.ProductId)
	assert.Equal(t, products[0].ISBN, result[0].ProductDetail.ISBN)
	assert.Equal(t, products[0].Author, result[0].ProductDetail.Author)
	assert.Equal(t, products[0].Description, result[0].ProductDetail.Description)
	assert.Equal(t, products[0].Title, result[0].ProductDetail.Title)
	assert.Equal(t, products[0].Subtitle, result[0].ProductDetail.Subtitle)
	assert.Equal(t, products[0].Publisher, result[0].ProductDetail.Publisher)
	assert.Equal(t, products[0].PublishTime, result[0].ProductDetail.PublishTime)
	assert.Equal(t, products[0].AvailableStock, result[0].Stock)
	assert.Equal(t, products[0].Price, result[0].Price)

	assert.Equal(t, products[1].ProductId, result[1].ProductDetail.ProductId)
	assert.Equal(t, products[1].ISBN, result[1].ProductDetail.ISBN)
	assert.Equal(t, products[1].Author, result[1].ProductDetail.Author)
	assert.Equal(t, products[1].Description, result[1].ProductDetail.Description)
	assert.Equal(t, products[1].Title, result[1].ProductDetail.Title)
	assert.Equal(t, products[1].Subtitle, result[1].ProductDetail.Subtitle)
	assert.Equal(t, products[1].Publisher, result[1].ProductDetail.Publisher)
	assert.Equal(t, products[1].PublishTime, result[1].ProductDetail.PublishTime)
	assert.Equal(t, products[1].AvailableStock, result[1].Stock)
	assert.Equal(t, products[1].Price, result[1].Price)
}
