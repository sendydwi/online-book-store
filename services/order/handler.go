package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		Svc: Service{
			Repo: Repository{DB: db},
		},
	}
}

func (o *OrderHandler) RegisterHandler(c *gin.Engine) {

}

func (o *OrderHandler) CreateOrder(c *gin.Context) {

}

func (o *OrderHandler) GetOrderDetail(c *gin.Context) {

}

func (o *OrderHandler) GetOrderHistories(c *gin.Context) {

}
