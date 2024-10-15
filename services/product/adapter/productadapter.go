package adapter

import (
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/product/entity"
)

func ProductModelToProductResponse(model entity.Product) apiproduct.ProductResponse {
	bookDetail := apiproduct.ProductDetail{}
	bookDetail.ProductId = model.ProductId
	bookDetail.ISBN = model.ISBN
	bookDetail.Author = model.Author
	bookDetail.Description = model.Description
	bookDetail.Title = model.Title
	bookDetail.Subtitle = model.Subtitle
	bookDetail.Publisher = model.Publisher
	bookDetail.PublishTime = model.PublishTime

	response := apiproduct.ProductResponse{
		ProductDetail: bookDetail,
		Stock:         model.AvailableStock,
		Price:         model.Price,
	}
	return response
}

func ProductModelListToProductResponseList(model []entity.Product) []apiproduct.ProductResponse {
	responseList := []apiproduct.ProductResponse{}
	for _, m := range model {
		response := ProductModelToProductResponse(m)
		responseList = append(responseList, response)
	}

	return responseList
}
