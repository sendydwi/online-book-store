package product

import (
	"errors"

	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/product/adapter"
	"gorm.io/gorm"
)

type ProductServiceInterface interface {
	GetProductById(productId int) (*apiproduct.ProductResponse, error)
	GetProductList(page, size int) (*apiproduct.ProductListResponse, error)
}

type Service struct {
	Repo ProductRepositoryInterface
}

func (s *Service) GetProductById(productId int) (*apiproduct.ProductResponse, error) {
	product, err := s.Repo.GetProductById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	bookResponse := adapter.ProductModelToProductResponse(*product)

	return &bookResponse, nil
}

func (s *Service) GetProductList(page, size int) (*apiproduct.ProductListResponse, error) {
	products, err := s.Repo.GetProductList(page, size)
	if err != nil {
		return nil, err
	}

	productResponseList := adapter.ProductModelListToProductResponseList(*products)
	response := apiproduct.ProductListResponse{
		ProductList: productResponseList,
		Page:        page,
		Size:        size,
	}

	return &response, nil
}
