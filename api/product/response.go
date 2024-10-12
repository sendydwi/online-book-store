package apiproduct

import "time"

type ProductResponse struct {
	BookDetail BookDetail `json:"book_detail"`
	Stock      int        `json:"stock"`
}

type BookDetail struct {
	ProductId   string    `json:"product_id"`
	ISBN        string    `json:"isbn"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Author      string    `json:"author"`
	PublishTime time.Time `json:"publish_time"`
	Publisher   string    `json:"publisher"`
	TotalPage   int       `json:"total_page"`
	Description string    `json:"description"`
}

type ProductListResponse struct {
	ProductList []ProductResponse `json:"product_list"`
	Page        int               `json:"page"`
	Size        int               `json:"size"`
}
