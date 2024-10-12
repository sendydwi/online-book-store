package product

import (
	"github.com/sendydwi/online-book-store/services/product/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (p *ProductRepository) GetProductById(bookId string) (*model.BookModel, error) {
	var book model.BookModel
	err := p.DB.Model(model.BookModel{}).Where("book_id = ?", bookId).First(book).Error

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (p *ProductRepository) GetProductList(page, size int) (*[]model.BookModel, error) {
	var books []model.BookModel
	err := p.DB.Model(model.BookModel{}).Offset((page - 1) * size).Limit(size).Find(books).Error

	if err != nil {
		return nil, err
	}

	return &books, nil
}
