package adapter

import (
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/product/model"
)

func BookModelToProductResponse(model model.BookModel) apiproduct.ProductResponse {
	bookDetail := apiproduct.BookDetail{}
	bookDetail.ISBN = model.ISBN
	bookDetail.Author = model.Author
	bookDetail.Description = model.Description
	bookDetail.Title = model.Title
	bookDetail.Subtitle = model.Subtitle
	bookDetail.Publisher = model.Publisher
	bookDetail.PublishTime = model.PublishTime

	response := apiproduct.ProductResponse{
		BookDetail: bookDetail,
		Stock:      model.AvailableStock,
	}
	return response
}

func BookModelListToProductResponseList(model []model.BookModel) []apiproduct.ProductResponse {
	responseList := []apiproduct.ProductResponse{}
	for _, m := range model {
		response := BookModelToProductResponse(m)
		responseList = append(responseList, response)
	}

	return responseList
}
