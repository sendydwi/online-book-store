package cart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		Svc: Service{
			Repo: Repository{DB: db},
		},
	}
}

func (c *CartHandler) RegisterHandler(r *gin.Engine) {

}

func (c *CartHandler) AddItem() {

}

func (c *CartHandler) RemoveItem() {

}

func (c *CartHandler) DeleteItem() {

}
