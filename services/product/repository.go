package product

import (
	"github.com/sendydwi/online-book-store/services/product/entity"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetProductById(productId int) (*entity.Product, error)
	GetProductList(page, size int) (*[]entity.Product, error)
}

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) GetProductById(productId int) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.Where("product_id = ? AND is_active = ?", productId, true).First(&product).Error

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) GetProductList(page, size int) (*[]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Offset((page - 1) * size).Limit(size).Find(&products).Error

	if err != nil {
		return nil, err
	}
	return &products, nil
}
