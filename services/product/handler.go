package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		Svc: Service{
			Repo: Repository{DB: db},
		},
	}
}

func (p *ProductHandler) RegisterHandler(c *gin.Engine) {

}

func (p *ProductHandler) GetProductList(c *gin.Context) {

}

func (p *ProductHandler) GetProductDetail(c *gin.Context) {

}
