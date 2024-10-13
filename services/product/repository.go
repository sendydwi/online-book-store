package product

import (
	"github.com/sendydwi/online-book-store/services/product/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) GetProductById(bookId int) (*entity.Product, error) {
	var book entity.Product
	err := p.DB.Where("book_id = ? AND is_active = ?", bookId, true).First(book).Error

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (p *ProductRepository) GetProductList(page, size int) (*[]entity.Product, error) {
	var books []entity.Product
	err := p.DB.Offset((page - 1) * size).Limit(size).Find(books).Error

	if err != nil {
		return nil, err
	}
	return &books, nil
}
