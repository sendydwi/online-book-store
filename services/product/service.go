package product

import (
	"errors"
	"fmt"

	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/product/adapter"
	"gorm.io/gorm"
)

type Service struct {
	Repo ProductRepository
}

func (s *Service) GetProductById(bookId int) (*apiproduct.ProductResponse, error) {
	book, err := s.Repo.GetProductById(bookId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book id %v not found", bookId)
		}
		return nil, err
	}
	bookResponse := adapter.ProductModelToProductResponse(*book)

	return &bookResponse, nil
}

func (s *Service) GetProductList(page, size int) (*apiproduct.ProductListResponse, error) {
	books, err := s.Repo.GetProductList(page, size)

	if err != nil {
		return nil, err
	}

	bookResponseList := adapter.ProductModelListToProductResponseList(*books)
	response := apiproduct.ProductListResponse{
		ProductList: bookResponseList,
		Page:        page,
		Size:        size,
	}

	return &response, nil
}
